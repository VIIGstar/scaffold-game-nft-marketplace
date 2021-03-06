package user_handlers

import (
	"logur.dev/logur"
	"scaffold-game-nft-marketplace/internal/repository"
)

type userHandler struct {
	logger logur.LoggerFacade
	repo   repository.Registry
}

func New(logger logur.LoggerFacade, registry repository.Registry) userHandler {
	return userHandler{
		logger: logger,
		repo:   registry,
	}
}
