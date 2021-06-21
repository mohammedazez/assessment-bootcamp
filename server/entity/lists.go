package entity

import "time"

type List struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Website string `json:"website"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAT time.Time `json:"updated_at"`
}

