package models

type Empresa struct {
	CNPJ        string `json:"cnpj`
	RazaoSocial string `json:"-"`
}
