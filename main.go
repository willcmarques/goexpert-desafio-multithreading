package main

import (
	"log"
	"os"

	"github.com/willcmarques/goexpert-desafio-multithreading/dto"
	"github.com/willcmarques/goexpert-desafio-multithreading/service"
)

func main() {
	ch1 := make(chan dto.Address)
	ch2 := make(chan dto.Address)
	cep := "29172680"
	go func() {
		address := service.ConsultaCepByApiCep(cep)
		ch1 <- address
	}()
	go func() {
		ch2 <- service.ConsultaCepByViaCep(cep)
	}()

	for {
		select { //Inseri um exit para não dar tempo de ver o resultado no log.
		case address := <-ch1:
			log.Println(address)
			os.Exit(0)
		case address := <-ch2:
			log.Print(address)
			os.Exit(0)
		default:
			log.Println("Timeout")
		}
	}
}
