// main.go
package main

import (
	"fmt"
	"mymath"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("Hello, world.  Sqrt(2) = %v\n", mymath.Sqrt(2))
	}
	return
}
