package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{}
	url := "http://localhost:8080/person/3"
	request, _ := http.NewRequest("GET", url, nil)

	response, _ := client.Do(request)
	responseBody, _ := io.ReadAll(response.Body)

	fmt.Println(string(responseBody))
}
