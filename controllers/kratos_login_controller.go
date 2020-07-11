package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ory/kratos-client-go/client/common"
)

// Login controller for login flow
func Login(c *gin.Context) {
	// get request query
	getRequest := c.Query("request")

	// if request query is empty, initialize browser-based registration flow
	if len(getRequest) == 0 {
		re := fmt.Sprintf("%s%sself-service/browser/flows/login", LoginProviderCompleteHost, LoginProviderSelfPublicAPIPath)
		c.Redirect(http.StatusTemporaryRedirect, re)
	}

	params := common.NewGetSelfServiceBrowserLoginRequestParams()
	params.WithRequest(getRequest)

	resp, err := AdminKratosClient.Common.GetSelfServiceBrowserLoginRequest(params)
	if err != nil {
		log.Println("[!] get self-service browser login request error ->", err.Error())
		return
	}

	config := resp.GetPayload().Methods["password"].Config

	tpl := template.Must(template.ParseFiles("views/login.html"))
	err = tpl.Execute(c.Writer, config)
	if err != nil {
		log.Println("[!] template execute error ->", err.Error())
		return
	}
}
