// Packate principal
package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	initValidations()
	usc := userController{}
	db, err := getDbConnection()

	if err != nil {
		panic("There was a problem connecting to the database")
	}

	usc.init(*db)

	r := gin.New()

	r.POST("/user", func(context *gin.Context) {
		userData := User{}
		var ve validator.ValidationErrors
		if err := context.ShouldBindJSON(&userData); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": getError(ve, err)})
			return
		}

		exist, errSql := usc.exists(userData)

		if errSql != nil {
			context.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"error": "There was a problem please try again later"},
			)
			return
		}

		if exist {
			context.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"error": "There is already a registered user with the email/user/phone data"},
			)
			return
		} else {
			usc.create(&userData)
		}

		context.JSON(http.StatusAccepted, userData)
	})

	r.POST("/login", func(context *gin.Context) {
		data := UserLogin{}
		usl := LoginData{}
		var ve validator.ValidationErrors
		if err := context.ShouldBindJSON(&data); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": getError(ve, err)})
			return
		}

		if errLogin := usc.login(data, &usl); errLogin != nil {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": errLogin.Error()})
			return
		}

		context.JSON(http.StatusAccepted, usl)
	})

	r.GET("/me", func(context *gin.Context) {
		h := AuthHeader{}
		usrDat := User{}
		if err := context.ShouldBindHeader(&h); err != nil {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "User token not provided"})
			return
		}

		valid, dt := valid(h.Authorization)

		if valid {
			id, _ := strconv.ParseUint(dt["id"], 10, 32)
			if errUsr := usc.getUser(id, &usrDat); errUsr != nil {
				context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "there is no user matching the ID"})
				return
			}
			context.JSON(http.StatusAccepted, usrDat)
		} else {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "User token not valid"})
			return
		}
	})

	r.Run(":3000")
}
