package client

import (
	"SiteAPI/pkg/secure"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetToken(w http.ResponseWriter, r *http.Request) {
	if secure.Token != "" {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", "http://localhost:8080/homePage", nil)
		req.Header.Set("Token", secure.Token)
		res, err := client.Do(req)
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("two")
			fmt.Fprintf(w, err.Error())

		}
		fmt.Fprintf(w, string(body))

	} else {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", "http://localhost:8080/autorisation", nil)
		res, err := client.Do(req)
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("two")
			fmt.Fprintf(w, err.Error())

		}
		fmt.Fprintf(w, string(body))
	}
}
