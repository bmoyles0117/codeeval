package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func baseConvert(value, power int, chars string) string {
	bytes := make([]byte, power)
	for i := 0; i < power; i++ {
		bytes[i] = chars[0]
	}

	i := 0
	totalChars := len(chars)

	for value > 0 {
		bytes[i] = chars[value%totalChars]
		value /= totalChars
		i += 1
	}

	return string(bytes)
}

func getPossibilities(power int, chars string) []string {
	totalChars := len(chars)

	totalPossibilities := len(chars)
	for i := 1; i < power; i++ {
		totalPossibilities *= totalChars
	}

	possibilities := make([]string, totalPossibilities)

	for i := 0; i < totalPossibilities; i++ {
		possibilities[i] = baseConvert(i, power, chars)
	}

	return possibilities
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

		power, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("> failed to parse int: %s", err)
		}

		chars := []rune{}
		taken := map[rune]bool{}
		for _, char := range parts[1] {
			if _, exists := taken[char]; !exists {
				chars = append(chars, char)
				taken[char] = true
			}
		}

		possibilities := getPossibilities(power, string(chars))
		sort.Strings(possibilities)

		fmt.Println(strings.Join(possibilities, ","))
	}
}
