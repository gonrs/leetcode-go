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
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Code([]byte(scanner.Text())))
}

// use can use fmt package
func Code(chars []byte) int {
	if len(chars) <= 1 {
		return len(chars)
	}
	currentGroup := chars[0]
	count := 0
	result := 0
	index := 0
	for _, char := range chars {
		if char == currentGroup {
			count++
		} else {
			result++
			chars[index] = currentGroup
			index++
			if count > 1 {
				for _, digit := range fmt.Sprintf("%d", count) {
					chars[index] = byte(digit)
					result++
					index++
				}
			}
			currentGroup = char
			count = 1
		}
	}
	result++
	chars[index] = currentGroup
	index++
	if count > 1 {
		for _, digit := range fmt.Sprintf("%d", count) {
			chars[index] = byte(digit)
			result++
			index++
		}
	}
	// res := "func Code(chars []byte) int {\n	if len(chars) <= 1 {\n		return len(chars)\n	}\n	currentGroup := chars[0]\n	count := 0\n	result := 0\n	index := 0\n	for _, char := range chars {\n		if char == currentGroup {\n			count++\n		} else {\n			result++\n			chars[index] = currentGroup\n			index++\n			if count > 1 {\n				for _, digit := range fmt.Sprintf(\"%d\", count) {\n					chars[index] = byte(digit)\n					result++\n					index++\n				}\n			}\n			currentGroup = char\n			count = 1\n		}\n	}\n	result++\n	chars[index] = currentGroup\n	index++\n	if count > 1 {\n		for _, digit := range fmt.Sprintf(\"%d\", count) {\n			chars[index] = byte(digit)\n			result++\n			index++\n		}\n	}\n	return result\n}"
	return result
}
