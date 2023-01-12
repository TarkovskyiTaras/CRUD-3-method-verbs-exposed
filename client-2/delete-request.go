package main

import "net/http"

func main() {
	client := &http.Client{}
	url := "http://localhost:8080/person/2"

	request, _ := http.NewRequest("DELETE", url, nil)
	client.Do(request)
}
