package datastructure

import "fmt"

func MapSandbox() {
	var m map[string]int = make(map[string]int)

	m["height"] = 178
	fmt.Println(m)
}
