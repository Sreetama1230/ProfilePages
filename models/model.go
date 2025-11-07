package models

type User struct {
	ID      uint    `json:"id" gorm:"primaryKey"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Profile Profile `json:"profile" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Profile struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	UserID  uint   `json:"-"`
	Bio     string `json:"bio" gorm:"primaryKey"`
	Twitter string `json:"twitter"`
}
