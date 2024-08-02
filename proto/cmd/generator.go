package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	pkgPrefix    = "package "
	goPkgPrefix  = "option go_package "
	svcPrefix    = "service "
	hostPrefix   = " option google.api.default_host "
	rpcPrefix    = " rpc "
	msgPrefix    = "message "
	streamPrefix = "stream "
)

var (
	protoSeparator = []byte(":")
	rpcRegexp      = regexp.MustCompile(`(?P<name>.+?)\s(?P<request>(?:stream\s)?.+?)\s(?P<response>(?:stream\s)?.+?)$`)
	msgRegexp      = regexp.MustCompile(`(\s+)message\s+(.+?)$`)
)

var (
	rawProtosFilePath    = flag.String("protos", "", "")
	proto2GoJSONFilePath = flag.String("proto2go", "", "")
)

func main_generator() {
	strings.Count("this is a test", " ")

	rawProtosFile, err := openFile(rawProtosFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	protos := make(map[string]*ProtoDef)

	s := bufio.NewScanner(rawProtosFile)
	for s.Scan() {
		lineParts := bytes.SplitN(s.Bytes(), protoSeparator, 2)

		protoFilePath := string(lineParts[0])
		protoFileName := filepath.Base(string(protoFilePath))
		protoFile := strings.TrimSuffix(protoFilePath, filepath.Ext(protoFileName))

		var proto *ProtoDef
		if _, ok := protos[protoFilePath]; !ok {
			protos[protoFilePath] = &ProtoDef{
				Proto: &protoFile,
			}
		}
		proto = protos[protoFilePath]

		entry := string(lineParts[1])

		switch {
		case strings.HasPrefix(entry, pkgPrefix):
			pkgName := entry[len(pkgPrefix):]
			proto.Package = &pkgName

		case strings.HasPrefix(entry, goPkgPrefix):
			goPkg := entry[len(goPkgPrefix):]
			goPkgParts := strings.SplitN(goPkg, " ", 2)

			goPkgFull := goPkgParts[0]
			goPkgAlias := ""

			if len(goPkgParts) > 1 {
				goPkgAlias = goPkgParts[1]
			} else {
				goPkgAlias = filepath.Base(goPkgFull)
			}

			proto.Go = &Go{
				Package:     &goPkgAlias,
				PackageFull: &goPkgFull,
			}

		case strings.HasPrefix(entry, svcPrefix):
			serviceName := entry[len(svcPrefix):]

			service := &Service{Service: &serviceName}
			if proto.Services == nil {
				proto.Services = []*Service{service}
			} else {
				proto.Services = append(proto.Services, service)
			}

		case strings.HasPrefix(entry, hostPrefix):
			hostName := entry[len(hostPrefix):]
			proto.Services[len(proto.Services)-1].Host = &hostName

		case strings.HasPrefix(entry, rpcPrefix):
			rpc := entry[len(rpcPrefix):]

			rpcMatch := rpcRegexp.FindStringSubmatch(rpc)

			if len(rpcMatch) > 0 {
				service := proto.Services[len(proto.Services)-1]

				if service.RPCs == nil {
					service.RPCs = make([]*RPC, 0)
				}

				method := rpcMatch[1]
				request := rpcMatch[2]
				response := rpcMatch[3]

				isStreamRequest := strings.HasPrefix(request, streamPrefix)
				isStreamResponse := strings.HasPrefix(response, streamPrefix)

				if isStreamRequest {
					request = request[len(streamPrefix):]
				}
				if isStreamResponse {
					response = response[len(streamPrefix):]
				}

				service.RPCs = append(service.RPCs, &RPC{
					Method: &method,
					Request: &Message{
						Type:   &request,
						Stream: isStreamRequest,
					},
					Response: &Message{
						Type:   &response,
						Stream: isStreamResponse,
					},
				})
			}

		case strings.HasPrefix(entry, msgPrefix):
			msg := entry[len(msgPrefix):]

			if proto.Messages == nil {
				proto.Messages = make([]*Message, 0)
			}

			_msg := &Message{Type: &msg}
			if proto.Messages == nil {
				proto.Messages = []*Message{_msg}
			} else {
				proto.Messages = append(proto.Messages, _msg)
			}

		case msgRegexp.MatchString(entry):
			msgMatch := msgRegexp.FindStringSubmatch(entry)

			depth := len(msgMatch[1])
			msg := &Message{Type: &msgMatch[2]}

			message := proto.Messages[len(proto.Messages)-1]
			for i := 1; i < depth; i++ {
				message = message.Messages[len(message.Messages)-1]
			}

			if message.Messages == nil {
				// just found a new level
				message.Messages = []*Message{msg}
			} else {
				// level was already found, update parent
				message.Messages = append(message.Messages, msg)
			}
		}
	}

	jsonFile, err := openFile(proto2GoJSONFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	s = bufio.NewScanner(jsonFile)
	for s.Scan() {
		var v Proto2Go
		if unmarshalErr := json.Unmarshal(s.Bytes(), &v); unmarshalErr != nil {
			fmt.Fprintln(os.Stderr, unmarshalErr.Error())
			os.Exit(2)
		}

		if protoDef, ok := protos[v.Proto]; ok {
			if protoDef.Go != nil {
				protoDef.Go.SrcFile = &v.Go
			}
		}
	}

	for _, protoDef := range protos {
		jsonBytes, err := json.Marshal(protoDef)
		if err == nil {
			io.WriteString(os.Stdout, string(jsonBytes)+"\n")
		} else {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}
}
