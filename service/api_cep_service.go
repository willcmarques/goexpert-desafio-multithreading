package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/willcmarques/goexpert-desafio-multithreading/dto"
)

type ApiCepResponse struct {
	Code     string `json:"code"`
	State    string `json:"state"`
	City     string `json:"city"`
	District string `json:"district"`
	Address  string `json:"address"`
}

func ConsultaCepByApiCep(cep string) dto.Address {
	cep = cep[:5] + "-" + cep[5:]
	url := fmt.Sprintf("https://cdn.apicep.com/file/apicep/%s.json", cep)
	c := http.Client{Timeout: time.Second}
	res, err := c.Get(url)
	if err != nil {
		log.Panicf("error when consulting CEP by API CEP: %s", err)
	}
	if res.StatusCode != 200 {
		log.Panicf("the APICEP returned an error: %s", res.Status)
	}
	var apiCepResponse ApiCepResponse
	json.NewDecoder(res.Body).Decode(&apiCepResponse)
	return convertApiCepResponseToAddress(&apiCepResponse)
}

func convertApiCepResponseToAddress(response *ApiCepResponse) dto.Address {
	return dto.Address{
		Cep:        response.Code,
		Logradouro: response.Address,
		Bairro:     response.District,
		Cidade:     response.City,
		UF:         response.State,
		Api:        "API CEP",
	}
}
