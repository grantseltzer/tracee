package main

import "C"

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

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
	prog, err := bpfModule.GetProgram("kprobe__sys_mmap")
	if err != nil {
		panic(err)
	}

	_, err = prog.AttachKprobe("__x64_sys_mmap")
	if err != nil {
		panic(err)
	}

	eventsChannel := make(chan []byte)
	rb, err := bpfModule.InitRingBuf("events", eventsChannel)
	if err != nil {
		panic(err)
	}
	rb.Start()

	go func() {
		for {
			select {
			case z := <-eventsChannel:
				pid := binary.LittleEndian.Uint32(z[0:5])
				fmt.Printf("PID %d", int(pid))
				fmt.Printf("\t%s\n", strings.SplitN(string(z[5:]), "\u0000", 2))
			case <-sig:
				log.Fatal("exiting")
				rb.Stop()
				rb.Close()
				log.Fatal("exited")
			}
		}
	}()

	<-sig
	rb.Stop()
	rb.Close()

	fmt.Println("complete")
}
