package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type User struct {
	Login    string `json:"userName"`
	Password string
	UserID   string `json:"userID"`
}

func main() {
	data := map[string]string{
		"userName": "Paul",
		"Password": "pass123",
	}
	user, err := json.MarshalIndent(data, "\n", "")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(user))

	jsonFile, err := ioutil.ReadFile("users.json")
	jsonvar(err, jsonFile)

}

func jsonvar(err error, jsonFile []byte) (user []User) {
	if err != nil {
		fmt.Println(err)
	} else {

		json.Unmarshal(jsonFile, &user)

		fmt.Println(user)
	}
	return user
}
