package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Address struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Cidade     string `json:"cidade"`
	UF         string `json:"uf"`
}

type ApiCepResponse struct {
	Code     string `json:"code"`
	State    string `json:"state"`
	City     string `json:"city"`
	District string `json:"district"`
	Address  string `json:"address"`
}

type ViaCepResponse struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
}

func ConsultaCepByApiCep(cep string) {
	cep = cep[:5] + "-" + cep[5:]
	url := fmt.Sprintf("https://cdn.apicep.com/file/apicep/%s.json", cep)
	res, err := http.Get(url)
	if err != nil {
		log.Panicf("error when consulting CEP by API CEP: %s", err)
	}
	if res.StatusCode != 200 {
		log.Panicf("the APICEP returned an error: %s", res.Status)
	}
	var apiCepResponse ApiCepResponse
	json.NewDecoder(res.Body).Decode(&apiCepResponse)
	log.Println(apiCepResponse)
}

func ConsultaCepByViaCep(cep string) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	res, err := http.Get(url)
	if err != nil {
		log.Panicf("error when consulting CEP by Via CEP: %s", err)
	}
	if res.StatusCode != 200 {
		log.Panicf("the VIA Cep API returned an error: %s", res.Status)
	}
	var viaCepResponse ViaCepResponse
	json.NewDecoder(res.Body).Decode(&viaCepResponse)
	log.Println(viaCepResponse)
}

func main() {
	ConsultaCepByApiCep("29172680")
	ConsultaCepByViaCep("29172680")
}
