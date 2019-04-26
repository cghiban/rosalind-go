package main

import (
	"bufio"
	"fmt"
	"os"
)

func v1(text string) {
	for i := len(text) - 1; i >= 0; i-- {
		//fmt.Println(i, text[i])
		//continue

		switch text[i] {
		case 'T':
			fmt.Print("A")
		case 'C':
			fmt.Print("G")
		case 'G':
			fmt.Print("C")
		case 'A':
			fmt.Print("T")
			//default:
			//	fmt.Print(text[i])
		}
	}
	fmt.Println("")
}

func v2(text string) string {
	out := ""
	for i := len(text) - 1; i >= 0; i-- {
		switch text[i] {
		case 'T':
			out += "A"
		case 'C':
			out += "G"
		case 'G':
			out += "C"
		case 'A':
			out += "T"
		}
	}
	fmt.Println(out)
	return out
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	v1(text)
	v2(text)
}
