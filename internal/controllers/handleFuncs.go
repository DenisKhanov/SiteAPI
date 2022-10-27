package controllers

import (
	"SiteAPI/internal/service"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func HandleFuncs() {
	//port := os.Getenv("PORT")
	rout := mux.NewRouter().StrictSlash(true)

	rout.HandleFunc("/", service.Index).Methods("GET")
	rout.HandleFunc("/post/{id:[0-9]+}", service.ViewPosts).Methods("GET")
	rout.HandleFunc("/check_user", service.Authentication).Methods("POST")
	rout.HandleFunc("/new_user", service.NewUser).Methods("POST")
	rout.HandleFunc("/create_post", service.SaveArticle).Methods("POST")
	rout.HandleFunc("/edit_post/post/{id:[0-9]+}", service.EditPost).Methods("GET", "POST")
	rout.HandleFunc("/delete_post/post/{id:[0-9]+}", service.Delete).Methods("GET")

	http.Handle("/", rout)

	//http.ListenAndServe(":"+port, nil)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
