package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	//fmt.Println(input)
	m := map[string]int{
		"A": 0,
		"C": 0,
		"T": 0,
		"G": 0,
	}
	//fmt.Println(m)
	for _, c := range input {
		//fmt.Printf("%d: %s\n", i, string(c))
		//if v, ok := m["Gigel"]; ok {
		//}
		m[string(c)]++
	}
	//fmt.Println(m)
	for _, c := range []string{"A", "C", "G", "T"} {
		fmt.Printf("%d ", m[c])
	}
	fmt.Println("")
}
