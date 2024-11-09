package handler

import (
	"errors"
	"fmt"
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"
	"go-fiber-gorm/model/response"
	"go-fiber-gorm/utils"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var err error

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

func UserHandlerGetAll(c *fiber.Ctx) error {
	var users []entity.User

	err := database.DB.Find(&users).Error
	if err != nil {
		log.Println(err)
	}

	return utils.Response(c, 200, "User created successfully", users)
}

func UserHandlerCreate(c *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err = c.BodyParser(user); err != nil {
		log.Println(err)
	}

	validationErrors := []ErrorResponse{}

	validate := validator.New()
	errValidation := validate.Struct(user)
	if errValidation != nil {
		for _, err := range errValidation.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}

		return utils.Response(c, 400, "Error", validationErrors)
	}

	newUser := entity.User{
		Name:    user.Name,
		Address: user.Address,
		Phone:   user.Phone,
		Email:   user.Email,
	}

	err = database.DB.Table("users").Create(&newUser).Error
	if err != nil {
		log.Println(err)
	}

	return utils.Response(c, 200, "User created successfully", newUser)
}

func UserHandlerGetById(c *fiber.Ctx) error {
	userId := c.Params("id")

	var user response.User

	err = database.DB.Table("users").Where("id", userId).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.Response(c, 404, "User not found", nil)
		}

		return utils.Response(c, 500, err.Error(), nil)
	}

	return utils.Response(c, 200, "User found", user)
}
