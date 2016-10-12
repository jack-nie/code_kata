package code_kata

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type SpecialRule struct {
	quantity int
	price    int
}

type Rule struct {
	name        string
	price       int
	specialRule SpecialRule
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func isNotNumeric(str string) bool {
	_, err := strconv.ParseInt(str, 10, 64)
	return err != nil
}

func analyzePricingRules(file_path string) []Rule {
	dat, err := os.Open(file_path)

	check(err)

	scanner := bufio.NewScanner(dat)
	var rule []Rule
	var r Rule
	var specialRule SpecialRule

	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())

		if (len(arr) != 2 && len(arr) != 5) || arr[0] == "Price" {
			continue
		}

		r.name = arr[0]
		r.price, _ = strconv.Atoi(arr[1])

		if len(arr) == 5 {
			specialRule.quantity, _ = strconv.Atoi(arr[2])
			specialRule.price, _ = strconv.Atoi(arr[4])
		}

		r.specialRule = specialRule
		rule = append(rule, r)
	}
	return rule
}
