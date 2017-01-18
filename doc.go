/*
Metaphone and double metaphone alogrithm implementaion in Golang.
More information about this algorithm can be found on Wikipedia at
http://en.wikipedia.org/wiki/Metaphone.

Example
	// test.go
	package main

	import (
		"fmt"

		"github.com/vividvilla/metaphone"
	)

	func main() {
		inp := "Bangalore"
		mp := metaphone.Metaphone(inp)
		fmt.Printf("metaphone for input %s : %s \n", inp, mp)

		mp, alt := metaphone.DoubleMetaphone(inp)
		fmt.Printf("doublemetaphone for input %s : %s %s \n", inp, mp, alt)
	}


Output
	metaphone for input Bangalore : BNKLR
	doublemetaphone for input Bangalore : PNKL PNKL
*/
package metaphone
