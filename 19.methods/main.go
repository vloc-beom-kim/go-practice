package main

import "fmt"

type rect struct {
	width, height int
}

func (self *rect) area() int {
	return self.width * self.height
}

func (self rect) perim() int {
	return 2*self.width + 2*self.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("Area", r.area())
	fmt.Println("perim", r.perim())

	rp := &r
	fmt.Println("Area", rp.area())
	fmt.Println("perim", rp.perim())
}
