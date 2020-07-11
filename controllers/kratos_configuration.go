package controllers

import (
	"log"
	"net/url"
	"os"

	kratos "github.com/ory/kratos-client-go/client"
)

var (
	// KratosAdminAPI is Kratos administrative endpoint
	KratosAdminAPI string

	// KratosPublicAPI is Kratos public endpoint
	KratosPublicAPI string

	// AdminKratosClient sdk
	AdminKratosClient *kratos.OryKratos

	// PublicKratosClient sdk
	PublicKratosClient *kratos.OryKratos
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

	AdminKratosClient = kratos.NewHTTPClientWithConfig(
		nil,
		&kratos.TransportConfig{
			Schemes:  []string{kratosAdminHost.Scheme},
			Host:     kratosAdminHost.Host,
			BasePath: kratosAdminHost.Path,
		},
	)

	PublicKratosClient = kratos.NewHTTPClientWithConfig(
		nil,
		&kratos.TransportConfig{
			Schemes:  []string{kratosPublicHost.Scheme},
			Host:     kratosPublicHost.Host,
			BasePath: kratosPublicHost.Path,
		},
	)
}
