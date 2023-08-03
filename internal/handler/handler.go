package handler

import (
	"one-lab-final/internal/config"
	"one-lab-final/internal/service"
)

type Handler struct {
	Services service.Service
	Config   *config.Config
}

func New(services service.Service, config *config.Config) *Handler {
	return &Handler{
		Services: services,
		Config:   config,
	}
}
