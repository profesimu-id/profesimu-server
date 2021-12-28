package models

type Profession struct {
	ID          uint64 `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Salary      int
	Description string `gorm:"type:varchar(255)" json:"description"`
	CategoryID  uint
}
