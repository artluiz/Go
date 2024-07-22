package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Anime struct {
	Name  string `json:"nome"`
	Genre string `json:"gênero"`
	Eps   int    `json:"episódios"`
}

// métodos são criados ao criar função que recebe uma struct
func (a Anime) Exibe() {
	fmt.Println("Exibindo anime pelo método: ", a)
}

type aoty struct {
	Anime
	Studio string `json:"estúdio"`
}

func main() {

	a := Anime{
		Name:  "Yuru Camp",
		Genre: "SoL",
		Eps:   12,
	}
	fmt.Println(a)

	a2 := aoty{Anime{"Frieren", "Fantasia", 26}, "MadHouse"}

	fmt.Println(a2)

	a2.Exibe()

	animeJson, err := json.Marshal(a2)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(string(animeJson))

	jsonA3 := `{"nome":"Bocchi The Rock","gênero":"Música","episódios":12,"estúdio":"CloverWorks"}`
	a3 := aoty{}

	// como é uma função, deve ser passado endereço de a3 para acessar o espaço de memória real
	json.Unmarshal([]byte(jsonA3), &a3)
	fmt.Println(a3)

}
