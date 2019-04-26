package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	for _, c := range text {
		//fmt.Println(c, c == 'a')

		switch c {
		case 'T':
			fmt.Print("U")
		default:
			fmt.Print(string(c))
		}
	}
}
