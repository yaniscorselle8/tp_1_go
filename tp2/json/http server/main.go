package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type User struct {
	userID   int
	userName string
	Password string
}

func ServerHandler(response http.ResponseWriter, request *http.Request) {
	id := request.FormValue("id")
	for userID := range user {
		if id == user[userID] {
			userName := request.FormValue("userName")
			Password := request.FormValue("Password")
			data := map[string]string{
				"userID":   id,
				"userName": userName,
				"Password": Password,
			}
			user, err := json.MarshalIndent(data, "\n", "")
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(user))

		}
	}
}

var user map[string]string

func main() {
	data := map[string]string{
		"userID":   "2",
		"userName": "Paul",
		"Password": "pass123",
	}
	user, err := json.MarshalIndent(data, "\n", "")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(user))

	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	} else {
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var result map[string]interface{}
		json.Unmarshal([]byte(byteValue), &result)

		fmt.Println(result)
	}

	//ttp.HandleFunc("/", ServerHandler(w.Header().Set("Content-Type", "application/json; charset=utf-8"), "http://localhost:8000/?id=id1"))
	http.ListenAndServe("http://localhost:8000/?id=id1", nil)
}
