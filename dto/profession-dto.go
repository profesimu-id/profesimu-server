package dto

type ProfesiUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	Salary      int    `json:"salary" form:"salary" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	CategoryId  uint64 `json:"category_id,omitempty" form:"category_id,omitempty"`
}

type ProfesiCreateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	Salary      int    `json:"salary" form:"salary" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	CategoryId  uint64 `json:"category_id,omitempty" form:"category_id,omitempty"`
}
