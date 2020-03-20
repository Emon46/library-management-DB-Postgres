package main

import (
	"log"
	"net/http"

	"gopkg.in/macaron.v1"

	"github.com/go-macaron/binding"

	"github.com/Emon331046/libraryManagement/pkg/model"

	"github.com/Emon331046/libraryManagement/pkg/middleware"

	"github.com/Emon331046/libraryManagement/pkg/api"
)

func main() {
	//*********************go-macaron**********************
	m := macaron.Classic()
	m.Use(macaron.Renderer())

	m.Use(middleware.JwtMiddleWare)

	m.Get("/login", binding.Json(model.UserModel{}), api.Login)
	m.Post("/register", binding.Json(model.UserModel{}), api.Register)
	m.Get("/user-profile/:userId([0-9]+)", api.UserProfile)
	m.Patch("/edit-profile", binding.Json(model.UserModel{}), api.EditUserProfile)

	m.Post("/purchase-book", binding.Json(model.BookHistoryDb{}), api.AddNewPurchase)
	m.Put("/return-book", binding.Json(model.BookHistoryDb{}), api.ReturnBook)
	m.Post("/book", binding.Json(model.Bookdb{}), api.AddNewBook)
	m.Get("/book", api.ShowAllBooks)
	m.Get("/book/:bookId([0-9]+)", api.ShowBook)
	m.Delete("/delete-book/:bookId([0-9]+)", api.DeleteBook)
	log.Fatal(http.ListenAndServe("0.0.0.0:4000", m))

	//*********************gorilla-mux***********************

	//r := mux.NewRouter()
	//
	//http.Handle("/", r)
	//srv := &http.Server{
	//	Handler: r,
	//	Addr:    "127.0.0.1:8000",
	//	// Good practice: enforce timeouts for servers you create!
	//	WriteTimeout: 15 * time.Second,
	//	ReadTimeout:  15 * time.Second,
	//}
	//
	//log.Fatal(srv.ListenAndServe())

}
