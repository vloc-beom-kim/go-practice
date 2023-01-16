package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	yundream := person{name: "yundream", age: 25}
	jackson := person{name: "jackson", age: 31}
	fmt.Println(yundream)

	sp := &jackson
	sp.name = "Jackson json"
	sp.age = 48
	fmt.Println(sp.age)
	fmt.Println(jackson)
}
