package main

import "chatroom/Routes"

func main() {
	r := Routes.SetupRouter()
	r.Run()
}
