package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var MySigningKey = []byte("secret1234")

func JwtMiddleWare(next http.Handler) http.Handler {

	fmt.Println("here middleware ")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.URL.Path)

		//if get then here i have given access the user with or without token
		if r.Method == "GET" || r.URL.Path == "/register" {
			log.Println(r.Method, "request")
			next.ServeHTTP(w, r)
		} else {
			//in case of other write operation i need to check if the user is valid
			fmt.Println("checking with jwt middleware")

			//retrieve the auth header
			authHeader := r.Header.Get("Authorization")
			//if there is no auth header
			if authHeader == "" {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("need auth token"))
				return
			} else {

				// parsing the token from (bearer tokenString) with the secret key
				token, _ := jwt.Parse(strings.Split(authHeader, " ")[1], func(token *jwt.Token) (interface{}, error) {
					return MySigningKey, nil
				})
				//checking if there is any claim and s the token valid
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					fmt.Println("user validity : ok jwt")
					//fmt.Println(claims)
					//retrieving the claims info

					userId, okId := claims["userId"].(float64)

					//as i have set it capital i can use it outside the package
					CurrentUserId := int(userId)
					CurrentUserMail, okMail := claims["userMail"].(string)
					CurrentUserType, okType := claims["userType"].(string)

					if okId && okMail && okType {
						r.Header.Set("current_user_id", string(CurrentUserId))
						r.Header.Set("current_user_type", string(CurrentUserType))
						r.Header.Set("current_user_mail", string(CurrentUserMail))
						//fmt.Println(CurrentUserId, CurrentUserMail, CurrentUserType, okMail, okType)
						next.ServeHTTP(w, r)
					} else {
						w.WriteHeader(http.StatusNotAcceptable)
						w.Write([]byte("the token is not valid .missing some info"))
					}

				} else {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("need auth token"))
					return
				}
			}
		}
	})
}

func GenerateJWT(userMail string, userType string, userId int) (string, error) {
	//signing method declare
	token := jwt.New(jwt.SigningMethodHS256)

	//passing the parameter that i want to keep i my token
	claims := token.Claims.(jwt.MapClaims)
	claims["userMail"] = userMail
	claims["userType"] = userType
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Minute * 300).Unix()

	//generating the token string with my signing keys and claims parameter
	tokeString, err := token.SignedString(MySigningKey)
	if err != nil {
		return "", err
	}
	//fmt.Println("otkkkkkkkkkkkkkkkeeeeeeeeennnnn", tokeString)
	return tokeString, err

}
