package runtimex

import "runtime"

func Caller(skip int) (*runtime.Frame, bool, bool) {
	rpc := make([]uintptr, 1)
	n := runtime.Callers(skip+1, rpc[:])
	if n < 1 {
		return nil, false, false
	}
	frame, more := runtime.CallersFrames(rpc).Next()
	return &frame, more, true
}
