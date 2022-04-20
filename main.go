package main

import (
	"backend-ewallet/configs"
	"backend-ewallet/utils"
	"fmt"
	"log"

	uc "backend-ewallet/delivery/controllers/user"
	"backend-ewallet/delivery/routes"
	userRepo "backend-ewallet/repository/user"

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

	userRepo := userRepo.New(db)

	userController := uc.New(userRepo)
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	routes.RegisterPath(e, userController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
