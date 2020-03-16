package main

import (
	"fmt"
	"reflect"
)

type triangle struct {
	base float64
	height float64
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}

type square struct {
	sideLength float64
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	txt := fmt.Sprintf("Area of %s: %.2f", reflect.TypeOf(s).Name(), s.getArea())
	fmt.Println(txt)
}

func main () {
	sq := square{
		sideLength: 5,
	}

	tr := triangle{
		base:   5,
		height: 8,
	}

	printArea(sq)
	printArea(tr)
}