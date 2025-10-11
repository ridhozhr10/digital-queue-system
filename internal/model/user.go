package model

type User struct {
	ID           int64
	Username     string
	Email        string
	PasswordHash string
	Role         string
}
