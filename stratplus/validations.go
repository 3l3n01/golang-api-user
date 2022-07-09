package main

import (
	"regexp"
	"unicode"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Personalized validation to check the security level of the password
func verifyPassword(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	for _, c := range s {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}

	if len(s) >= 6 && len(s) <= 12 {
		hasMinLen = true
	}

	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

// Init valitations
func initValidations() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
			return verifyPassword(fl.Field().String())
		})
		v.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
			r, _ := regexp.Compile("^[0-9]{10}$")
			return r.MatchString(fl.Field().String())
		})
	}
}
