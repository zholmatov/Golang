package main

import "fmt"

type describer interface {
	describe() string
}

type base struct {
	num int
}

// container struct embeds base
type container struct {
	base
	str string
}

func (b base) describe() string {
	return fmt.Sprintf("Num is %v", b.num)
}

func main() {

	co := container{
		base: base{num: 1},
		str:  "Something",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println(co.describe())

	// declaring describer interface to co container
	var d describer = co
	fmt.Println("describer:", d.describe())
}
