package service

import (
	"SiteAPI/internal/db"
	"SiteAPI/pkg/structs"
	"github.com/gorilla/mux"
	"net/http"
)

//Редактирование статьи
func EditPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	fullText := r.FormValue("full_text")

	dataBase := db.DbConnect()
	dataBase.Model(&structs.Article{}).Where("id", id).Updates(&structs.Article{Title: title,
		Anons: anons, FullText: fullText})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
