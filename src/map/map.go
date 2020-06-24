// map

package main

import "fmt"

func main() {
	// m := make(map[string]int)
	// m["taguchi"] = 200
	// m["fkoji"] = 300
	m := map[string]int{"taguchi": 100, "fkoji": 200}
	fmt.Println(m)       // map[fkoji:200 taguchi:100]
	fmt.Println(len(m)) // 2

	delete(m, "taguchi")
	fmt.Println(m)       // map[fkoji:200]

	v, ok := m["fkoji"]
	fmt.Println(v)       // 200
	fmt.Println(ok)      // true
}