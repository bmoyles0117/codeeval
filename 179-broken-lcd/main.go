package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var digits = map[rune]int{
	'0': 0xFC, // 1111 1100
	'1': 0x60, // 0110 0000
	'2': 0xDA, // 1101 1010
	'3': 0xF2, // 1111 0010
	'4': 0x66, // 0110 0110
	'5': 0xB6, // 1011 0110
	'6': 0xBE, // 1011 1110
	'7': 0xE0, // 1110 0000
	'8': 0xFE, // 1111 1110
	'9': 0xE6, // 1110 0110
	'.': 0x01, // 0000 0001
}

func ledStringToInt(led string) int {
	val := 0
	for i := len(led) - 1; i >= 0; i-- {
		if led[i] == '1' {
			val += 1 << uint(len(led)-i-1)
		}
	}

	return val
}

func isCapableOfShowing(leds, ledsRequirement []int, sequenced bool) bool {
	// fmt.Println("is capable of showing?", leds, ledsRequirement)
	if len(ledsRequirement) <= 0 {
		// fmt.Println("REQUIREMENTS MET")
		return true
	}

	if len(leds) <= 0 {
		// fmt.Println("HERE", ledsRequirement)
		return false
	}

	if sequenced && (leds[0]&ledsRequirement[0] != ledsRequirement[0]) {
		return false
	}

	// if ledsRequirement[0]&leds[0] != ledsRequirement[0] {
	// 	return false
	// }

	for i, led := range leds {
		if led&ledsRequirement[0] == ledsRequirement[0] {
			if isCapableOfShowing(leds[i+1:], ledsRequirement[1:], true) {
				return true
			}
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

		parts := strings.Split(string(line), ";")
		ledStrings, target := strings.Split(parts[0], " "), parts[1]

		leds := []int{}
		for _, ledString := range ledStrings {
			leds = append(leds, ledStringToInt(ledString))
		}

		ledsRequirement := []int{}
		for _, char := range target {
			if char == '.' {
				ledsRequirement[len(ledsRequirement)-1] = ledsRequirement[len(ledsRequirement)-1] | 0x01
			} else {
				ledsRequirement = append(ledsRequirement, digits[char])
			}
		}

		if isCapableOfShowing(leds, ledsRequirement, false) {
			fmt.Println("1")
		} else {
			fmt.Println("0")
		}
	}
}
