package migration

import "time"

type User struct {
	ID int `json:"primaryKey"`
	Fullname string `json:"fullname"`
	Address string `json:"address"`
	Email string `json:"email"`
	Password string `json:"password"`
	List []List `gorm:"foreignKey:UserID"`
}

type List struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Website   string    `json:"website"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAT time.Time `json:"updated_at"`
}


