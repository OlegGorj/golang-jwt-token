package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

// Secret key for signing JWT tokens
var mySigningKey = []byte("secret")

// Handler for the home page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

// Handler for the login endpoint
func login(w http.ResponseWriter, r *http.Request) {
	// Get the username and password from the request body
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Check if the username and password are valid
	if username == "testuser" && password == "password123" {
		// Set the expiration time of the token to 5 minutes from now
		expirationTime := time.Now().Add(5 * time.Minute)

		// Create the JWT claims, which includes the username and expiration time
		claims := &jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Subject:   username,
		}

		// Create the JWT token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Sign the token with the secret key
		signedToken, err := token.SignedString(mySigningKey)
		if err != nil {
			fmt.Println("Error signing token: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Send the token as a response to the client
		w.Write([]byte(signedToken))
	} else {
		// Return a 401 Unauthorized status if the username or password is invalid
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid username or password"))
	}
}

// Handler for the protected endpoint
func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	// Get the JWT token from the Authorization header
	authHeader := r.Header.Get("Authorization")
	tokenString := authHeader[len("Bearer "):]

	// Parse the JWT token
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Println("Error parsing token: ", err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid token"))
		return
	}

	// Check if the token is valid
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid token"))
		return
	}

	// Get the username from the token's claims
	username := claims.Subject

	// Send a response with the protected data
	fmt.Fprintf(w, "Welcome, %s! This is protected data.", username)
}

func main() {
	// Define the server routes
	http.HandleFunc("/", homePage)
	http.HandleFunc("/login", login)
	http.HandleFunc("/protected", protectedEndpoint)

	// Start the server
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
