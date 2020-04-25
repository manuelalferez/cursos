package main

func main() {
	myServer := newServer(":3000")
	myServer.setHandler("GET", "/", HandlerRoot)
	myServer.setHandler("POST", "/user", UserPostRequest)
	myServer.setHandler("POST", "/create", PostRequest)
	myServer.setHandler("POST", "/api", myServer.addMiddleware(HandlerAPI, checkAuth(), Loggin()))
	myServer.listen()
}
