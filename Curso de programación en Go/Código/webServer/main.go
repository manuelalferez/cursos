package main

func main() {
	myServer := newServer(":3000")
	myServer.setHandler("/", HandlerRoot)
	myServer.setHandler("/api", myServer.addMiddleware(HandlerAPI, checkAuth(), Loggin()))
	myServer.listen()
}
