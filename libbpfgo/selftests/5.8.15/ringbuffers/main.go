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
	defer bpfModule.Close()

	bpfModule.BPFLoadObject()

	prog, err := bpfModule.GetProgram("self")
	if err != nil {
		panic(err)
	}

	_, err = prog.AttachKprobe("__x64_sys_execve")
	if err != nil {
		panic(err)
	}

	eventsChannel := make(chan []bytes)
	rb, err := InitRingBuf("events", eventsChannel)
	if err != nil {
		panic(err)
	}
	rb.Start()

	go func() {
		for {
			switch {
				case z := <-eventsChannel:
					fmt.Println(z)
			}
		}
	}()

	<-sig
	rb.Stop()
	rb.Close()

	fmt.Println("complete")
}
