package api

import (
	"fmt"
	"net/http"
	"strconv"

	"gopkg.in/macaron.v1"

	"github.com/Emon331046/libraryManagement/pkg/middleware"

	"github.com/Emon331046/libraryManagement/pkg/model"

	"github.com/Emon331046/libraryManagement/pkg/db"
)

var Users []model.UserModel

func Register(ctx *macaron.Context, user model.UserModel) {
	result := db.CreateUser(model.UserDbFormat(user))
	if result == nil {
		fmt.Println("***************crete user false***************")

		ctx.JSON(http.StatusNotImplemented, "the user already exist")
		return
	}

	ctx.JSON(http.StatusCreated, result)

}

func UserProfile(ctx *macaron.Context) {
	key := ctx.Params(":userId")
	userId, err1 := strconv.Atoi(key)

	if err1 != nil {

		ctx.JSON(http.StatusBadRequest, "invalid user profile")
	}

	userResult := db.GetUser(userId)
	if userResult != nil {
		ctx.JSON(http.StatusOK, userResult)
		return
	}

	ctx.JSON(http.StatusBadRequest, nil)
	return
}

func EditUserProfile(ctx *macaron.Context, user model.UserModel) {
	currentUserType := ctx.Req.Header.Get("current_user_type")
	currentUserMail := ctx.Req.Header.Get("current_user_mail")
	if currentUserType != "user" {
		ctx.JSON(http.StatusNotAcceptable, "type of user didn't match")
		return
	}
	if currentUserMail == "" {
		ctx.JSON(http.StatusNotAcceptable, "mail not valid")
		return
	} else {
		user.Mail = currentUserMail
		resultUser, err := db.UpdateUserProfile(model.UserDbFormat(user))
		if err == nil {

			ctx.JSON(http.StatusOK, resultUser)
			return
		} else {
			ctx.JSON(http.StatusNotImplemented, "update failed")
			return
		}
	}
}

func Login(ctx *macaron.Context, user model.UserModel) {

	user1 := db.LoginUser(model.UserDbFormat(user))
	//fmt.Println("hurrrrreeeeeeeeeeeeee", user1)

	if user1 != nil {
		tokenString, err := middleware.GenerateJWT(user1.Mail, user1.UserType, user1.ID)
		if err == nil {

			ctx.JSON(http.StatusOK, tokenString)
			return
		}

		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	ctx.JSON(http.StatusUnauthorized, "user password or mail didn't match")
	return

}
