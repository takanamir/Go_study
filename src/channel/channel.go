// channel

package main

import (
	"fmt"
	"time"
)

func task1(result chan string) { // chan型としてあげて、受け渡すデータの型を書いていく
	time.Sleep(time.Second * 2)
	// fmt.Println("task1 finished!")
	result <- "task1 result"
}

func task2() {
	fmt.Println("task2 finished!")
}

func main() {
	result := make(chan string)

	go task1(result) // task1をバックグラウンドで実行
	go task2()       // task2もバックグラウンドで実行

	fmt.Println(<-result)

	time.Sleep(time.Second * 3)
	fmt.Println("All finished!")
}