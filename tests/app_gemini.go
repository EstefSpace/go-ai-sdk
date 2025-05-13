package main

import (
	"fmt"
	"go-ai-sdk/gemini"
)

func main() {
	client := gemini.NewClient("AIzaSyAcsgKNsXXF4cixFMqtDvLj5s-AvmduAPU", "gemini-2.5-flash-preview-04-17")

	response, err := client.Ask("Explica como funciona la Inteligencia Artificial", "No uses markdown, es decir los ** los # los ##")

	if err != nil {
		fmt.Println("Ocurrio un error")
		fmt.Println(err)
	}

	fmt.Println(response)
}
