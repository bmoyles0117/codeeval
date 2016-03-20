package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func contains(haystack, needle string, sequenced bool) bool {
	if len(needle) <= 0 {
		return true
	}

	if len(haystack) < len(needle) {
		return false
	}

	if sequenced && haystack[0] != needle[0] {
		return false
	}

	for i, char := range haystack {
		if char == rune(needle[0]) && contains(haystack[i+1:], needle[1:], true) {
			return true
		}
	}

	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("> failed to read line: %s", err)
		}

		parts := strings.Split(string(line), ",")
		haystack, needle := parts[0], parts[1]

		if contains(haystack, needle, false) {
			fmt.Println("1")
		} else {
			fmt.Println("0")
		}
	}
}
