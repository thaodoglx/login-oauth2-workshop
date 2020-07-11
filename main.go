package main

import (
	"gl.atisicloud.com/dev/sim-infinyscloud-authentication-provider/server"
)

// Run application system ...
func main() {
	server := new(server.Server)
	server.Run()
}
