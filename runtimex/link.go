package runtimex

import (
	_ "runtime"
	_ "unsafe" // for linkname
)

//go:linkname Fastrand runtime.fastrand
func Fastrand() uint32
