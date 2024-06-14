package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func main() {
	println("Digite o CEP: ")
	var cep string
	fmt.Scanln(&cep)
	c := resty.New()
	// viacep.com.br/ws/01001000/json/ 
	url := "https://viacep.com.br/ws/" + cep + "/json/"
	res, err := c.R().Get(url)
	if err != nil {
		panic(err)
	}
	println(res.String())
	fmt.Scanln()
}
