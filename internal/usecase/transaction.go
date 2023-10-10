package usecase

import "template/internal/repository"

type TransactionHandler struct {
	g repository.GatheringRepository
	u repository.UserRepository
}

func NewTransactionsUsecase(g repository.GatheringRepository, u repository.UserRepository) GatheringUcase {
	return &TransactionHandler{g, u}
}
