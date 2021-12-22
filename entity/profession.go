package entity

type Profession struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Salary      int
	Description string   `gorm:"type:varchar(255)" json:"description"`
	CategoryId  uint64   `gorm:"not null" json:"-"`
	Category    Category `gorm:"foreignKey:CategoryRefer;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"category"`
}
