package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func getFibonacci(n int) int {
	previous := 0
	value := 1

	for i := 1; i < n; i++ {
		p := value
		value += previous
		previous = p
	}

	return value
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

		n, _ := strconv.Atoi(string(line))

		if n == 0 {
			fmt.Println("0")
		} else {
			fmt.Println(getFibonacci(n))
		}
	}
}
