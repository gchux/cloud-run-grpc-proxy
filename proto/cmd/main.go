package main

import (
	"flag"
	"os"
)

type (
	Proto2Go struct {
		Proto string `json:"proto"`
		Go    string `json:"go"`
	}

	Message struct {
		Type     *string    `json:"type"`
		Stream   bool       `json:"stream,omitempty"`
		Messages []*Message `json:"msgs,omitempty"`
	}

	RPC struct {
		Method   *string  `json:"method"`
		Request  *Message `json:"request"`
		Response *Message `json:"response"`
	}

	Go struct {
		SrcFile     *string `json:"src,omitempty"`
		Package     *string `json:"pkg,omitempty"`
		PackageFull *string `json:"pkg_full,omitempty"`
		Namespace   *string `json:"-"`
	}

	Service struct {
		Service *string `json:"svc,omitempty"`
		Host    *string `json:"host,omitempty"`
		RPCs    []*RPC  `json:"rpcs,omitempty"`
	}

	ProtoDef struct {
		Proto    *string    `json:"proto"`
		Package  *string    `json:"pkg"`
		Services []*Service `json:"svcs,omitempty"`
		Messages []*Message `json:"msgs,omitempty"`
		Go       *Go        `json:"go,omitempty"`
	}
)

var exec = flag.String("exec", "", "")

func openFile(filePath *string) (*os.File, error) {
	return os.OpenFile(*filePath, os.O_RDONLY, 0o444 /* -r--r--r-- */)
}

func main() {
	flag.Parse()

	switch *exec {
	default:
		main_generator()
	case "codegen":
		main_codegen()
	case "generator":
		main_generator()
	}
}
