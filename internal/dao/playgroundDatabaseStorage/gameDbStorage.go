package gameStorage

import postgres "github.com/go-pg/pg/v10"

type dbStorage struct {
	postgres.DB
}

//constructor for storage
func NewdbStorage(db *postgres.DB) *dbStorage {
	return &dbStorage{*db}
}

//TODO - Implement database storage to keep track of the player's score vs AI
