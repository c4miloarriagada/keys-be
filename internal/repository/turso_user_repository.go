package repository

import (
	"database/sql"
	"errors"

	"github.com/c4miloarriagada/keys-be/internal/domain"
)

type tursoUserRepository struct {
	db *sql.DB
}

func NewTursoUserRepository(db *sql.DB) domain.UserRepository {
	return &tursoUserRepository{db: db}
}

func (r *tursoUserRepository) GetByID(id int) (*domain.User, error) {
	row := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)

	var user domain.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *tursoUserRepository) Save(user *domain.User) error {
	_, err := r.db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	return err
}

func (r *tursoUserRepository) GetAll() ([]domain.User, error) {
	rows, err := r.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
