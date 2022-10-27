package main

import (
	"SiteAPI/internal/controllers"
	"SiteAPI/internal/db"
)

// @title           Swagger Example API
// @version         1.8.7
// @description     This is a sample server celler server.

// @host      localhost:8080
// @BasePath  /

func main() {
	db.InitialMigration()
	controllers.HandleFuncs()
}

//id := r.URL.Query().Get("id") Парсинг URL при передачи данных в конце URL после знака ?
