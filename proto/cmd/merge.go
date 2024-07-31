package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type (
	Proto2Package struct {
		Proto   string `json:"proto"`
		Package string `json:"package"`
	}

	Proto2GoPackage struct {
		Proto       string `json:"proto"`
		PackageFull string `json:"pkg_full"`
		Package     string `json:"pkg"`
	}

	Proto2Service struct {
		Proto   string `json:"proto"`
		Service string `json:"service"`
	}

	Proto2Go struct {
		Proto string `json:"proto"`
		Go    string `json:"go"`
	}

	Proto2RPC struct {
		Proto    string `json:"proto"`
		RPC      string `json:"rpc"`
		Request  string `json:"request"`
		Response string `json:"response"`
	}

	ProtoDefMap map[string]*ProtoDef
)

var (
	proto2rpc   = flag.String("proto2rpc", "", "")
	proto2pkg   = flag.String("proto2pkg", "", "")
	proto2svc   = flag.String("proto2svc", "", "")
	proto2go    = flag.String("proto2go", "", "")
	proto2gopkg = flag.String("proto2go_pkg", "", "")
)

func loadProto2RPC(jsonFilePath *string, protoDefMap ProtoDefMap) uint64 {
	jsonFile, err := openJSON(jsonFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	size := uint64(0)
	s := bufio.NewScanner(jsonFile)
	for s.Scan() {
		var v Proto2RPC
		if unmarshalErr := json.Unmarshal(s.Bytes(), &v); unmarshalErr != nil {
			fmt.Fprintln(os.Stderr, unmarshalErr.Error())
			os.Exit(2)
		}

		isStreamRequest := strings.HasPrefix(v.Request, "stream ")
		isStreamResponse := strings.HasPrefix(v.Response, "stream ")

		requestMessage := v.Request
		if isStreamRequest {
			requestMessage = requestMessage[7:]
		}
		responseMessage := v.Response
		if isStreamResponse {
			responseMessage = responseMessage[7:]
		}

		rpc := &RPC{
			Method: &v.RPC,
			Request: &Message{
				&requestMessage,
				isStreamRequest,
			},
			Response: &Message{
				&responseMessage,
				isStreamResponse,
			},
		}

		protoDef, ok := protoDefMap[v.Proto]
		if ok {
			protoDef.RPCs = append(protoDef.RPCs, rpc)
		} else {
			protoDef = &ProtoDef{
				Proto: &v.Proto,
				RPCs:  []*RPC{rpc},
			}
			protoDefMap[v.Proto] = protoDef
		}

		size += 1
	}

	return size
}

func loadProto2Package(jsonFilePath *string, protoDefMap ProtoDefMap) uint64 {
	jsonFile, err := openJSON(jsonFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	size := uint64(0)
	s := bufio.NewScanner(jsonFile)
	for s.Scan() {
		var v Proto2Package
		if unmarshalErr := json.Unmarshal(s.Bytes(), &v); unmarshalErr != nil {
			fmt.Fprintln(os.Stderr, unmarshalErr.Error())
			os.Exit(2)
		}

		if protoDef, ok := protoDefMap[v.Proto]; ok {
			protoDef.Package = &v.Package
			size += 1
		}
	}
	return size
}

func loadProto2Service(jsonFilePath *string, protoDefMap ProtoDefMap) uint64 {
	jsonFile, err := openJSON(jsonFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	size := uint64(0)
	s := bufio.NewScanner(jsonFile)
	for s.Scan() {
		var v Proto2Service
		if unmarshalErr := json.Unmarshal(s.Bytes(), &v); unmarshalErr != nil {
			fmt.Fprintln(os.Stderr, unmarshalErr.Error())
			os.Exit(4)
		}

		if protoDef, ok := protoDefMap[v.Proto]; ok {
			if protoDef.Services == nil {
				protoDef.Services = []*string{&v.Service}
			} else {
				protoDef.Services = append(protoDef.Services, &v.Service)
			}
			size += 1
		}
	}
	return size
}

func loadProto2Go(jsonFilePath *string, protoDefMap ProtoDefMap) uint64 {
	jsonFile, err := openJSON(jsonFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	size := uint64(0)
	s := bufio.NewScanner(jsonFile)
	for s.Scan() {
		var v Proto2Go
		if unmarshalErr := json.Unmarshal(s.Bytes(), &v); unmarshalErr != nil {
			fmt.Fprintln(os.Stderr, unmarshalErr.Error())
			os.Exit(5)
		}

		if protoDef, ok := protoDefMap[v.Proto]; ok {
			if protoDef.Go == nil {
				protoDef.Go = &Go{
					SrcFile: &v.Go,
				}
			} else {
				protoDef.Go.SrcFile = &v.Go
			}
			size += 1
		}
	}
	return size
}

func loadProto2GoPackage(jsonFilePath *string, protoDefMap ProtoDefMap) uint64 {
	jsonFile, err := openJSON(jsonFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	size := uint64(0)
	s := bufio.NewScanner(jsonFile)
	for s.Scan() {
		var v Proto2GoPackage
		if unmarshalErr := json.Unmarshal(s.Bytes(), &v); unmarshalErr != nil {
			fmt.Fprintln(os.Stderr, unmarshalErr.Error())
			os.Exit(3)
		}

		if protoDef, ok := protoDefMap[v.Proto]; ok {
			goPkgFullParts := goPkgRegexp.FindStringSubmatch(v.PackageFull)
			if len(goPkgFullParts) == 0 {
				continue
			}

			namespace := fmt.Sprintf("%s_%s_%s",
				goPkgFullParts[1], goPkgFullParts[2], v.Package)

			if protoDef.Go == nil {
				protoDef.Go = &Go{
					Package:     &v.Package,
					PackageFull: &v.PackageFull,
					Namespace:   &namespace,
				}
			} else {
				protoDef.Go.Package = &v.Package
				protoDef.Go.PackageFull = &v.PackageFull
				protoDef.Go.Namespace = &namespace
			}
			size += 1
		}
	}

	return size
}

func minOf(nums ...uint64) uint64 {
	min := nums[0]
	for _, i := range nums {
		if min > i {
			min = i
		}
	}
	return min
}

func main_merge() {
	protoDefMap := make(ProtoDefMap)

	sizeOfProto2RPC := loadProto2RPC(proto2rpc, protoDefMap)
	sizeOFPRoto2Go := loadProto2Go(proto2go, protoDefMap)
	sizeOFPRoto2GoPkg := loadProto2GoPackage(proto2gopkg, protoDefMap)
	sizeOFProto2Pkg := loadProto2Package(proto2pkg, protoDefMap)
	sizeOfProto2Svc := loadProto2Service(proto2svc, protoDefMap)

	size := minOf(sizeOFProto2Pkg, sizeOfProto2Svc, sizeOFPRoto2Go, sizeOFPRoto2GoPkg)
	RPCs := make([]*ProtoDef, size)

	index := 0
	for _, protoDef := range protoDefMap {
		if protoDef.Package != nil &&
			protoDef.Services != nil &&
			protoDef.Go != nil &&
			protoDef.Go.SrcFile != nil &&
			protoDef.Go.PackageFull != nil &&
			!strings.Contains(*protoDef.Go.PackageFull, "/internal/") {
			RPCs[index] = protoDef
			for _, rpc := range protoDef.RPCs {
				var services strings.Builder
				sizeOfServices := len(protoDef.Services)
				for i := sizeOfServices - 1; i >= 0; i-- {
					services.WriteString(*protoDef.Services[i])
					if i > 0 {
						services.WriteString(",")
					}
				}
				fmt.Fprintf(os.Stderr, "%s.[%s]/%s(%s,%s) @%s\n",
					*protoDef.Package, services.String(), *rpc.Method, *rpc.Request, *rpc.Response, *protoDef.Go.PackageFull)
			}
			index += 1
		}
	}

	RPCs = slices.DeleteFunc(RPCs,
		func(rpc *ProtoDef) bool {
			return rpc == nil
		},
	)
	RPCs = slices.Clip(RPCs)

	jsonBytes, err := json.Marshal(RPCs)
	if err == nil {
		io.WriteString(os.Stdout, string(jsonBytes)+"\n")
	}
	fmt.Fprintf(os.Stderr, "SizeOf[protos:%d|packages:%d|services:%d|go:%d] => %d/%d\n",
		sizeOfProto2RPC, sizeOFProto2Pkg, sizeOfProto2Svc, sizeOFPRoto2Go, index, size)
}
