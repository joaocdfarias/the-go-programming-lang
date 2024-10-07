package main

import (
	"fmt"
)

func main() {
	var x int = 10
	var y float64 = 30.2
	var z float64 = float64(x) + y
	var d int = x + int(y)

	fmt.Println(x) // 10
	fmt.Println(y) // 30.2
	fmt.Println(z) // 10.0 + 30.2 = 40.2
	fmt.Println(d) // 10 + 30 = 40
}
