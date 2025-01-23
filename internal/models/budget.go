package models

import "time"

type Budget struct {
	BaseModel
	Title       string      `gorm:"index;type:varchar(255);not null" json:"title"`
	Slug        string      `gorm:"index;type:varchar(255);not null;uniqueIndex:unique_idx_user_id_slug_year_month" json:"slug"`
	Description *string     `gorm:"type:text" json:"description"`
	UserID      uint        `gorm:"not null;column:user_id;uniqueIndex:unique_idx_user_id_slug_year_month" json:"user_id"`
	Amount      float64     `gorm:"type:decimal(10,2);not null" json:"amount"`
	Categories  []*Category `gorm:"constraint:OnDelete:CASCADE;many2many:budget_categories;" json:"categories"`
	Date        time.Time   `gorm:"type:timestamptz;not null" json:"date"`
	Month       uint8       `gorm:"type:SMALLINT;check:month >= 1 AND month <= 12;not null;index:idx_year_month;uniqueIndex:unique_idx_user_id_slug_year_month" json:"month"`
	Year        uint16      `gorm:"type:INT;check:year >= 0;not null;index:idx_year_month;uniqueIndex:unique_idx_user_id_slug_year_month" json:"year"`
}

// slug, year, month, user_id
func (Budget) TableName() string {
	return "budgets"
}
