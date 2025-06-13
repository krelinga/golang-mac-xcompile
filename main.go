package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")
	resp, err := http.Get("https://tcserver.i.krel.ing/")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}
	fmt.Println(string(body))
}
