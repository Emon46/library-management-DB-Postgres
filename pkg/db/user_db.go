package db

import (
	"fmt"

	"github.com/Emon331046/libraryManagement/pkg/model"
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
		return nil
	}
	if okk {
		//fmt.Println("found a user", *model.APIFormat(user))
		return model.APIFormat(user)

	} else {
		return nil
	}
}
func LoginUser(user model.UserDb) *model.UserModel {
	var user1 model.UserDb
	//fmt.Println(user)
	okk, err := eng.Where("mail = ? AND password = ?", user.Mail, user.Password).Get(&user1)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if okk {
		//fmt.Println("found a user", *model.APIFormat(user1))
		return model.APIFormat(user1)

	} else {
		return nil
	}
}
func UpdateUserProfile(user model.UserDb) (*model.UserModel, error) {
	var user1 model.UserDb
	okk, err := eng.Where("mail = ? ", user.Mail).Get(&user1)
	if err != nil {
		return nil, err
	}
	if okk {
		//fmt.Println(user1)

		//edit
		if user.Name != "" {
			//fmt.Println("name ", user.Name, userVar.ID, i, userVar.Name, Users[i].Name)
			user1.Name = user.Name
		}
		if user.PhoneNo != "" {
			//fmt.Println("phn ", user.PhoneNo)
			user1.PhoneNo = user.PhoneNo
		}

		if user.Password != "" {
			//fmt.Println("pass ", user.Password)
			user1.Password = user.Password
		}
		_, err := eng.ID(user1.ID).Update(user1)
		//fmt.Println(user1, err)
		if err != nil {
			return nil, err
		}

	}
	return model.APIFormat(user1), nil

}
