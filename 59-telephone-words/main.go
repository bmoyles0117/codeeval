package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var digitMappings = map[rune]string{
	'0': "0",
	'1': "1",
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func getPermutations(possibleValueSets []string, offsets []int) []string {
	if offsets == nil {
		return []string{}
	}

	permutation := ""
	for i := 0; i < len(possibleValueSets); i++ {
		permutation += string(possibleValueSets[i][offsets[i]])
	}

	permutations := append([]string{permutation}, getPermutations(possibleValueSets, incrementOffsets(possibleValueSets, offsets))...)

	return permutations
}

func incrementOffsets(possibleValueSets []string, offsets []int) []int {
	changed := false

	for i := len(possibleValueSets) - 1; i >= 0 && !changed; i-- {
		offsets[i] = (offsets[i] + 1) % len(possibleValueSets[i])
		if offsets[i] != 0 {
			changed = true
		}
	}

	// nil out the offsets once we've maxed out
	if !changed {
		return nil
	}

	return offsets
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

		possibleValueSets := []string{}
		for _, char := range string(line) {
			possibleValueSets = append(possibleValueSets, digitMappings[char])
		}

		offsets := make([]int, len(possibleValueSets))

		fmt.Println(strings.Join(getPermutations(possibleValueSets, offsets), ","))
	}
}
