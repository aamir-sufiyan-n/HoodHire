package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var validate = validator.New()

func BindAndValidate[T any](c fiber.Ctx) (*T,error) {
	var body T
	if err:= c.Bind().Body(&body);err !=nil{
		return nil,errors.New("invalid credentials")
	}
	if err:=validate.Struct(body);err!=nil{
		errorr:=make(map[string]string)
		for _,e:=range err.(validator.ValidationErrors){
			errorr[e.Field()]=e.Tag()
		}
		return nil,errors.New("validation error")
	}
	return  &body,nil
}