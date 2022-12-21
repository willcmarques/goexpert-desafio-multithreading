package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/willcmarques/goexpert-desafio-multithreading/dto"
)

type ViaCepResponse struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
}

func ConsultaCepByViaCep(cep string) dto.Address {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	c := http.Client{Timeout: time.Second}
	res, err := c.Get(url)
	if err != nil {
		log.Panicf("error when consulting CEP by Via CEP: %s", err)
	}
	if res.StatusCode != 200 {
		log.Panicf("the VIA Cep API returned an error: %s", res.Status)
	}
	var viaCepResponse ViaCepResponse
	json.NewDecoder(res.Body).Decode(&viaCepResponse)
	return convertViaCepResponseToAddress(&viaCepResponse)
}

func convertViaCepResponseToAddress(response *ViaCepResponse) dto.Address {
	return dto.Address{
		Cep:        response.Cep,
		Logradouro: response.Logradouro,
		Bairro:     response.Bairro,
		Cidade:     response.Localidade,
		UF:         response.UF,
		Api:        "Via CEP",
	}
}
