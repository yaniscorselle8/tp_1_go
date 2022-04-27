package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Response struct {
	respText string
	err      error
}

func callServer(address string, channel chan Response) {
	requestResult, err := http.Get(address)
	//fmt.Println(address)
	//fmt.Println(requestResult)
	//fmt.Println(requestResult.StatusCode)
	var resultErreur Response
	var resultResponse Response
	if err != nil {
		resultErreur.err = err
		log.Fatal(err)
		go run(channel, resultErreur)
	}
	if requestResult.StatusCode != 200 {
		resultErreur.err = errors.New("Le code retourné par le serveur indique une erreur: " + strconv.Itoa(requestResult.StatusCode))
		go run(channel, resultErreur)

	} else {

		body, err := ioutil.ReadAll(requestResult.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer requestResult.Body.Close()
		resultResponse.respText = string(body)
		//fmt.Println(resultResponse)
		go run(channel, resultResponse)
	}

}

func run(c chan Response, value Response) {
	c <- value // envoyer une valeur d'un channel
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		maFonction()
	}()
	wg.Wait()
	channel1 := make(chan Response)
	channel2 := make(chan Response)

	go callServer("http://localhost:8080/?id=id1", channel1)
	go callServer("http://localhost:8080/?id=id42", channel2)
	select {
	case msg1 := <-channel1:
		fmt.Println("received", msg1)
	case msg2 := <-channel2:
		fmt.Println("received", msg2)
	}
	fmt.Println("Fin du programme")

}
func maFonction() {
	fmt.Println("j'ai fini !")
}
