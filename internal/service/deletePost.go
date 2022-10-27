package service

import (
	"SiteAPI/internal/db"
	"SiteAPI/pkg/structs"
	"github.com/gorilla/mux"
	"net/http"
)

//Удаление статьи
func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	dataBase := db.DbConnect()

	var post structs.Article
	dataBase.Delete(&post, id)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
