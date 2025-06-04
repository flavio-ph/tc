package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ReceitaWSResponse struct {
	RazaoSocial string `json:"nome"`
}

func ConsultaReceitaWS(cnpj string) (string, error) {
	url := fmt.Sprintf("https://receitaws.com.br/v1/cnpj/%s", cnpj)

	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("erro ao consultar ReceitaWS: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Resposta da ReceitaWS: %d", resp.StatusCode)
	}

	var data ReceitaWSResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", fmt.Errorf("erro ao decodificar JSON da ReceitaWS: %w", err)
	}

	return data.RazaoSocial, nil
}
