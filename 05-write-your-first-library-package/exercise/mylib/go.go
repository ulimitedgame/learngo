package mylib

import "runtime"

func Version() string {
	return runtime.Version()
}
