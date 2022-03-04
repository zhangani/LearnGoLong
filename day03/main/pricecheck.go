package main

import (
	"../shopping"
	"fmt"
)

func PriceCheck(itemId int) (float64, bool) {
	fmt.Println("main===", itemId)
	item := db.LoadItem(itemId)
	if item == nil {
		return 0, false
	}
	fmt.Println("PriceCheck===", item)
	return item.Price, true
}

func main() {
	fmt.Println(PriceCheck(1))
}
