// for

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		// if i == 3 { break }
		if i == 3 {
			continue
		}
		fmt.Print(i, " ")
	}

	fmt.Println("")

	j := 0
	for j < 10 {
		fmt.Print(j, " ")
		j++
	}

	fmt.Println("")

	k := 0
	for {
		fmt.Print(k, " ")
		k++
		if k == 3 {
			break
		}
	}
}