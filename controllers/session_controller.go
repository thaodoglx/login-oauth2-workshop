package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"gl.atisicloud.com/dev/sim-infinyscloud-authentication-provider/models"
	"gl.atisicloud.com/dev/sim-infinyscloud-authentication-provider/utils"
)

var oauth2Store = sessions.NewCookieStore([]byte(LoginProviderSessionsSecret))

// GetCurrentSession for retrieve current user's sesssion from ORY Kratos API
func GetCurrentSession(c *gin.Context) (*models.KratosSession, error) {
	url := fmt.Sprintf("%s%ssessions/whoami", LoginProviderCompleteHost, LoginProviderSelfPublicAPIPath)

	resp, err := utils.Fetch(c, http.MethodGet, url, nil, utils.ContentTypeJSON, true)
	if err != nil {
		return nil, err
	}

	s := &models.KratosSession{}
	err = json.Unmarshal(resp, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// IsLoggedIn function for check current sessions from ory_kratos_session cookie
func IsLoggedIn(c *gin.Context) bool {
	cookies := c.Request.Cookies()

	for _, v := range cookies {
		if strings.Contains(v.Name, "ory_kratos_session") {
			return true
		}
	}

	return false
}

// Sessions for show sessions data on the browser or API
func Sessions(c *gin.Context) {
	s, err := GetCurrentSession(c)
	if err != nil {
		log.Println("[!] get current session ->", err.Error())
		return
	}

	if len(s.SID) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "empty session",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sid":      s.SID,
		"username": s.Identity.Traits.Username,
		"email":    s.Identity.Traits.Email,
	})
	return
}
