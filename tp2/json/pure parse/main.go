package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type User struct {
	userName string
	Password string
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

}
