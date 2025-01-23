package requests

type CreateBudgetRequest struct {
	CategoryIDs []uint  `json:"category_ids" validate:"required,dive,min=1"`
	Amount      float64 `json:"amount" validate:"required,numeric,min=1"`
	Date        string  `json:"date,omitempty" validate:"omitempty,datetime=2006-01-01"`
	Title       string  `json:"title" validate:"required,min=2,max=200"`
	Description *string `json:"description" validate:"omitempty,min=5,max=500"`
}

type UpdateBudgetRequest struct {
	CategoryIDs []uint  `json:"category_ids" validate:"omitempty,dive,min=1"`
	Amount      float64 `json:"amount" validate:"omitempty,numeric,min=1"`
	Date        string  `json:"date,omitempty" validate:"omitempty,datetime=2006-01-01"`
	Title       string  `json:"title" validate:"omitempty,min=2,max=200"`
	Description *string `json:"description" validate:"omitempty,min=5,max=500"`
}
