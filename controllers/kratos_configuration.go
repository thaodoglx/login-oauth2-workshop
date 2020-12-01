package controllers

import (
	"log"
	"net/url"
	"os"

	runtime "github.com/go-openapi/runtime/client"
	kratos "github.com/ory/kratos-client-go/client"
)

var (
	// KratosAdminAPI is Kratos administrative endpoint
	KratosAdminAPI string

	// KratosPublicAPI is Kratos public endpoint
	KratosPublicAPI string

	// KratosClient sdk
	KratosClient *kratos.OryKratos
)

// LoadKratosConfiguration for load all kratos configuration variables from .env files and define several values.
// Run this function in built-in `init` function
func LoadKratosConfiguration() {
	KratosAdminAPI = os.Getenv("KRATOS_ADMIN_API")
	KratosPublicAPI = os.Getenv("KRATOS_PUBLIC_API")

	kratosAdminHost, err := url.Parse(KratosAdminAPI)
	if err != nil {
		log.Println("[!] url parse error ->", err.Error())
		return
	}

	kratosPublicHost, err := url.Parse(KratosPublicAPI)
	if err != nil {
		log.Println("[!] url parse error ->", err.Error())
		return
	}

	KratosClient = kratos.NewHTTPClientWithConfig(
		nil,
		&kratos.TransportConfig{
			Schemes:  []string{kratosAdminHost.Scheme},
			Host:     kratosAdminHost.Host,
			BasePath: kratosAdminHost.Path,
		},
	)

	publicTransport := runtime.New(
		kratosPublicHost.Host,
		kratosPublicHost.Path,
		[]string{kratosPublicHost.Scheme},
	)
	KratosClient.Public.SetTransport(publicTransport)
}
