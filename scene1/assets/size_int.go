package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, playground")
	var size int = 10000000000000
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MaxFloat64)
	size2 := 9223372036854775807
	size3 := 12000000000000000
	//strsize := "12000000000000000"

	fmt.Printf("Size is %d and %T\n", size, size)
	fmt.Printf("Size is %d and %T\n", size2, size2)
	fmt.Printf("Size is %d and %T\n", size3, size3)
	abc := float64(size2)
	//xyz := 922337203685477580.00
	xyz := 1999999999999999999.00
	fmt.Printf("Size is %.0f and %T\n", abc, abc)
	fmt.Printf("Size is %.0f and %T\n", xyz, xyz)
	fmt.Printf("Size is %d and %T\n", int(xyz), int(xyz))
}

