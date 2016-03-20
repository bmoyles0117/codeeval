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

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("> failed to read line: %s", err)
		}

		parts := strings.Split(string(line), " : ")

		values := strings.Split(parts[0], " ")

		for _, swapRaw := range strings.Split(parts[1], ", ") {
			swapParts := strings.Split(swapRaw, "-")

			position1, _ := strconv.Atoi(swapParts[0])
			position2, _ := strconv.Atoi(swapParts[1])

			v := values[position1]
			values[position1] = values[position2]
			values[position2] = v
		}

		fmt.Println(strings.Join(values, " "))
	}
}
