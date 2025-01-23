package requests

type CategoryRequest struct {
	Name     string `json:"name" validate:"required"`
	IsCustom bool   `default:"false" json:"is_custom"`
}

type AssociateUserToCategoriesRequest struct {
	CategoryIDs []uint `json:"category_ids" validate:"required"`
}
