package controllers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/ory/kratos-client-go/client/common"
)

// Registration controllers for register new user
func Registration(c *gin.Context) {
	// get request query
	getRequest := c.Query("request")

	// if request query is empty, initialize browser-based registration flow
	if len(getRequest) == 0 {
		re := fmt.Sprintf("%s%sself-service/browser/flows/registration", LoginProviderCompleteHost, LoginProviderSelfPublicAPIPath)
		c.Redirect(http.StatusTemporaryRedirect, re)
	}

	params := common.NewGetSelfServiceBrowserRegistrationRequestParams()
	params.WithRequest(getRequest)

	resp, err := KratosClient.Common.GetSelfServiceBrowserRegistrationRequest(params)
	if err != nil {
		log.Println("[!] get self-service browser registration request error ->", err.Error())
		return
	}

	config := resp.GetPayload().Methods["password"].Config

	tpl := template.Must(template.ParseFiles("views/registration.html"))
	err = tpl.Execute(c.Writer, config)
	if err != nil {
		log.Println("[!] template execute error ->", err.Error())
		return
	}
}
