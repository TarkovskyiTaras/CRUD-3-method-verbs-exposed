package main

import (
	"bytes"
	"net/http"
)

func main() {
	//jsonMessage := `{
	//					"id" : 1,
	//					"first_name" : "Taras",
	//					"last_name" : "Tarkovskyi",
	//					"dob" : "January 23d",
	//					"address_and_phone" : "Kyiv, 0933115485"
	//				}`

	//jsonMessage := `{
	//					"id" : 2,
	//					"first_name" : "Nikita",
	//					"last_name" : "Tarkovskyi",
	//					"dob" : "January 23d",
	//					"address_and_phone" : "Kyiv, 33333333"
	//				}`

	jsonMessage := `{
						"id" : 3,
						"first_name" : "Sergey",
						"last_name" : "Onishenko",
						"dob" : "December 28th",
						"address_and_phone" : "Kyiv, 5555555555"
					}`

	client := &http.Client{}
	myUrl := "http://localhost:8080/person/"
	body1 := bytes.NewBuffer([]byte(jsonMessage))

	request, _ := http.NewRequest("taras", myUrl, body1)
	client.Do(request)
}
