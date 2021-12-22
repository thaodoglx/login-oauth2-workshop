package router

import (
	"github.com/gin-gonic/gin"
	"gl.atisicloud.com/dev/sim-infinyscloud-authentication-provider/controllers"
)

// Routes methods for define route endpoint
func Routes(route *gin.Engine) {
	r := route.Group("/")

	// Login Provider routes
	r.GET("/", controllers.Home)
	r.GET("/registration", controllers.Registration)
	r.GET("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)
	r.GET("/sessions", controllers.Sessions)

	// OAuth 2.0 Credentials routes
	r.GET("/credentials", controllers.Credentials)
	r.GET("/credentials/create/oauth2_client", controllers.CreateOAuth2Client)

	// OAuth 2.0 ORY Hydra workflow routes
	r.GET("/oauth2/login", controllers.OAuth2Login)
	r.GET("/oauth2/consent", controllers.OAuth2Consent)
	r.GET("/oauth2/callback", controllers.OAuth2Callback)
	r.POST("/oauth2/access_token", controllers.GenerateAccessToken)
}
