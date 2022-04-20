package main

import (
	"backend-ewallet/configs"
	"backend-ewallet/utils"
	"fmt"
	"log"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
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

	fmt.Println(db)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
