package services

import (
	"errors"

	"github.com/RhoNit/budgetapi/cmd/api/requests"
	"github.com/RhoNit/budgetapi/common"
	"github.com/RhoNit/budgetapi/internal/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (u *UserService) RegisterUser(userRequest *requests.RegisterUserRequest) (*models.User, error) {
	// hash the password
	hashedPasswd, err := common.HashedPassword(userRequest.Password)
	if err != nil {
		return nil, errors.New("couldn't able to hash the password")
	}

	// map RegisterUserRequest type into models.User type
	user := models.User{
		FirstName:      userRequest.FirstName,
		LastName:       userRequest.LastName,
		Email:          userRequest.Email,
		HashedPassword: hashedPasswd,
	}

	// Save to the database
	if result := u.db.Create(&user); result.Error != nil {
		return nil, errors.New("saving to database error")
	}

	return &user, nil
}

func (u *UserService) LoginUser(userRequest *requests.LoginUserRequest) error {
	// compare the request-body password with hashed one in the DB

	return nil
}

func (u *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user *models.User
	result := u.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
