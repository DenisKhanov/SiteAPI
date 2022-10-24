package service

import (
	"SiteAPI/internal/db"
	"SiteAPI/pkg/structs"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

//Главная страница
func Index(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("pkg/templates/index.html",
		"pkg/templates/header.html", "pkg/templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	//Выборка данных
	dataBase := db.DbConnect()
	var posts []structs.Article
	dataBase.Find(&posts)
	dataBase.Clauses()
	tmp.ExecuteTemplate(w, "index", posts)
	json.NewEncoder(w).Encode(posts)
}
