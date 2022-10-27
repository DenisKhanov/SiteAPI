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
Authentication
Данные передаются в формате:
{
		"Login": "users11@mail.ru",
		"Password": "12344321"
	}
*/
func Authentication(w http.ResponseWriter, r *http.Request) {

	req := &structs.Users{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		fmt.Println(w, r, http.StatusBadRequest, err)
	}
	req.Password = secure.HashPassword(req.Password)
	if req.Login == "" || req.Password == "" {
		fmt.Println("Поля логин или пароль не могут быть пустыми")
	} else {
		//Вытягивание строки из БД по логину и проверка с введенным паролем
		dataBase := db.DbConnect()
		var user structs.Users
		dataBase.Find(&user, "login", req.Login)
		if user.Login == "" {
			fmt.Printf("Пользователь с таким email %s не зарегистрирован", req.Login)
		} else {
			if user.Password == req.Password {
				fmt.Printf("%s, мы вас узнали!", req.Login)
				token, err := secure.GenerateJWT(req.Login)
				if err != nil {
					fmt.Println(err)
				}
				secure.Token = token
				err = json.NewEncoder(w).Encode(secure.Token)
				if err != nil {
					fmt.Println(err)
				}

			} else {
				fmt.Println("Сочетание логина и пароля не верны!")
			}
		}
	}
}
