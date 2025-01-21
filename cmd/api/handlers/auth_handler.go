package handlers

import (
	"errors"

	"github.com/RhoNit/budgetapi/cmd/api/requests"
	"github.com/RhoNit/budgetapi/cmd/api/services"
	"github.com/RhoNit/budgetapi/common"
	"github.com/RhoNit/budgetapi/internal/mailer"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (h *Handler) RegisterUserHandler(c echo.Context) error {
	request := new(requests.RegisterUserRequest)

	// bind the request body
	if err := (&echo.DefaultBinder{}).BindBody(c, request); err != nil {
		h.Logger.Error(err)
		return common.SendBadRequestResponse(c, err.Error())

	}

	// var validate *validator.Validate
	// var validationErrors []*validation.ValidationError
	validationErrors := h.ValidateRequestBody(c, *request)

	if validationErrors != nil {
		h.Logger.Error(validationErrors)
		return common.SendFailedValidationResponse(c, validationErrors)
	}

	// check if email already exists
	userService := services.NewUserService(h.DB)
	_, err := userService.GetUserByEmail(request.Email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return common.SendBadRequestResponse(c, "Email already exists")
	}

	// create the user
	// var registeredUser *models.User
	registeredUser, err := userService.RegisterUser(request)
	if err != nil {
		// return common.SendBadRequestResponse(c, "incorrect request body.. error occurred while unmarshalling request body into golang struct type")
		return common.SendInternalServerErrorResponse(c, err.Error())
	}

	mailData := mailer.EmailData{
		Subject: "🛠️📨📥Testing Email Delivery",
		MetaData: struct {
			FirstName      string
			GetStartedLink string
			SupportLink    string
		}{
			FirstName:      registeredUser.FirstName,
			GetStartedLink: "#",
			SupportLink:    "#",
		},
	}
	// send welcome note to the user
	err = h.Mailer.Send(request.Email, "hello.html", mailData)
	if err != nil {
		h.Logger.Error(err)
		return err
	}

	// c.Logger().Info(request)

	return common.SendSuccessResponse(c, "User Registration successful!", registeredUser)
}

func (h *Handler) LoginUserHandler(c echo.Context) error {
	request := new(requests.LoginUserRequest)
	// bind payload with golang-struct type
	if err := (&echo.DefaultBinder{}).BindBody(c, request); err != nil {
		h.Logger.Error(err)
		return common.SendBadRequestResponse(c, err.Error())
	}

	// check validation of request body i.e. payload
	validationErrors := h.ValidateRequestBody(c, *request)

	if validationErrors != nil {
		h.Logger.Error(validationErrors)
		return common.SendFailedValidationResponse(c, validationErrors)
	}

	// verify user email id already exists or not
	userService := services.NewUserService(h.DB)
	retrievedUser, err := userService.GetUserByEmail(request.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		h.Logger.Error(err)
		return common.SendBadRequestResponse(c, "Invalid email or password")
	}

	// compare password hash with plain password
	result := common.CheckPasswordHash(request.Password, retrievedUser.HashedPassword)
	if !result {
		return common.SendBadRequestResponse(c, "Invalid email or password")
	}

	// send response with access token
	accessToken, refreshToken, err := common.GenerateJWTTokens(retrievedUser)
	if err != nil {
		h.Logger.Error(err)
		return common.SendInternalServerErrorResponse(c, err.Error())
	}
	return common.SendSuccessResponse(c, "user login successful", map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user":          retrievedUser,
	})
}
