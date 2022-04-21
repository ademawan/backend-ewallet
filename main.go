package main

import (
	"backend-ewallet/configs"
	"backend-ewallet/utils"
	"fmt"
	"log"

	ac "backend-ewallet/delivery/controllers/auth"
	tc "backend-ewallet/delivery/controllers/transaction"
	uc "backend-ewallet/delivery/controllers/user"
	"backend-ewallet/delivery/routes"

	authRepo "backend-ewallet/repository/auth"
	transactionRepo "backend-ewallet/repository/transaction"
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

	authRepo := authRepo.New(db)
	userRepo := userRepo.New(db)
	transactionRepo := transactionRepo.New(db)

	authController := ac.New(authRepo)
	userController := uc.New(userRepo)
	transactionController := tc.New(transactionRepo)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	routes.RegisterPath(e, userController, authController, transactionController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
