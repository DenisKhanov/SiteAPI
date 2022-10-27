package service

import (
	"SiteAPI/internal/db"
	"SiteAPI/pkg/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

// Index Главная страница
func Index(w http.ResponseWriter, _ *http.Request) {

	//Выборка данных
	dataBase := db.DbConnect()
	var posts []structs.Article
	dataBase.Find(&posts)
	err := json.NewEncoder(w).Encode(posts)
	if err != nil {
		fmt.Println(err)
	}
}
