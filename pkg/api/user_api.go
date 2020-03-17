package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/emon331046/libraryManagement/pkg/model"

	"github.com/emon331046/libraryManagement/pkg/db"
)

var Users []model.UserModel

func Register(w http.ResponseWriter, r *http.Request) {
	var user model.UserModel
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	result := db.CreateUser(model.UserDbFormat(user))
	if result == nil {
		fmt.Println("***************crete user false***************")
		myResponse := MyData{
			Status:  http.StatusCreated,
			Error:   nil,
			Message: "cannot create  user",
			Success: "true",
			Data:    result,
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(myResponse)
		return
	}

	myResponse := MyData{
		Status:  http.StatusCreated,
		Error:   nil,
		Message: "created new user",
		Success: "true",
		Data:    result,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(myResponse)

}

func UserProfile(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["user_id"]
	userId, err1 := strconv.Atoi(key)

	if err1 != nil {
		myResponse := MyData{
			Status:  http.StatusBadRequest,
			Error:   err1,
			Message: "no match for this user",
			Success: "true",
			Data:    nil,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(myResponse)
	}

	userResult := db.GetUser(userId)
	if userResult != nil {
		myResponse := MyData{
			Status:  http.StatusBadRequest,
			Error:   err1,
			Success: "false",
			Message: "user available",
			Data:    userResult,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(myResponse)
		return
	}

	myResponse := MyData{
		Status:  http.StatusBadRequest,
		Error:   err1,
		Success: "false",
		Message: "No id available",
		Data:    userResult,
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(myResponse)
	return
}

func EditUserProfile(w http.ResponseWriter, r *http.Request) {
	//_, err := strconv.Atoi(r.Header.Get("current_user_id"))
	currentUserType := r.Header.Get("current_user_type")
	currentUserMail := r.Header.Get("current_user_mail")
	if currentUserType != "user" {
		myResponse := MyData{
			Status:  http.StatusNotAcceptable,
			Error:   errors.New("user type didn't match"),
			Message: "only user can edit his profile info",
			Success: "true",
			Data:    nil,
		}

		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(myResponse)
		return
	}
	if currentUserMail == "" {

		http.Error(w, "mail not valid", 404)
		return
	} else {

		var user UserModel
		err1 := json.NewDecoder(r.Body).Decode(&user)

		if err1 != nil {
			http.Error(w, err1.Error(), 404)
			return
		}

		for i, userVar := range Users {
			if userVar.Mail == currentUserMail {
				//fmt.Println("dfssssssssssssss ", user.Name)
				//edit
				if user.Name != "" {
					//fmt.Println("name ", user.Name, userVar.ID, i, userVar.Name, Users[i].Name)
					Users[i].Name = user.Name
				}
				if user.PhoneNo != "" {
					//fmt.Println("phn ", user.PhoneNo)
					Users[userVar.ID].PhoneNo = user.PhoneNo
				}

				if user.Password != "" {
					//fmt.Println("pass ", user.Password)
					Users[userVar.ID].Password = user.Password
				}

				myResponse := MyData{
					Status: http.StatusOK,

					Error:   nil,
					Success: "true",
					Message: "updated user",
					Data:    Users[i],
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(myResponse)
				return
			}
		}
	}

}
