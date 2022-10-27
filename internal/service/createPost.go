package service

import (
	"SiteAPI/internal/db"
	"SiteAPI/pkg/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

/*
SaveArticle
Данные передаются в виде:
{
	"Title": "Jsonishe",
	"Anons": "API",
	"FullText": "text"
}
*/
func SaveArticle(w http.ResponseWriter, r *http.Request) {

	req := &structs.Article{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		fmt.Println(w, r, http.StatusBadRequest, err)
	}
	dataBase := db.DbConnect()
	dataBase.Create(req)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
