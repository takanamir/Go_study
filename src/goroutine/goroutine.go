// goroutine : 並行処理

package main

import (
	"fmt"
	"time"
)

// task2 が始まる前に task1 の処理待ちをする

func task1() {
	time.Sleep(time.Second * 2) // 時間待ち
	fmt.Println("task1 finished!")
}

func task2() {
	fmt.Println("task2 finished!")
}

func main() {
	go task1()
	go task2()

	time.Sleep(time.Second * 3)
	fmt.Println("All finished!")
}