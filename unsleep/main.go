package main

import (
	"log"
	"syscall"
	"time"
)

// Execution States
const (
	EsSystemRequired = 0x00000001
	EsContinuous     = 0x80000000
)

var pulseTime = 30 * time.Second

// cf https://stackoverflow.com/a/45615169

func main() {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	setThreadExecStateProc := kernel32.NewProc("SetThreadExecutionState")

	pulse := time.NewTicker(pulseTime)

	log.Println("Starting keep alive poll... (silence)")
	for {
		select {
		case <-pulse.C:
			setThreadExecStateProc.Call(uintptr(EsSystemRequired))
		}
	}
}
