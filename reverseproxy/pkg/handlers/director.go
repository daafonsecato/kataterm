package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"net/http/httputil"
	"net/url"
)

func NewReverseProxy(targetHost string, targetPort string) *httputil.ReverseProxy {
	target := &url.URL{
		Scheme: "http",
		Host:   targetHost + ":" + targetPort,
	}

	return httputil.NewSingleHostReverseProxy(target)
}

func (controller *SessionController) CustomDirector() *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			host := req.Host
			subdomains := strings.Split(host, ".")

			// Assuming the format is [sessionID].[service].terminal.kataterm.com
			if len(subdomains) < 4 {
				log.Println("Invalid host:", host)
				return // Or set a default targetHost
			}
			fmt.Println(subdomains)
			sessionID := subdomains[0]
			service := subdomains[1]
			machineID, err := controller.sessionStore.GetAWSInstanceID(sessionID)
			if err != nil {
				log.Println("Error getting machine ID:", err)
				return // Or set a default targetHost
			}
			// Get the host for the machine
			targetHost, err := controller.sessionStore.GetMachineHost(machineID)
			if err != nil {
				log.Println("Error getting machine host:", err)
				return // Or set a default targetHost
			}

			var port string
			switch service {
			case "backend":
				port = "8000"
			case "ttyd":
				port = "7681"
			case "code-server":
				port = "8080"
			// Add more cases as needed for different services
			default:
				log.Println("Service not recognized:", service)
				return // Or set a default port
			}

			// Modify the request to route to the target host and port
			req.URL.Host = targetHost + ":" + port
		},
	}
}
