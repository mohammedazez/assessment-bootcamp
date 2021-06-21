package migration

import "time"

type User struct {
	ID       int `json:"primaryKey"`
	Fullname string
	Address  string
	Email    string
	Password string
	List     []List `gorm:"foreignKey:UserID"`
}

type List struct {
	ID        int `json:"primaryKey"`
	UserID    int
	Website   string
	Password  string
	CreatedAt time.Time
	UpdatedAT time.Time
}
