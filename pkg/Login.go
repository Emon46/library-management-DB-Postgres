package pkg

import (
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request)  {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w,err.Error(),404)
		return
	}

	for _, userVar := range Users {
		if userVar.Mail == user.Mail && userVar.Password == user.Password {
			myResponse := MyData{
				Status: http.StatusOK,

				Error: "null",
				Success: "true",
				Message: "Logged in successfully",
				Data: userVar,
			}
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(myResponse)
			return
		}
	}
	myResponse := MyData{
		Status: http.StatusBadRequest,
		Error: "password or username didn't match",
		Success: "false",
		Message: "Log in failed",
		Data: nil,
	}
	w.WriteHeader(http.StatusBadGateway)
	json.NewEncoder(w).Encode(myResponse)
	return
}