package main

import (
	"fmt"
)

// struct em GoLang
type Carro struct {
	Nome string
}

// método em GoLang
func (c Carro) andar() {
	fmt.Println(c.Nome, "capotou")
}

func main() {

	carro := Carro{
		Nome: "Chevetão de Drift",
	}

	carro.andar()

	res := func(x ...int) func() int {
		res := 0

		for _, v := range x {
			res += v
		}

		return func() int {
			return res * res
		}

	}
	fmt.Println(Mult(2, 3, 4, 5, 56, 7, 675, 67, 5, 67, 567, 5, 7, 56, 75, 76))
	fmt.Println(res(2, 3, 4, 5, 56, 7, 675, 67, 5, 67, 567, 5, 7, 56, 75, 76)())
}

// Possível passar um range para uma função variática
func Mult(x ...int) int {
	res := 0

	for _, v := range x {
		res += v
	}

	return res
}
