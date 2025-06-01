package models

import "time"

type Agenda struct {
	HorarioFormatado string  `json:"horario"`
	Empresa          Empresa `json:"empresa"`

	ID          int       `json:"-"`
	CreatedAt   time.Time `json:"-"`
	Horario     time.Time `json:"-"`
	EmpresaCNPJ string    `json:"-" db:"empresa_cnpj"`
}
