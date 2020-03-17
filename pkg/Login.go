package pkg

//func Login(w http.ResponseWriter, r *http.Request) {
//	var user User
//	err := json.NewDecoder(r.Body).Decode(&user)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	for _, userVar := range Users {
//		if userVar.Mail == user.Mail && userVar.Password == user.Password {
//			tokenString, err := GenerateJWT(userVar.Mail, userVar.UserType, userVar.ID)
//
//			myResponse := MyData{
//				Status: http.StatusOK,
//
//				Error:   err,
//				Success: "true",
//				Message: "Logged in successfully",
//				Data:    tokenString,
//			}
//			w.WriteHeader(http.StatusOK)
//			json.NewEncoder(w).Encode(myResponse)
//			return
//		}
//	}
//	myResponse := MyData{
//		Status:  http.StatusUnauthorized,
//		Error:   err,
//		Success: "false",
//		Message: "Log in failed",
//		Data:    nil,
//	}
//	w.WriteHeader(http.StatusUnauthorized)
//	json.NewEncoder(w).Encode(myResponse)
//	return
//}
