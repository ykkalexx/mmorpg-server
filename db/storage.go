package db

import (
	"database/sql"

	"github.com/mmorpg-server/types"
)

type Storage interface {
	CreatePlayer(*types.Player) error
	GetPlayerByID(int) (*types.Player, error)
	GetAllPlayers() ([]*types.Player, error)
	DeletePlayer(int) error	
	UpdatePlayer(int) error
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostGresStorage() (*PostgresStorage, error) {
	connStr := "user=alex dbname=mmorpg password=alex sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStorage{
		db: db,
	}, nil
}

func (s *PostgresStorage) CreatePlayer(*types.Player) error {
	return nil
}

func (s *PostgresStorage) GetPlayerByID(id int) (*types.Player, error) {
	return nil, nil
}

func (s *PostgresStorage) GetAllPlayers() ([]*types.Player, error) {
    return nil, nil
}

func (s *PostgresStorage) DeletePlayer(id int) error {
	return nil
}

func (s *PostgresStorage) UpdatePlayer(id int) error {
	return nil
}

