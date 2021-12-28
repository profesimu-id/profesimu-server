package models

type User struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Email       string `gorm:"type:varchar(255),unique" json:"email"`
	Password    string `gorm:"->;<-" json:"password"`
	Role        string `gorm:"type:varchar(255)" json:"role"`
	Token       string
	Profile     Profile      `gorm:"foreignKey:ID"`
	Professions []Profession `gorm:"many2many:bookmarks;"`
}

type Profile struct {
	ID      uint `gorm:"primaryKey" json:"id"`
	Phone   string
	Address string
}
