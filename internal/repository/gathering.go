package repository

import "database/sql"

type GatheringHandler struct {
	db *sql.DB
}

func NewGatheringRepository(db *sql.DB) GatheringRepository {
	return &GatheringHandler{db}
}
