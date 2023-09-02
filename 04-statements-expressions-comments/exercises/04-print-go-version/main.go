// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"runtime"
)

// ---------------------------------------------------------
// EXERCISE: Print the Go Version
//
//  1. Look at the runtime package documentation
//  2. Find the func that returns the Go version
//  3. Print the Go version by calling that func
//
// HINT
//
//	It's here: https://golang.org/pkg/runtime
//
// EXPECTED OUTPUT
//
//	"go1.10"
//
// ---------------------------------------------------------
func showSystemInfo() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Go Version\t", runtime.Version())
	fmt.Println("Go Compiler\t", runtime.Compiler)
	fmt.Println("Go Goroutines\t", runtime.NumGoroutine())
}
func main() {
	// ?
	fmt.Println(runtime.Version())

	go func() {
		fmt.Println("hello")
	}()
	showSystemInfo()

}
