package domain

type User struct {
	ID       int64
	Name     string
	Email    string
	LastName string
	Password string
}

type UserRepository interface {
	GetByID(id int) (*User, error)
	Save(user *User) error
	GetAll() ([]User, error)
}
