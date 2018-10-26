package main

import (
	"fmt"

	"github.com/otiai10/gosseract"
)

func main() {
	ocrClient := gosseract.NewClient()
	defer ocrClient.Close()

	ocrClient.SetImage("images/Pingdom edit screen for Search.png")
	text, err := ocrClient.Text()

	if err != nil {
		fmt.Printf("%v", err)
	}
	else {
		fmt.Println(text)
	}
}
