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

type Item struct {
	name     string
	quantity int
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

func analyzePricingRules(file_path string) map[string]Rule {
	var rule map[string]Rule
	var r Rule
	var specialRule SpecialRule

	rule = make(map[string]Rule)
	dat, err := os.Open(file_path)
	check(err)
	scanner := bufio.NewScanner(dat)
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
		rule[arr[0]] = r
	}
	return rule
}

func (item *Item) isSpecialRuleMeeted(rule map[string]Rule) bool {
	specialRule := rule[item.name].specialRule
	if &specialRule != nil {
		return item.quantity >= specialRule.quantity
	}
	return false
}

func addToCart(item_name string, cart *map[string]Item) bool {

	item := (*cart)[item_name]
	if item.quantity != 0 {
		item.quantity += 1
		(*cart)[item_name] = item
	} else {
		item.name = item_name
		item.quantity = 1
		(*cart)[item_name] = item
	}
	return true
}

func (item *Item) calculateSpecialPrice(rule map[string]Rule) int {
	var sum int
	var specialCouples int
	var remainQuantity int
	sum = 0

	if item.isSpecialRuleMeeted(rule) {
		specialCouples = item.quantity / rule[item.name].specialRule.quantity
		remainQuantity = item.quantity % rule[item.name].specialRule.quantity
		sum = specialCouples*rule[item.name].specialRule.price + remainQuantity*rule[item.name].price
	} else {
		sum = item.quantity * rule[item.name].price
	}
	return sum
}

func calculateTotalPrice(item_names []string, file_path string) int {
	var cart map[string]Item
	var sum int
	sum = 0
	rule := analyzePricingRules(file_path)
	cart = make(map[string]Item)

	for _, item := range item_names {
		addToCart(item, &cart)
	}

	for _, value := range cart {
		sum += value.calculateSpecialPrice(rule)
	}
	return sum
}
