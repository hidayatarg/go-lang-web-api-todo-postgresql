package domain

type UserRepo interface {
	GetByEmail(email string) (*User, error)
}

type DB struct {
	UserRepo UserRepo
}

type Domain struct {
	DB DB
}