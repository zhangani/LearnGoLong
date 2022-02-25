package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

type Saiyan struct {
	*Person
	Power int
}

// func main() {
// 	// 使用它
// 	goku := &Saiyan{
// 		Person: &Person{"Goku"},
// 		Power:  9001,
// 	}
// 	goku.Introduce()
// }

func main() {
	goku := &Saiyan{
		Person: &Person{"Goku111"},
	}
	fmt.Println(goku.Name)
	fmt.Println(goku.Person.Name)
}
