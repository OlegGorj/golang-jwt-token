package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Login to the server to get a JWT token
	response, err := http.PostForm("http://localhost:8080/login",
		map[string][]string{
			"username": {"testuser"},
			"password": {"password123"},
		})
	if err != nil {
		fmt.Println("Error logging in: ", err)
		return
	}
	defer response.Body.Close()

	// Read the token from the response body
	tokenBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading token: ", err)
		return
	}
	tokenString := string(tokenBytes)

	// Make a request to the protected endpoint with the JWT token in the Authorization header
	req, err := http.NewRequest("GET", "http://localhost:8080/protected", nil)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+tokenString)
	response, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making request: ", err)
		return
	}
	defer response.Body.Close()

	// Read the response body and print it
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response: ", err)
		return
	}
	fmt.Println(string(bodyBytes))
}
