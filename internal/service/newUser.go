package service

import (
	"SiteAPI/internal/db"
	"SiteAPI/pkg/secure"
	"SiteAPI/pkg/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

/*
NewUser
Данные передаются в формате:
{
	"Login": "users11@mail.ru",
	"Password": "12344321"
}
*/
func NewUser(w http.ResponseWriter, r *http.Request) {

	req := &structs.Users{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		fmt.Println(w, r, http.StatusBadRequest, err)
	}
	req.Password = secure.HashPassword(req.Password)

	dataBase := db.DbConnect()
	var user structs.Users
	dataBase.Find(&user, "login", req.Login)

	if len(req.Login) < 4 || len(req.Password) < 4 {
		fmt.Println("Логин и пароль не могут быть менее 4 символов")
	} else if user.Login == "" {
		//Добавление данных
		dataBase.Create(req)
		fmt.Println("Пользователь успешно зарегистрирован")
	} else {
		fmt.Println("Пользователь уже зарегистрирован")
	}
}
