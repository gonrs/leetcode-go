package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	strings.Clone("")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Code(scanner.Text()))
}
// use can use fmt, strings package
func Code(s string) string {
	return s
}
