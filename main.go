package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	str := os.Args[1]
	len := utf8.RuneCountInString(str)

	printHeader(len)
	printBody(str)
	printFooter(len)
}

func printHeader(len int) {
	fmt.Print("＿")
	for i := 0; i <= len; i++ {
		fmt.Print("人")
	}
	fmt.Println("＿")
}

func printBody(str string) {
	fmt.Print("＞ ")
	fmt.Print(str)
	fmt.Println(" ＜")
}

func printFooter(len int) {
	fmt.Print("￣")
	for i := 0; i <= len; i++ {
		fmt.Print("^Y")
	}
	fmt.Println("￣")
}
