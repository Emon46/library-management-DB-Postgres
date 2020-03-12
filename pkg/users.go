package pkg

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
	PhoneNo  string `json:"phone_no"`
	UserType string `json:"user_type"`
}

var UserIdCount = 1
var Users []User

func (user *User) CreateUser() bool {

	Users = append(Users, User{
		ID:       user.ID,
		Name:     user.Name,
		Mail:     user.Mail,
		Password: user.Password,
		PhoneNo:  user.PhoneNo,
		UserType: user.UserType,
	})
	return true
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

	for _, userVar := range Users {
		if userVar.ID == userId {
			myResponse := MyData{
				Status: http.StatusOK,

				Error:   nil,
				Success: "true",
				Message: "Found a user",
				Data:    userVar,
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(myResponse)
			return
		}
	}
	myResponse := MyData{
		Status:  http.StatusBadRequest,
		Error:   err1,
		Success: "false",
		Message: "No id available",
		Data:    nil,
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(myResponse)
	return
}

func EditUserProfile() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		userMail, _, ok := r.BasicAuth()
		if ok {

			for i, userVar := range Users {
				if userVar.Mail == userMail {
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
						Message: "Found a user",
						Data:    Users[i],
					}
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(myResponse)
					return
				}
			}
		}

	})
}
