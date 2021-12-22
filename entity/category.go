package entity

type Category struct {
	ID          uint64       `gorm:"primary_key:auto_increment" json:"id"`
	Name        string       `gorm:"type:varchar(255)" json:"name"`
	Professions []Profession `gorm:"foreignKey:CategoryId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"professions"`
}
