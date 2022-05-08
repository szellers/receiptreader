package main

import (
	"fmt"
	"github.com/otai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()

	client.SetImage("receipt_test")

	text, _ := client.Text()
	fmt.Println(text)
}
