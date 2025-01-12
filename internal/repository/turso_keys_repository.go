package repository

import (
	"database/sql"
	"errors"

	"github.com/c4miloarriagada/keys-be/internal/domain"
)

type tursoKeyRepository struct {
	db *sql.DB
}

func NewTursoKeysRepository(db *sql.DB) domain.KeysRepository {
	return &tursoKeyRepository{db: db}
}

func (r *tursoKeyRepository) GetByID(id int64) (*domain.Key, error) {
	row := r.db.QueryRow("SELECT id, name, description, pass, alias, created_at, valid_until FROM keys WHERE id = ?", id)

	var keys domain.Key
	if err := row.Scan(&keys.ID, &keys.Name, &keys.Description, &keys.Pass, &keys.Alias, &keys.CreatedAt, &keys.ValidUntil); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &keys, nil
}

func (r *tursoKeyRepository) Save(keys *domain.Key) error {
	_, err := r.db.Exec("INSERT INTO keys (name, description, pass, alias, created_at, valid_until) VALUES (?, ?, ?, ?, ?, ?)",
		keys.Name, keys.Description, keys.Pass, keys.Alias)

	if err != nil {
		return err
	}

	return nil
}
