package main

import "fmt"

type List struct {
	Num int
}

func main() {
	lookup := make(map[string]int)

	lookup["goku"] = 9001

	power, exists := lookup["value"], true

	fmt.Println(power, exists)

	total := len(lookup)

	fmt.Println("total====", total)
	// returns 1
	index := len(lookup)

	// has no return, can be called on a non-existing key
	delete(lookup, "goku")
	fmt.Println(lookup)

	fmt.Println("index====", index)

}
