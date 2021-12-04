package main

import "fmt"

// define interface
type geometry interface {
	area() float32
	perimeter() float32
}

// define struct
type rectangle struct {
	width  float32
	height float32
}

type circle struct {
	radius float32
}

// define function for specific struct
func (r rectangle) area() float32 {
	return r.height * r.width
}

func (r rectangle) perimeter() float32 {
	return 2 * (r.height + r.width)
}

func (c circle) area() float32 {
	return 3.14 * c.radius * c.radius / 2
}

func (c circle) perimeter() float32 {
	return 2 * 3.14 * c.radius
}

// here we can use the interface like this
func measure(m geometry) {
	fmt.Println("Area is: ", m.area())
	fmt.Println("Perimeter is: ", m.perimeter())
}
func main() {

	myRect := rectangle{
		width:  2.0,
		height: 3.0,
	}
	myCircle := circle{
		radius: 3.0,
	}
	measure(myRect)
	measure(myCircle)
}
