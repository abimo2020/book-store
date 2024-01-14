package main

import (
	"github.com/abimo2020/book-store/routes"
	"github.com/abimo2020/book-store/utils"
	"github.com/go-playground/validator"
)

func main() {
	e := routes.New()
	e.Validator = &utils.CustomValidator{
		Validator: validator.New(),
	}
	e.Logger.Fatal(e.Start(":8080"))
}
