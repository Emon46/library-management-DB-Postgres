package pkg

import (
	"encoding/json"
	"net/http"
)
type TheData interface{}

type MyData struct {
	Success  string   `json:"success"`
	Status  int   `json:"status"`
	Error  string   `json:"error"`
	Message  string   `json:"message"`
	Data TheData `json:"data"`
}

func Register(w http.ResponseWriter, r *http.Request)  {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w,err.Error(),404)
		return
	}

	for _, userVar := range Users {
		if userVar.Mail == user.Mail {
			myResponse := MyData{
				Status: 502,
				Error: "A user is already registered with this mail",
				Success: "false",

			}
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(myResponse)
			return
		}
	}

	user.ID = len(Users)+1
	user.CreateUser()
	myResponse := MyData{
		Status: 201,
		Error: "null",
		Message : "created new user",
		Success: "true",
		Data: Users,

	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(myResponse)

}
