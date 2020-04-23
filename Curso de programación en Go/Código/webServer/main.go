package main

func main() {
	myServer := newServer(":3000")
	myServer.listen()
}
