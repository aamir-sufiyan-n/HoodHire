package utils

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(pass string) (string,error) {
	hash,err:=bcrypt.GenerateFromPassword([]byte(pass),bcrypt.DefaultCost)
	if err!=nil{
		return "",err
	}
	return string(hash),nil
}

func ComparePass(hash,pass string)bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hash),[]byte(pass))
	return err==nil
}
