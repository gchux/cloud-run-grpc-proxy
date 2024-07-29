package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

type (
	Proto2Package struct {
		Proto   string `json:"proto"`
		Package string `json:"package"`
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

	RPC struct {
		Method   *string `json:"method"`
		Request  *string `json:"request"`
		Response *string `json:"response"`
	}

	ProtoDef struct {
		Proto   *string `json:"proto"`
		RPCs    []*RPC  `json:"rpc"`
		Package *string `json:"pkg"`
		Service *string `json:"svc"`
		Go      *string `json:"go"`
	}

	ProtoDefMap map[string]*ProtoDef
)

var (
	proto2rpc = flag.String("proto2rpc", "", "")
	proto2pkg = flag.String("proto2pkg", "", "")
	proto2svc = flag.String("proto2svc", "", "")
	proto2go  = flag.String("proto2go", "", "")
)

func openJSON(jsonFile *string) (*os.File, error) {
	return os.OpenFile(*jsonFile, os.O_RDONLY, 0o444 /* -r--r--r-- */)
}

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

		rpc := &RPC{
			Method:   &v.RPC,
			Request:  &v.Request,
			Response: &v.Response,
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
			os.Exit(3)
		}

		if protoDef, ok := protoDefMap[v.Proto]; ok {
			protoDef.Service = &v.Service
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
			os.Exit(3)
		}

		if protoDef, ok := protoDefMap[v.Proto]; ok {
			protoDef.Go = &v.Go
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

func main() {
	flag.Parse()

	protoDefMap := make(ProtoDefMap)

	sizeOfProto2RPC := loadProto2RPC(proto2rpc, protoDefMap)
	sizeOFPRoto2Go := loadProto2Go(proto2go, protoDefMap)
	sizeOFProto2Pkg := loadProto2Package(proto2pkg, protoDefMap)
	sizeOfProto2Svc := loadProto2Service(proto2svc, protoDefMap)

	size := minOf(sizeOFProto2Pkg, sizeOfProto2Svc, sizeOFPRoto2Go)

	RPCs := make([]*ProtoDef, size)
	index := 0
	for _, protoDef := range protoDefMap {
		if protoDef.Package != nil &&
			protoDef.Service != nil &&
			protoDef.Go != nil {
			RPCs[index] = protoDef
			for _, rpc := range protoDef.RPCs {
				fmt.Fprintf(os.Stderr, "%s.%s/%s(%s,%s) @%s\n",
					*protoDef.Package, *protoDef.Service, *rpc.Method, *rpc.Request, *rpc.Response, *protoDef.Go)
			}
			index += 1
		}
	}
	jsonBytes, err := json.Marshal(RPCs)
	if err == nil {
		io.WriteString(os.Stdout, string(jsonBytes)+"\n")
	}
	fmt.Fprintf(os.Stderr, "SizeOf[protos:%d|packages:%d|services:%d|go:%d] => %d\n",
		sizeOfProto2RPC, sizeOFProto2Pkg, sizeOfProto2Svc, sizeOFPRoto2Go, size)
}
