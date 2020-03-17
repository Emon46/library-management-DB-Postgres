package db

import (
	"fmt"

	"github.com/emon331046/libraryManagement/pkg/model"
)

func CreateUser(user model.UserDb) *model.UserModel {
	var user1 model.UserDb
	okk, err := eng.Where("mail= ?", user.Mail).Get(&user1)
	if err != nil {
		fmt.Println("create user get method error: ", err)

	}
	//fmt.Println("pasisi ", okk)
	if okk {
		fmt.Println("****************i am already registered****************", user1.Mail)
		return nil

	} else {
		fmt.Println("*****************insert called*********")
		eng.Insert(user)
		eng.Where("mail=?", user.Mail).Get(&user1)
		return model.APIFormat(user1)

	}

}

func GetUser(userId int) *model.UserModel {
	var user model.UserDb
	okk, err := eng.Where("id= ?", userId).Get(&user)
	if err != nil {
		fmt.Println(err)
	}
	if okk {
		//fmt.Println("found a user", *model.APIFormat(user))
		return model.APIFormat(user)

	} else {
		return nil
	}
}
