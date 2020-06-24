// range

package main

import "fmt"

func main() {
	 s := []int{2, 3, 8}
	 for i, j := range s {
	     fmt.Println(i, j)
	 }

	 fmt.Println("")

	 for _, x := range s {
	     fmt.Println(x)
	 }

	 fmt.Println("")

	m := map[string]int{"taguchi": 200, "fkoji": 300}
	for k, v := range m {
		fmt.Println(k, v)
	}
}