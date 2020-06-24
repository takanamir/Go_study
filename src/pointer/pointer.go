// ポインタ
// アドレスを指し示す変数
// 演算はできない

package main

import "fmt"

func main() {
	a := 5
	var pa *int
	pa = &a // &a = aのアドレス
	// paの領域にあるデータの値 = *pa
	fmt.Println(pa)  // 0xc00000a0b8
	fmt.Println(*pa) // 5
}