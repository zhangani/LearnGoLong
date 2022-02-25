package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}
type editObj struct {
	Name  string
	Power int
}

func main() {
	goku := &Saiyan{"Goku", 9001}
	Name := &editObj{"Goku", 9001}
	goku.Super()
	Name.edit()
	fmt.Println(goku.Power) // 将会打印出 19001
	fmt.Println(Name.Name)  // 将会打印出 19001
}

func (s *Saiyan) Super() {
	s.Power += 10000
}
func (s *editObj) edit() {
	s.Name = s.Name + "222222"
}
