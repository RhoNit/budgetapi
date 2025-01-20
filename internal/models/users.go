package models

type User struct {
	BaseModel
	FirstName      string `gorm:"type:varchar(200)" json:"first_name"`
	LastName       string `gorm:"type:varchar(200)" json:"last_name"`
	Gender         string `gorm:"type:varchar(50)" json:"gender"`
	Email          string `gorm:"type:varchar(100);not null;unique" json:"email"`
	HashedPassword string `gorm:"type:varchar(250);not null" json:"-"`
}

func (u *User) TableName() string {
	return "users"
}
