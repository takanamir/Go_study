// 配列
/*
var a1 int
var a2 int
*/

package main

import "fmt"

func main() {
	// var a [5]int // a[0] - a[4]
	// a[2] = 3
	// a[4] = 10
	// fmt.Println(a[2])
	// b := [3]int{1, 3, 5}
	b := [...]int{1, 3, 5}
	fmt.Println(b)      // [1 3 5]
	fmt.Println(len(b)) // 3
}