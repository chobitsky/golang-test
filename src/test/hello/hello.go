package main

import (
	"fmt"
	"strconv"
	"test/hoge"
)

var message string = "hello world"

func main() {
	local_message := "日本語"
	for i := 0; i < 5; i++ {
		// 文字列結合
		fmt.Println(message+strconv.Itoa(i))
	}
	fmt.Println(local_message)
	hoge.HogePrint()
	fmt.Println(hoge.Sum(1, 4))
}
