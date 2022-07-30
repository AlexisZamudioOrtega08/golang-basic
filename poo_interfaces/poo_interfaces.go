package main

import (
	"fmt"
)

type shapes interface {
	getArea() float32
}

// an interface is a collection of method that are shared by different types of objects

type circle struct {
	radius float32
}

type rectangle struct {
	base   float32
	height float32
}

func (c circle) getArea() float32 {
	return 3.14 * c.radius * c.radius
}

func (s rectangle) getArea() float32 {
	return float32(s.base * s.height)
}

func main() {
	c := circle{5}
	s := rectangle{5, 10}
	shapes := []shapes{c, s}
	for _, shape := range shapes {
		fmt.Println(shape.getArea())
	}
}
