package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"agendamento-api/internal/models"
	"agendamento-api/internal/service"
	"agendamento-api/pkg/httputils"
)

type AgendaHandler struct {
	service *service.AgendaService
}

func NewAgendaHandler(s *service.AgendaService) *AgendaHandler {
	return &AgendaHandler{service: s}
}

func (h *AgendaHandler) RequestAgendaHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Empresa struct {
			CNPJ string `json:"cnpj"`
		} `json:"empresa"`
		Horario string `json:"horario"` // Recebe o horário como string "HH:MM"
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		httputils.RespondWithError(w, http.StatusBadRequest, "Requisição inválida: formato JSON incorreto.")
		return
	}

	if req.Empresa.CNPJ == "" || len(req.Empresa.CNPJ) != 14 {
		httputils.RespondWithError(w, http.StatusBadRequest, "CNPJ inválido: deve conter 14 dígitos e não pode estar vazio.")
		return
	}

	parsedTime, err := time.Parse("15:04", req.Horario)
	if err != nil {
		httputils.RespondWithError(w, http.StatusBadRequest, "Formato de horário inválido. Use HH:MM (ex: 10:00).")
		return
	}

	agenda := &models.Agenda{
		Empresa: models.Empresa{
			CNPJ: req.Empresa.CNPJ,
		},
		Horario: parsedTime,
	}
	err = h.service.RequestAgenda(agenda)
	if err != nil {

		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Agendamento solicitado com sucesso!"})
}

func (h *AgendaHandler) ListAgendasHandler(w http.ResponseWriter, r *http.Request) {
	agendas, err := h.service.ListAgendas()
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httputils.RespondWithJSON(w, http.StatusOK, agendas)
}

func (h *AgendaHandler) CheckAvailabilityHandler(w http.ResponseWriter, r *http.Request) {
	disponibilidades, err := h.service.CheckAvailability()
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httputils.RespondWithJSON(w, http.StatusOK, disponibilidades)
}
