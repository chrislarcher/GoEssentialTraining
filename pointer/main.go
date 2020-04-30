package main

import (
	"fmt"
)

// Point is a 2d point
type Point struct {
	X int
	Y int
}

// Move moves the point
func (p Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

func (p *Point) Move2(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
	p := Point{1, 2}
	p.Move(2, 3) // p is not updated because move doesn't use a pointer
	fmt.Printf("%+v\n", p)

	p2 := &Point{1, 2}
	p2.Move2(2, 3) // this will update p because it uses a pointer
	fmt.Printf("%+v\n", p2)
}