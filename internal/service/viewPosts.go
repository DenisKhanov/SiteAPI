package service

import (
	"SiteAPI/internal/db"
	"SiteAPI/pkg/structs"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// ViewPosts Развернутый просмотр статьи
func ViewPosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.WriteHeader(http.StatusOK)

	dataBase := db.DbConnect()
	var post structs.Article
	dataBase.Find(&post, id)
	err := json.NewEncoder(w).Encode(post)
	if err != nil {
		fmt.Println(err)
	}
}
