package model

import (
	"google.golang.org/genproto/googleapis/type/date"
)

type FullUser struct {
	ID           string
	Title        string
	FirstName    string
	LastName     string
	PasswordHash string
	Gender       string
	Email        string
	DateOfBirth  date.Date
	RegisterDate date.Date
	PhoneNumber  string
}
