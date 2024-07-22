package main

import "fmt"

type Anime struct {
	Genre string
}

func (a Anime) t_genre() {
	a.Genre = "CGDCT"
	fmt.Println("Func sem Ponteiro: ", a.Genre)
}

func (a *Anime) p_genre() {
	a.Genre = "CGDCT"
}

func main() {
	i := 10
	var p *int = &i

	// Imprime o endereço de p
	fmt.Println(&p)

	// Imprime o endereço o qual p aponta
	fmt.Println(p)

	// Imprime o valor da variável que p aponta
	fmt.Println(*p)

	// Altera o valor da variável que p aponta
	*p = 50
	fmt.Println("Após alteração \np: ", *p, "\ni: ", i)

	// Declaração rápida de ponteiro
	b := &i
	*b = 100
	fmt.Println("Após alteração \nb: ", *b, "\np: ", *p, "\ni: ", i)

	v := 20
	abc(&v)
	fmt.Println(v)

	// Ponteiros com Struct
	YuruCamp := Anime{
		Genre: "SoL",
	}

	YuruCamp.t_genre()
	fmt.Println("Pré ponteiro: ", YuruCamp.Genre)

	YuruCamp.p_genre()
	fmt.Println("Pós ponteiro: ", YuruCamp.Genre)
}

func abc(a *int) {
	*a = 200
}
