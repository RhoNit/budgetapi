package models

type Category struct {
	BaseModel
	Name     string `gorm:"unique;type:varchar(200);not null" json:"name"`
	Slug     string `gorm:"type:varchar(200);unique;not null" json:"slug"`
	IsCustom bool   `gorm:"type:bool;default:false" json:"is_custom"`
}

func (Category) TableName() string {
	return "categories"
}
