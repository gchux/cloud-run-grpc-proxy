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
	aliasfixJSONFilePath = flag.String("aliasfix", "", "")
)

func main_generator() {
	strings.Count("this is a test", " ")

	rawProtosFile, err := openFile(rawProtosFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	protos := make(map[string]*ProtoDef)
	goPkg2Proto := make(map[string][]string)

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

			if protos, ok := goPkg2Proto[goPkgFull]; ok {
				goPkg2Proto[goPkgFull] = append(protos, protoFilePath)
			} else {
				goPkg2Proto[goPkgFull] = []string{protoFilePath}
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
	rawProtosFile.Close()

	jsonFile, err := openFile(aliasfixJSONFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(4)
	}
	d := json.NewDecoder(jsonFile)
	for {
		var v AliasFix
		err = d.Decode(&v)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(5)
		}

		// see: https://github.com/googleapis/google-cloud-go/blob/main/internal/aliasfix/mappings.go
		// migration mappings currently ( and probably never will... ) only contains `StatusMigrated`
		// so there's no way to make any other decisions here...
		// see: https://github.com/googleapis/google-cloud-go/issues/10700
		for _, protoFile := range goPkg2Proto[v.New] {
			if protoDef, ok := protos[protoFile]; ok {
				if protoDef.Go == nil {
					pkg := "pb"
					protoDef.Go = &Go{
						Package:           &pkg,
						PackageFull:       &v.New,
						LegacyPackageFull: &v.Old,
					}
				}
				protoDef.Go.Migrated = true
			}
		}
	}
	jsonFile.Close()

	jsonFile, err = openFile(proto2GoJSONFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
	d = json.NewDecoder(jsonFile)
	for {
		var v Proto2Go
		err = d.Decode(&v)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(3)
		}

		if protoDef, ok := protos[v.Proto]; ok {
			if protoDef.Go != nil {
				protoDef.Go.SrcFile = &v.Go
				if !protoDef.Go.Migrated && strings.HasPrefix(v.Go, "googleapis/") {
					pkg := *protoDef.Go.SrcFile
					pkgParts := strings.Split(pkg, "/")
					sizeOfPkgParts := len(pkgParts)
					pkgFull := "google.golang.org/genproto/" + strings.Join(pkgParts[:sizeOfPkgParts-1], "/")
					pkgAlias := strings.Join(strings.SplitN(pkgParts[sizeOfPkgParts-1], ".", 3)[:2], "")
					protoDef.Go.Package = &pkgAlias
					protoDef.Go.LegacyPackageFull = &pkgFull
				}
			}
		}

	}
	jsonFile.Close()

	for _, protoDef := range protos {
		jsonBytes, err := json.Marshal(protoDef)
		if err == nil {
			io.WriteString(os.Stdout, string(jsonBytes)+"\n")
		} else {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}
}
