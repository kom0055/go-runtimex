package runtimex

import (
	"log"
	"runtime/debug"

	"github.com/kom0055/go-runtimex/langx/bytesconv"
)

func GoRoutine(method func()) {
	if method == nil {
		return
	}
	go func() {
		defer Guard()
		method()
	}()

}

func Guard() {
	if r := recover(); r != nil {
		log.Printf("recover panic: %v,  stack: %v\r\n",
			r, bytesconv.BytesToString(debug.Stack()))
	}
}
