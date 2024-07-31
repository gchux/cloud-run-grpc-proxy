package main

import (
	"flag"
	"os"
	"regexp"
)

type (
	Message struct {
		Type   *string
		Stream bool
	}
	RPC struct {
		Method   *string  `json:"method"`
		Request  *Message `json:"request"`
		Response *Message `json:"response"`
	}

	Go struct {
		SrcFile     *string `json:"src"`
		Package     *string `json:"pkg"`
		PackageFull *string `json:"pkg_full"`
		Namespace   *string `json:"ns"`
	}

	ProtoDef struct {
		Proto    *string   `json:"proto"`
		RPCs     []*RPC    `json:"rpc"`
		Package  *string   `json:"pkg"`
		Services []*string `json:"svc"`
		Go       *Go       `json:"go"`
	}
)

var exec = flag.String("exec", "", "")

var goPkgRegexp = regexp.MustCompile(`.*?/go/(.+?)/(?:.*?/)*?(apiv.+?)/.*`)

func openJSON(jsonFile *string) (*os.File, error) {
	return os.OpenFile(*jsonFile, os.O_RDONLY, 0o444 /* -r--r--r-- */)
}

func main() {
	flag.Parse()

	switch *exec {
	default:
		main_codegen()
	case "merge":
		main_merge()
	case "codegen":
		main_codegen()
	}
}
