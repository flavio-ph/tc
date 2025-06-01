package service

import (
	"fmt"
	"time"

	"agendamento-api/internal/models"
	"agendamento-api/internal/repository"
)

type AgendaService struct {
	repo *repository.AgendaRepository
}

func NewAgendaService(repo *repository.AgendaRepository) *AgendaService {
	return &AgendaService{repo: repo}
}

func (s *AgendaService) RequestAgenda(agenda *models.Agenda) error {
	if len(agenda.Empresa.CNPJ) != 14 {
		return fmt.Errorf("CNPJ inválido: deve conter 14 dígitos")
	}

	now := time.Now()
	loc := now.Location()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, loc)
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 18, 0, 0, 0, loc)

	agenda.Horario = time.Date(now.Year(), now.Month(), now.Day(), agenda.Horario.Hour(), agenda.Horario.Minute(), 0, 0, loc)

	if agenda.Horario.Before(startOfDay) || !agenda.Horario.Before(endOfDay) {
		return fmt.Errorf("horário de agendamento fora do período permitido (8h às 17h)")
	}

	if agenda.Horario.Minute() != 0 {
		return fmt.Errorf("horário de agendamento deve ser em horas exatas (ex: 10:00, 11:00)")
	}

	existingAgenda, err := s.repo.GetAgendaByHorarioAndCNPJ(agenda.Horario, agenda.Empresa.CNPJ)
	if err != nil {
		return fmt.Errorf("erro interno ao verificar agendamento existente: %w", err)
	}
	if existingAgenda != nil {
		return fmt.Errorf("já existe um agendamento para esta empresa neste horário")
	}

	if err := s.repo.CreateAgenda(agenda); err != nil {
		return fmt.Errorf("não foi possível agendar o horário: %w", err)
	}
	return nil
}

func (s *AgendaService) ListAgendas() ([]models.Agenda, error) {
	agendas, err := s.repo.ListaAgendas()
	if err != nil {
		return nil, fmt.Errorf("erro ao listar agendas do repositório: %w", err)
	}

	for i := range agendas {
		agendas[i].HorarioFormatado = agendas[i].Horario.Format("15:04")
	}

	return agendas, nil
}

type HorarioDisponibilidade struct {
	Inicio     time.Time `json:"inicio"`
	Fim        time.Time `json:"fim"`
	Disponivel bool      `json:"disponivel"`
}

func (s *AgendaService) CheckAvailability() ([]HorarioDisponibilidade, error) {
	var disponibilidades []HorarioDisponibilidade
	now := time.Now()
	loc := now.Location()

	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, loc)
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 18, 0, 0, 0, loc)

	bookedAgendas, err := s.repo.GetAgendasByPeriod(startOfDay, endOfDay)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar agendas para verificação de disponibilidade: %w", err)
	}

	bookedTimes := make(map[time.Time]bool)
	for _, agenda := range bookedAgendas {
		normalizedTime := time.Date(
			agenda.Horario.Year(), agenda.Horario.Month(), agenda.Horario.Day(),
			agenda.Horario.Hour(), 0, 0, 0, loc)
		bookedTimes[normalizedTime] = true
	}

	for t := startOfDay; t.Before(endOfDay); t = t.Add(1 * time.Hour) {
		isAvailable := !bookedTimes[t]
		disponibilidades = append(disponibilidades, HorarioDisponibilidade{
			Inicio:     t,
			Fim:        t.Add(1 * time.Hour),
			Disponivel: isAvailable,
		})
	}

	return disponibilidades, nil
}
