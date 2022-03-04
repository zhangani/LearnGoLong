package db

import "fmt"

type Item struct {
	Price float64
}

func LoadItem(id int) *Item {
	fmt.Println("LoadItem===", 9.001)
	return &Item{
		Price: 9.001,
	}
}
