package repository

import (
	"database/sql"
	"fmt"
	"time"

	"agendamento-api/internal/models"
)

type AgendaRepository struct {
	db *sql.DB
}

func NewAgendaRepository(db *sql.DB) *AgendaRepository {
	return &AgendaRepository{db: db}
}

func (r *AgendaRepository) CreateAgenda(agenda *models.Agenda) error {
	query := `INSERT INTO agendas (empresa_cnpj, horario) VALUES (?, ?)`
	_, err := r.db.Exec(query, agenda.Empresa.CNPJ, agenda.Horario)
	if err != nil {
		return fmt.Errorf("falha ao criar agendamento no DB: %w", err)
	}
	return nil
}

func (r *AgendaRepository) ListaAgendas() ([]models.Agenda, error) {
	query := `SELECT id, empresa_cnpj, horario, created_at FROM agendas ORDER BY horario ASC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("falha ao listar agendas do DB: %w", err)
	}
	defer rows.Close()

	var agendas []models.Agenda
	for rows.Next() {
		var agenda models.Agenda
		err := rows.Scan(&agenda.ID, &agenda.Empresa.CNPJ, &agenda.Horario, &agenda.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("falha ao escanear linha da agenda: %w", err)
		}
		agendas = append(agendas, agenda)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante a iteração das agendas: %w", err)
	}

	return agendas, nil
}

func (r *AgendaRepository) GetAgendaByHorarioAndCNPJ(horario time.Time, cnpj string) (*models.Agenda, error) {
	query := `SELECT id FROM agendas WHERE horario = ? AND empresa_cnpj = ?`
	var id int
	err := r.db.QueryRow(query, horario, cnpj).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("falha ao verificar agendamento por horário e CNPJ: %w", err)
	}
	return &models.Agenda{ID: id}, nil
}

func (r *AgendaRepository) GetAgendasByPeriod(start, end time.Time) ([]models.Agenda, error) {
	query := `SELECT id, empresa_cnpj, horario FROM agendas WHERE horario >= ? AND horario < ? ORDER BY horario ASC`
	rows, err := r.db.Query(query, start, end)
	if err != nil {
		return nil, fmt.Errorf("falha ao obter agendas por período: %w", err)
	}
	defer rows.Close()

	var agendas []models.Agenda
	for rows.Next() {
		var agenda models.Agenda
		err := rows.Scan(&agenda.ID, &agenda.Empresa.CNPJ, &agenda.Horario)
		if err != nil {
			return nil, fmt.Errorf("falha ao escanear linha da agenda por período: %w", err)
		}
		agendas = append(agendas, agenda)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante iteração das agendas por período: %w", err)
	}

	return agendas, nil
}
