package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/c4miloarriagada/keys-be/internal/domain"
)

type tursoKeyRepository struct {
	db *sql.DB
}

func NewTursoKeysRepository(db *sql.DB) domain.KeyRepository {
	return &tursoKeyRepository{db: db}
}

func (r *tursoKeyRepository) GetByID(id int64) (*domain.Key, error) {
	row := r.db.QueryRow("SELECT id, name, description, pass, alias, created_at, valid_until FROM keys WHERE id = ?", id)

	var keys domain.Key
	if err := row.Scan(&keys.ID, &keys.Name, &keys.Description, &keys.Pass, &keys.Alias, &keys.CreatedAt, &keys.ValidUntil); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		log.Printf("failed to get the key: %v", err)
		return nil, errors.New("failed to get the key")
	}

	return &keys, nil
}

func (r *tursoKeyRepository) Save(keys *domain.Key) error {
	_, err := r.db.Exec("INSERT INTO keys (name, description, pass, alias, created_at, valid_until) VALUES (?, ?, ?, ?, ?, ?)",
		keys.Name, keys.Description, keys.Pass, keys.Alias)

	if err != nil {
		log.Printf("failed to save the key: %v", err)
		return errors.New("failed to save the key")
	}

	return nil
}
