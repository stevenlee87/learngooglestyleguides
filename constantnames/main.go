package main

import "fmt"

// Good:
const MaxPacketSize = 512
const (
	ExecuteBit = 1 << iota
	WriteBit
	ReadBit
)

var a int

func main() {
	fmt.Println("ExecuteBit is: ", ExecuteBit)
	fmt.Println("WriteBit is: ", WriteBit)
	fmt.Println("ReadBit is: ", ReadBit)

	fmt.Println("var a: ", a)
}
