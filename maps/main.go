package main

import (
	"fmt"
)

func main() {
	// Create a map with a key of type string and a value of type int
	m := make(map[int]string)
	m[1] = "a"
	m[2] = "b"
	m[3] = "c"
	for k, v := range m {
		fmt.Println("{" + fmt.Sprint(k) + ":" + fmt.Sprintf("%v", v) + "}")
	}
	// Create a mao witrh a key of type string and a value of type string
	m2 := make(map[string]string)
	m2["a"] = "1"
	m2["b"] = "2"
	m2["c"] = "3"
	for k, v := range m2 {
		fmt.Println("{" + k + ":" + fmt.Sprintf("%v", v) + "}")
	}

	var m3 map[string]interface{}
	m3 = make(map[string]interface{})
	m3["a"] = 1
	m3["b"] = "2"
	m3["c"] = 3
	for k, v := range m3 {
		fmt.Println("{" + k + ":" + fmt.Sprintf("%v", v) + "}")
	}
}
