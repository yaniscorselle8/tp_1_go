package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Login    string `json:"userName"`
	Password string
	UserID   string `json:"userID"`
}

var userMap map[string]User

func main() {
	var user []User

	jsonFile, err := ioutil.ReadFile("users.json")
	if err != nil {
		fmt.Println(err)
	} else {

		json.Unmarshal(jsonFile, &user)

		fmt.Println(user)
	}

	userMap = make(map[string]User)
	for i := 0; i < len(user); i++ {
		userMap[user[i].UserID] = user[i]
	}

	fmt.Println(userMap)

	handler := func(w http.ResponseWriter, req *http.Request) {
		id := req.FormValue("id")
		for UserID := range userMap {
			if id == UserID {
				data := map[string]string{
					"Login":    req.FormValue("Login"),
					"Password": req.FormValue("Password"),
					"UserID":   req.FormValue("UserID"),
				}
				user, err := json.MarshalIndent(data, "\n", "")
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(string(user))
				w.WriteHeader(http.StatusOK)

				content_userMap, _ := json.Marshal(userMap[id])
				w.Write(content_userMap)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}

		}
		w.Header().Set("Content-Type",
			"application/json; charset=utf-8",
		)

	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
