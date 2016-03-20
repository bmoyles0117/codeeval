package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	totalRaw, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("> failed to grab first value: %s", err)
	}

	total, err := strconv.Atoi(strings.Trim(totalRaw, "\n"))
	if err != nil {
		log.Fatalf("> failed to convert first value: %s", err)
	}

	values := make([]string, total)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("> failed to read line: %s", err)
		}

		text := string(line)

		for i, currentText := range values {
			if len(text) > len(currentText) {
				for x := len(values); x > i && x > 1; x-- {
					values[x-1] = values[x-2]
				}

				values[i] = text
				break
			}
		}
	}

	fmt.Println(strings.Join(values, "\n"))
}
