package server

import (
	"github.com/gin-gonic/gin"
	"gl.atisicloud.com/dev/sim-infinyscloud-authentication-provider/controllers"
	"gl.atisicloud.com/dev/sim-infinyscloud-authentication-provider/middleware"
	"gl.atisicloud.com/dev/sim-infinyscloud-authentication-provider/router"
)

// Server struct for define initialize golang web server (such as route, etc)
type Server struct{}

// Run methods for run golang web server
func (s *Server) Run() {
	// Initialize gin engine
	r := gin.Default()
	r.Use(middleware.AuthProviderMiddleware())

	// all routes endpoint here:
	router.Routes(r)

	r.Run(controllers.LoginProviderPort)
}
