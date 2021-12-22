package dto

// ProfesiUpdateDTO is a model that client use when update a profession
type ProfesiUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	Salary      int    `json:"salary" form:"salary" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	CategoryId  uint64 `json:"category_id,omitempty" form:"category_id,omitempty"`
}

// ProfesiCreateDTO is a model that client use when create a profession
type ProfesiCreateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	Salary      int    `json:"salary" form:"salary" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	CategoryId  uint64 `json:"category_id,omitempty" form:"category_id,omitempty"`
}
