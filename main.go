package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func getEnvUrl(envName string, defaultValue string) *url.URL {
	urlString := os.Getenv(envName)
	if urlString == "" {
		if defaultValue != "" {
			urlString = defaultValue
		} else {
			log.Fatal(envName + " is required")
		}
	}
	urlPtr, err := url.Parse(urlString)
	if err != nil {
		log.Fatal(err)
	}
	return urlPtr
}

func getEnvPort(envName string, defaultValue string) string {
	portString := os.Getenv(envName)
	if portString == "" {
		if defaultValue != "" {
			portString = defaultValue
		} else {
			log.Fatal(envName + " is required")
		}
	}
	return portString
}

func main() {
	log.Println("Starting the server...")

	var port = getEnvPort("PORT", "8080")

	var proxyUrl = getEnvUrl("PROXY_URL", "https://api.openai.com")

	reverseProxy := httputil.NewSingleHostReverseProxy(proxyUrl)

	originalDirector := reverseProxy.Director

	reverseProxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Header.Del("Proxy-Connection")
		req.Header.Del("Proxy-Authenticate")
		req.Header.Del("Proxy-Authorization")
		req.Header.Del("Via")
		req.Header.Del("X-Forwarded-For")
		req.URL.Scheme = proxyUrl.Scheme
		req.URL.Host = proxyUrl.Host
		req.Host = proxyUrl.Host

		dump, err := httputil.DumpRequestOut(req, true)
		if err == nil {
			log.Printf("\r\n%s\r\n%s", strings.Repeat("-", 50), dump)
		} else {
			log.Printf("Error dumping request: %s", err)
		}

	}

	var extProxyUrl = getEnvUrl("EXT_PROXY_URL", "http://10.0.0.8:8080")
	reverseProxy.Transport = &http.Transport{Proxy: http.ProxyURL(extProxyUrl)}

	log.Printf("Proxying to %s", proxyUrl)
	log.Printf("Listening on :%s", port)
	log.Printf("Using external proxy %s", extProxyUrl)
	log.Fatal(http.ListenAndServe(":"+port, reverseProxy))
}
