package secure

import (
	"SiteAPI/pkg/structs"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

var Token string
var mySigningKey = []byte("mysupersecurethrase")

func GenerateJWT(login string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &structs.UserToken{
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 12).Unix()},
		Username:       login,
	})

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil

}

func IsAuthorised(endpoint func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	fmt.Println("Start  " + Token)
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not Authorised!")
		}
	}
}
