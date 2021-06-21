package entity

type User struct {
	ID       int    `json:"primaryKey"`
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
	List     []List `gorm:"foreignKey:UserID"`
}
