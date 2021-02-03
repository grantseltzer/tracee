package main

import "C"

import (
	"fmt"
	"os"
	"os/signal"

	bpf "github.com/aquasecurity/tracee/libbpfgo"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	bpfModule, err := bpf.NewModuleFromFile("self.bpf.o")
	if err != nil {
		panic(err)
	}

	bpfModule.BPFLoadObject()

	prog, err := bpfModule.GetProgram("self")
	if err != nil {
		panic(err)
	}

	go bpf.TracePrint()

	_, err = prog.AttachKprobe("__x64_sys_execve")
	if err != nil {
		panic(err)
	}

	<-sig
	fmt.Println("complete")
}
