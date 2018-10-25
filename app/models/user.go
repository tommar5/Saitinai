package app

import (
	"github.com/fxtlabs/date"
)

type User struct {
	ID        int       `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Birthday  date.Date `json:"birthday"`
	Phone     string    `json:"phone"`
	City      string    `json:"city"`
	Address   string    `json:"address"`
	Role      int       `json:"role"`
}

type Users []User
