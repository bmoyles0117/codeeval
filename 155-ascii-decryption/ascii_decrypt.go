package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile("^[a-zA-Z]*$")

type GramReference struct {
	characterCodes []int
	total          int
}

func decode(characterCodes []int, offset int) string {
	buffer := bytes.NewBuffer([]byte{})
	for i := range characterCodes {
		buffer.WriteByte(byte(characterCodes[i] - offset))
	}

	return buffer.String()
}

func intArrayToString(vals []int) string {
	buffer := bytes.NewBuffer([]byte{})
	for i := range vals {
		buffer.WriteString(strconv.Itoa(vals[i]))
	}

	return buffer.String()
}

func main() {
	inputFile := os.Args[1]

	file, err := os.Open(inputFile)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF && line == "" {
			break
		}
		line = strings.TrimSuffix(line, "\n")

		parts := strings.SplitN(line, " | ", 3)

		wordLength, _ := strconv.Atoi(parts[0])
		lastLetter := parts[1]

		characterCodes := []int{}
		for _, code := range strings.Split(parts[2], " ") {
			characterCode, _ := strconv.Atoi(code)
			characterCodes = append(characterCodes, characterCode)
		}

		gramReferences := map[string]GramReference{}
		for i := 0; i < len(characterCodes)-wordLength; i += 1 {
			gramReferenceKey := intArrayToString(characterCodes[i : i+5])
			gramReference, exists := gramReferences[gramReferenceKey]
			if !exists {
				gramReference = GramReference{characterCodes: characterCodes[i : i+5]}
			}

			gramReference.total += 1

			gramReferences[gramReferenceKey] = gramReference
		}

		for _, gramReference := range gramReferences {
			if gramReference.total != 2 {
				continue
			}

			offset := gramReference.characterCodes[len(gramReference.characterCodes)-1] - int(lastLetter[0])

			if offset == int(lastLetter[0]) {
				continue
			}

			decoded := decode(gramReference.characterCodes, offset)
			if !re.MatchString(decoded) {
				continue
			}

			fmt.Println(decode(characterCodes, offset))
			break
		}
	}
}
