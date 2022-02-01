package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/skip2/go-qrcode"
)

func main() {
	a, err := os.Open("qrcode.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer a.Close()

	b, err := ioutil.ReadAll(a)
	if err != nil {
		log.Fatal(err)
	}

	abc := string(b)

	newQR := "qrcode.png"

	err = qrcode.WriteFile(abc, qrcode.Medium, 51, newQR)

	if err != nil {
		log.Fatal(err)
	}
}
