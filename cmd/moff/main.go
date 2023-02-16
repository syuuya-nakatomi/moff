package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// コマンドライン引数を解析する
	var text string
	flag.StringVar(&text, "text", "", "Text to reverse")
	flag.Parse()

	// テキストが指定されているかどうかを確認する
	if text == "" {
		flag.Usage()
		os.Exit(1)
	}

	// テキストを逆順にする
	reversed := reverse(text)
	fmt.Println(reversed)
}

// 文字列を逆順にする
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
