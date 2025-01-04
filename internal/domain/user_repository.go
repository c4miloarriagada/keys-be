package domain

type UserRepository interface {
	GetByID(id int) (*User, error)
	Save(user *User) error
	GetAll() ([]User, error)
}
