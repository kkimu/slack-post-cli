package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)

	if len(os.Args) != 4 {
		fmt.Println("指定された引数の数が間違っています。")
		os.Exit(1)
	}

	fmt.Printf("実行ファイル名: %s\n", os.Args[0])
	fmt.Printf("引数1: %s\n", os.Args[1])
	fmt.Printf("引数2: %s\n", os.Args[2])
	fmt.Printf("引数3: %s\n", os.Args[3])
}
