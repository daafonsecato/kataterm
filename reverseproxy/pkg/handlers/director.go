package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (controller *SessionController) ReverseProxyHandler(rw http.ResponseWriter, req *http.Request) {
	// Extract UUID and service name from the request host
	targetURL, err := ExtractUUIDAndServiceName(req.Host)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(rw, err)
		return
	}
	// Use the new host if it's different from targetURL
	if targetURL != nil {
		req.URL.Host = targetURL.Host
	}

	req.URL.Scheme = targetURL.Scheme
	req.RequestURI = ""
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(rw, err)
		return
	}
	for key, values := range resp.Header {
		for _, value := range values {
			rw.Header().Set(key, value)
		}
	}
	rw.WriteHeader(resp.StatusCode)
	io.Copy(rw, resp.Body)
}

func ExtractUUIDAndServiceName(host string) (*url.URL, error) {
	var targetURL *url.URL
	var err error

	// Extract the first subdomains from the host
	subdomains := strings.Split(host, ".")
	if len(subdomains) < 3 {
		log.Fatalf("Invalid host: %s", host)
		return nil, err
	}

	// Extract the UUID and service name using regular expressions
	uuid := subdomains[0]
	service := subdomains[1]

	switch service {
	case "backend":
		targetURL, err = url.Parse("http://backend-svc-" + uuid + ":8000")
	case "ttyd":
		targetURL, err = url.Parse("http://ttyd-svc-" + uuid + ":7681")
	case "codeeditor":
		targetURL, err = url.Parse("http://codeeditor-svc-" + uuid + ":8080")
	default:
		return nil, fmt.Errorf("unknown service type: %s", service)
	}

	if len(uuid) < 2 || len(service) < 2 {
		log.Fatalf("Invalid UUID or service name: %s", host)
		return nil, err
	}

	return targetURL, err
}
