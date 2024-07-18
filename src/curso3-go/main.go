package main

import (
	"errors"
	"fmt"
	"log"
	// "net/http"
)

func main() {
	// res, err := http.Get("https://www.youtube.com/watch?v=XsKiSaBesns&ab_channel=LucasMontano")

	res, err := Mult(2, 20)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(res)
}

func Mult(a int, b int) (int, error) {
	res := a * b

	if res > 20 {
		return 0, errors.New("Total acima de 20")
	}
	return res, nil
}
