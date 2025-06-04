package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"agendamento-api/database"
	"agendamento-api/internal/handler"
	"agendamento-api/internal/repository"
	"agendamento-api/internal/service"
)

func main() {
	database.InitDB()
	defer database.CloseDB()

	agendaRepo := repository.NewAgendaRepository(database.DB)
	agendaService := service.NewAgendaService(agendaRepo /*, receitaWSClient*/)
	agendaHandler := handler.NewAgendaHandler(agendaService)

	r := mux.NewRouter()
	r.HandleFunc("/agendas", agendaHandler.RequestAgendaHandler).Methods("POST")
	r.HandleFunc("/agendas", agendaHandler.ListAgendasHandler).Methods("GET")
	r.HandleFunc("/agendas:disponibilidade", agendaHandler.CheckAvailabilityHandler).Methods("GET")

	port := ":8080"
	log.Printf("🚀 Servidor da API de Agendamento iniciado na porta %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
