// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Declare float64
//
//  1. Declare and print a variable with a float64 type
//
//  2. The declared variable's name should be: brightness
//
// EXPECTED OUTPUT
//  0
// ---------------------------------------------------------

func main() {
	// var ? ?
	// ?
	var brightness float64
	brightness = 3.000000009
	great := 3.000030000
	fmt.Println("brightness is:", brightness)
	fmt.Println("great is:", great)

	if brightness > great {
		println("brightness is great")
	} else {
		println("brightness is not great")
	}
}
