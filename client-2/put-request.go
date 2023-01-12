package main

import (
	"bytes"
	"net/http"
)

func main() {
	client := &http.Client{}
	url := "http://localhost:8080/person/3"

	jsonMessage := `{
						"id" : 3,
						"first_name" : "11111111111",
						"last_name" : "1111111111111",
						"dob" : "1111111111111",
						"address_and_phone" : "11111111111111"
					}`

	body := bytes.NewBuffer([]byte(jsonMessage))

	request, _ := http.NewRequest("PUT", url, body)
	client.Do(request)
}
