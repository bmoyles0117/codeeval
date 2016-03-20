package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isPalindrome(value string) bool {
	for i := 0; i < len(value)/2; i++ {
		if value[i] != value[len(value)-i-1] {
			return false
		}
	}

	return true
}

func reverseString(value string) string {
	runes := make([]byte, len(value))

	for i := 0; i < len(value); i++ {
		runes[i] = value[len(value)-i-1]
	}

	return string(runes)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value := scanner.Text()

		for i := 0; i <= 100; i++ {
			if isPalindrome(value) {
				fmt.Printf("%d %s\n", i, value)
				break
			}

			intValue, _ := strconv.Atoi(value)
			intValueReversed, _ := strconv.Atoi(reverseString(value))
			value = strconv.Itoa(intValue + intValueReversed)
		}
	}
}
