package models

type Category struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	Name        string       `gorm:"type:varchar(255)" json:"name"`
	Professions []Profession `gorm:"foreignKey:CategoryID"`
}
