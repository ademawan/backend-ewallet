package main

import (
	"backend-ewallet/configs"
	"backend-ewallet/utils"
	"fmt"
	"log"

	uc "backend-ewallet/delivery/controllers/user"
	"backend-ewallet/delivery/routes"
	userRepo "backend-ewallet/repository/user"

	ac "backend-ewallet/delivery/controllers/auth"
	authRepo "backend-ewallet/repository/auth"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {

	config := configs.GetConfig()

	db := utils.InitDB(config)

	authRepo := authRepo.New(db)
	userRepo := userRepo.New(db)

	authController := ac.New(authRepo)
	userController := uc.New(userRepo)
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	routes.RegisterPath(e, userController, authController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
