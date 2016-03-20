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

const (
	PENNY       = 1
	NICKEL      = 5
	DIME        = 10
	QUARTER     = 25
	HALF_DOLLAR = 50
	ONE         = 100
	TWO         = 200
	FIVE        = 500
	TEN         = 1000
	TWENTY      = 2000
	FIFTY       = 5000
	ONE_HUNDRED = 10000
)

type DivisibleUnit struct {
	Name   string
	Amount int
}

var divisibleUnits = []DivisibleUnit{
	DivisibleUnit{"PENNY", 1},
	DivisibleUnit{"NICKEL", 5},
	DivisibleUnit{"DIME", 10},
	DivisibleUnit{"QUARTER", 25},
	DivisibleUnit{"HALF DOLLAR", 50},
	DivisibleUnit{"ONE", 100},
	DivisibleUnit{"TWO", 200},
	DivisibleUnit{"FIVE", 500},
	DivisibleUnit{"TEN", 1000},
	DivisibleUnit{"TWENTY", 2000},
	DivisibleUnit{"FIFTY", 5000},
	DivisibleUnit{"ONE HUNDRED", 10000},
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

		priceMultiplier := 100
		if strings.Contains(parts[0], ".") {
			priceMultiplier = 1
		}

		cashMultiplier := 100
		if strings.Contains(parts[1], ".") {
			cashMultiplier = 1
		}

		price, err := strconv.Atoi(strings.Replace(parts[0], ".", "", -1))
		if err != nil {
			log.Fatalf("> failed to read price: %s", err)
		}

		cash, err := strconv.Atoi(strings.Replace(parts[1], ".", "", -1))
		if err != nil {
			log.Fatalf("> failed to read cash: %s", err)
		}

		price *= priceMultiplier
		cash *= cashMultiplier

		if cash < price {
			fmt.Println("ERROR")
			continue
		} else if cash == price {
			fmt.Println("ZERO")
			continue
		}

		difference := cash - price

		units := []string{}

		for i := len(divisibleUnits) - 1; i >= 0; i-- {
			quotient := difference / divisibleUnits[i].Amount
			if quotient >= 1 {
				difference -= quotient * divisibleUnits[i].Amount
				for j := 0; j < quotient; j++ {
					units = append(units, divisibleUnits[i].Name)
				}
			}
		}

		fmt.Println(strings.Join(units, ","))
	}

}
