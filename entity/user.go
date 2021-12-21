package entity

type User struct {
	ID          uint16       `gorm:"primary_key:auto_increment" json:"id"`
	Name        string       `gorm:"type:varchar(255)" json:"name"`
	Email       string       `gorm:"uniqueIndex;type:varchar(255)" json:"email" binding:"required"`
	Password    string       `gorm:"->;<-;not null" json:"-" binding:"required"`
	Role        string       `gorm:"type:varchar(255)" json:"role"`
	Token       string       `gorm:"-" json:"token,omitempty"`
	Professions []Profession `gorm:"many2many:favorites;"`
}
