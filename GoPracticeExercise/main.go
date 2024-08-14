package main

import "fmt"

type ByteSize int //We create our own type to make it self-documenting

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {
	fmt.Printf("%d \t\t\t %#b\n", KB, KB)
	fmt.Printf("%d \t\t %#b\n", MB, MB)
	fmt.Printf("%d \t\t %#b\n", GB, GB)
	fmt.Printf("%d \t\t %#b\n", TB, TB)
	fmt.Printf("%d \t %#b\n", PB, PB)
	fmt.Printf("%d \t %#b\n", EB, EB)
}

// KB 1024 Bytes
// MB 1024 KBs
// GB 1024 MBs
// TB 1024 GB
// PB 1024 TB
// EB 1024 PB
// Its 2 to the power of 10 to upsize to each one.
//1 << (10 * iota) means "1 shifted left by 10 bits" which is equivalent to multiplying by 2^10. This results in 1024, so KB is 1024 bytes.
//Then iota increments the 1 to a 2 and basically becomes 1024 * 1024
//After that it just keeps adding another multiplication of 1024 to calculate the value of the user defined type ByteSize to eliminate repetitive code
