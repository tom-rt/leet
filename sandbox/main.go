package main

import (
	"fmt"
	c "sandbox/concurrency"
)

func main() {
	// Datastructures
	fmt.Println("#### Datastructure sandbox")
	// d.ArraySandbox()
	// d.SliceSandbox()
	// d.MapSandbox()

	// Concurrency
	fmt.Println("#### Concurrency sandbox")
	// c.RoutineSandbox()
	// c.WgSandbox()
	c.SelectSandbox()
}
