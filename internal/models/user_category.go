package models

type UserCategory struct {
	UserID     uint `gorm:"primaryKey;column:user_id" json:"user_id"`
	CategoryID uint `gorm:"primaryKey;column:category_id" json:"category_id"`
}

func (UserCategory) TableName() string {
	return "user_categories"
}
