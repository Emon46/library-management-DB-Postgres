package main

import (
	"encoding/json"
	"errors"
	"libraryManagement/pkg"
	"net/http"
)

func CheckAuth(r *http.Request, userType string) error {
	userMail, password, ok := r.BasicAuth()
	if !ok {
		return errors.New("unAuthorized")
	}
	for _, userVar := range pkg.Users {
		if userVar.Mail == userMail && userVar.Password == password && userVar.UserType == userType {
			return nil
		} else if userVar.Mail == userMail {
			return errors.New("invalid password")
		}
	}
	return errors.New("user not found")
}
func AuthMiddleware(next http.Handler, requesterType string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := CheckAuth(r, requesterType)
		if err != nil {
			myResponse := pkg.MyData{
				Status:  http.StatusUnauthorized,
				Error:   err,
				Message: "unAuthorized",
				Success: "false",
				Data:    nil,
			}
			w.WriteHeader(http.StatusUnauthorized)

			json.NewEncoder(w).Encode(myResponse)
		} else {

			next.ServeHTTP(w, r)
		}
	})
}
