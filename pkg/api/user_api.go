package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/emon331046/libraryManagement/pkg/middleware"

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

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}

func UserProfile(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["user_id"]
	userId, err1 := strconv.Atoi(key)

	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(nil)
	}

	userResult := db.GetUser(userId)
	if userResult != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(userResult)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(userResult)
	return
}

func EditUserProfile(w http.ResponseWriter, r *http.Request) {
	//_, err := strconv.Atoi(r.Header.Get("current_user_id"))
	currentUserType := r.Header.Get("current_user_type")
	currentUserMail := r.Header.Get("current_user_mail")
	if currentUserType != "user" {

		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(errors.New("type of user didn't match"))
		return
	}
	if currentUserMail == "" {

		http.Error(w, "mail not valid", 404)
		return
	} else {

		var user model.UserModel
		err1 := json.NewDecoder(r.Body).Decode(&user)
		if err1 != nil {
			http.Error(w, err1.Error(), 404)
			return
		}
		user.Mail = currentUserMail
		resultUser, err := db.UpdateUserProfile(model.UserDbFormat(user))
		if err == nil {

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resultUser)
			return
		}

	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	var user model.UserModel
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user1 := db.LoginUser(model.UserDbFormat(user))
	//fmt.Println("hurrrrreeeeeeeeeeeeee", user1)

	if user1 != nil {
		tokenString, err := middleware.GenerateJWT(user1.Mail, user1.UserType, user1.ID)
		if err == nil {

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(tokenString)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(nil)
	return

}
