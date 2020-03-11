package pkg

import (
	"encoding/json"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	for _, userVar := range Users {
		if userVar.Mail == user.Mail {
			myResponse := MyData{
				Status:  502,
				Error:   "A user is already registered with this mail",
				Success: "false",
				Message: "register failed",
			}
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(myResponse)
			return
		}
	}

	user.ID = UserIdCount
	UserIdCount++
	user.CreateUser()
	myResponse := MyData{
		Status:  http.StatusCreated,
		Error:   "null",
		Message: "created new user",
		Success: "true",
		Data:    Users,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(myResponse)

}
