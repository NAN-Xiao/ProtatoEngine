package main

import "potatoengine/src/server"

func main() {
	server.LaunchServer()
	sp := NewLoginSpace("Login")
	server.RegistSpace(server.E_Loging, sp)
	server.Serv()
	select {}
}
