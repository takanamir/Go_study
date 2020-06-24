// 構造体

package main

import "fmt"

type user struct {
	name  string
	score int
}

func main() {
	// u := new(user)
	// // (*u).name = "taguchi"
	// u.name = "taguchi"
	// u.score = 20

	// u := user{"taguchi", 50}
	u := user{name: "taguchi", score: 50}
	fmt.Println(u) // {taguchi 50}
}