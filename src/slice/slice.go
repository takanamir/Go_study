// 配列 : スライス

package main

import "fmt"

func main() {
	a := [5]int{2, 10, 8, 15, 4}
	s := a[2:4] // [8, 15]
	s[1] = 12

	fmt.Println(a)       // [2 10 8 12 4]
	fmt.Println(s)       // [8 12]
	fmt.Println(len(s)) // 2
	fmt.Println(cap(s)) // 3
}