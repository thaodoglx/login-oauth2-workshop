package middleware

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	con "gl.atisicloud.com/dev/sim-infinyscloud-authentication-provider/controllers"
)

// AuthProviderMiddleware function for this application middleware ...
func AuthProviderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get request url
		requestURL := c.Request.RequestURI

		// redirect to ory kratos endpoint if request url contains `/.ory/kratos/public`
		if strings.Contains(requestURL, con.LoginProviderSelfPublicAPIPath) {
			selfAPIPathLength := len(con.LoginProviderSelfPublicAPIPath)
			newRequestURL := fmt.Sprintf("%s/%s", con.KratosPublicAPI, requestURL[selfAPIPathLength:])

			proxyURL, err := url.Parse(newRequestURL)
			if err != nil {
				log.Println("[!] url parse error:", err.Error())
				return
			}

			proxyDirectory := func(req *http.Request) {
				req.URL.Scheme = proxyURL.Scheme
				req.URL.Host = proxyURL.Host
				req.URL.Path = proxyURL.Path
				req.Header.Add("X-Forwarded-Host", req.Host)
				req.Header.Add("X-Origin-Host", proxyURL.Host)
			}

			proxy := &httputil.ReverseProxy{
				Director: proxyDirectory,
			}

			proxy.ServeHTTP(c.Writer, c.Request)
		}

		// redirect user to login page if not logged in
		isLoggedIn := con.IsLoggedIn(c)

		if c.Request.URL.Path != "/login" && len(c.Query("request")) == 0 {
			if !isLoggedIn {
				c.Redirect(http.StatusTemporaryRedirect, "/login")
				return
			}
		}

		c.Next()
	}
}
