// 変数
// 変数名: 1文字目に注意

package main

import "fmt"

func main() {
    // var msg string
    // msg = "hello world"
    // var msg = "hello world"
    msg := "hello world"
    fmt.Println(msg)

    // var a, b int
    // a, b = 10, 15
    a, b := 10, 15

    var (
        c int
        d string
    )
    c = 20
    d = "hoge"

    fmt.Println(a + b)
    fmt.Println(c)
    fmt.Println(d)
}