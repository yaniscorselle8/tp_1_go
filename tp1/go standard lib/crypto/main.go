package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	hAsString1 := readImgAsByte("image_1.jpg")
	hAsString2 := readImgAsByte("image_2.jpg")
	hAsString3 := readImgAsByte("image_3.jpg")
	compareHash(hAsString1, hAsString2, hAsString3)
}

func readImgAsByte(fileName string) []byte {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	h := sha256.New()
	h.Write([]byte(content))

	hAsString := h.Sum(nil)
	fmt.Println(hAsString)

	return hAsString
}

func compareHash(hashImage1 []byte, hashImage2 []byte, hashImage3 []byte) {
	res1_2 := bytes.Compare(hashImage1, hashImage2)

	if res1_2 == 0 { // 1 et 2 sont pareilles
		res1_3 := bytes.Compare(hashImage1, hashImage3)
		if res1_3 == 0 { //1 et 3 sont pareilles
			fmt.Println("Something is wrong")
		} else { //1 et 3 différentes
			fmt.Println("Image 3 is different")
		}
	} else { //1 et 2 ne sont pas pareilles
		res1_3 := bytes.Compare(hashImage1, hashImage3)
		if res1_3 == 0 { //1 et 3 sont pareilles
			fmt.Println("Image 2 is different")
		} else { //1 et 3 différentes
			fmt.Println("Image 1 is different")
		}

	}

}
