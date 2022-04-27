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

	userMapInit(user)

	handler := HandlerInit()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func userMapInit(user []User) {
	userMap = make(map[string]User)
	for i := 0; i < len(user); i++ {
		userMap[user[i].UserID] = user[i]
	}

	fmt.Println(userMap)
}

func HandlerInit() func(w http.ResponseWriter, req *http.Request) {
	handler := func(w http.ResponseWriter, req *http.Request) {
		id := req.FormValue("id")
		fmt.Println(id)
		idExist := false
		for UserID := range userMap {
			if id == UserID {
				idExist = true
				data := map[string]string{
					"Login":    userMap[id].Login,
					"Password": userMap[id].Password,
					"UserID":   userMap[id].UserID,
				}
				user, err := json.MarshalIndent(data, "\n", "")
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(string(user))
				w.WriteHeader(http.StatusOK)

				content_userMap, _ := json.Marshal(userMap[id])
				w.Write(content_userMap)
			}

		}
		if idExist == false {
			w.WriteHeader(http.StatusNotFound)
		}
		w.Header().Set("Content-Type",
			"application/json; charset=utf-8",
		)

	}
	return handler
}
