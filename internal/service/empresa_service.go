package service

import (
	"fmt"
	"regexp"
	"strconv"
)

func ValidateCNPJ(cnpj string) (string, error) {

	re := regexp.MustCompile(`[^0-9]`)
	cnpj = re.ReplaceAllString(cnpj, "")

	if len(cnpj) != 14 {
		return "", fmt.Errorf("CNPJ inválido: deve conter 14 dígitos")
	}

	repetidos := []string{
		"00000000000000", "11111111111111", "22222222222222",
		"33333333333333", "44444444444444", "55555555555555",
		"66666666666666", "77777777777777", "88888888888888",
		"99999999999999",
	}
	for _, rep := range repetidos {
		if cnpj == rep {
			return "", fmt.Errorf("CNPJ inválido: todos os dígitos são iguais")
		}
	}

	pesos1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	pesos2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	soma := 0
	for i := 0; i < 12; i++ {
		num, _ := strconv.Atoi(string(cnpj[i]))
		soma += num * pesos1[i]
	}
	resto := soma % 11
	dig1 := 0
	if resto >= 2 {
		dig1 = 11 - resto
	}

	soma = 0
	for i := 0; i < 13; i++ {
		num, _ := strconv.Atoi(string(cnpj[i]))
		soma += num * pesos2[i]
	}
	resto = soma % 11
	dig2 := 0
	if resto >= 2 {
		dig2 = 11 - resto
	}

	if int(cnpj[12]-'0') != dig1 || int(cnpj[13]-'0') != dig2 {
		return "", fmt.Errorf("CNPJ inválido: dígitos verificadores incorretos")
	}

	return cnpj, nil
}
