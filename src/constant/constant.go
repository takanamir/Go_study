// 定数

package main

import "fmt"

func main() {
	// const name = "taguchi"
	// name = "fkoji"
	// fmt.Println(name)

	// 整数値の連番にしたい場合
	const (
		sun = iota // 0
		mon         // 1
		tue         // 2
	)
	fmt.Println(sun, mon, tue)

}