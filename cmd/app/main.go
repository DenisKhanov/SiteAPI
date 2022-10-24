package main

import (
	"SiteAPI/internal/controllers"
	"SiteAPI/internal/db"
)

func main() {
	db.InitialMigration()
	controllers.HandleFuncs()
}

//id := r.URL.Query().Get("id") Парсинг URL при передачи данных в конце URL после знака ?
