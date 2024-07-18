package math

var A string = "Testando"

// minusculo é função não exportada (private)
func soma(a int, b int) int {
	return a + b
}

// maiúsculo é função exportada (public)
func Soma(a int, b int) int {
	c := soma(a, b)
	return c
}
