# golang-jwt-token

## A simple app to implement JWT token auth flow

JWT stands for "JSON Web Token". It's a compact, URL-safe means of representing claims to be transferred between two parties. In other words, it's a way to securely transmit information between a client (like a web browser) and a server.

A JWT is composed of three parts: a header, a payload, and a signature. The header contains information about the type of token and the algorithm used to sign it. The payload contains the claims or information that is being transmitted, such as a user ID or access privileges. Finally, the signature is a hash of the header, payload, and a secret key that only the server knows. The signature ensures that the token has not been tampered with and that the sender is who they claim to be.

JWTs are often used for authentication and authorization. When a user logs in to a website, for example, the server can create a JWT that contains information about the user, such as their username and user ID. The server then sends this JWT to the client, which can store it in a cookie or local storage. The client can then send the JWT back to the server with each subsequent request, allowing the server to verify the user's identity and grant access to resources.


JWTs are a secure and efficient way to transmit information between a client and server, especially for authentication and authorization purposes. They are widely used in modern web applications and APIs.

### Server side 

Quick server side example, we define a few routes for our server:

- The `/` function is a simple endpoint that displays a welcome message.
- The `/login` endpoint authenticates a user and generates a JWT token that is sent back to the client.
- The `/protected` endpoint is protected by JWT authentication. The client must send a JWT token in the Authorization header to access this endpoint.

The login function generates a JWT token with a 5-minute expiration time, and includes the username "testuser" in the token's claims. The isAuthorized middleware extracts the JWT token from the client's request, parses it, and checks if it is valid. If the token is valid, the middleware calls the requested endpoint.

### Client side

This client uses the `http.PostForm` method to send a username and password to the server's `/login` endpoint and receive a JWT token in the response. 
The token is then included in the Authorization header of a `GET` request to the server's protected `/protected` endpoint using `http.NewRequest`. 
Finally, the client reads the response body and prints it to the console.

Note that this is just a very quick and simple example for demo purposes of how to use JWT tokens in a client-server application.
In a real-world scenario, we would likely use more secure methods for storing and transmitting tokens, such as HTTPS or secure cookies.

Well, util next time..