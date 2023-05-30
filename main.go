package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// Configure the proxy server
	proxy := NewProxyServer("http://localhost:8080")

	// Start the proxy server
	log.Println("Proxy server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", proxy))
}

// ProxyServer represents an HTTP proxy server
type ProxyServer struct {
	proxy *httputil.ReverseProxy
}

// NewProxyServer creates a new instance of ProxyServer
func NewProxyServer(targetURL string) *ProxyServer {
	target, err := url.Parse(targetURL)
	if err != nil {
		log.Fatal("Failed to parse target URL:", err)
	}

	return &ProxyServer{
		proxy: httputil.NewSingleHostReverseProxy(target),
	}
}

// ServeHTTP handles the incoming HTTP requests
func (p *ProxyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Proxying request:", r.Method, r.URL.Path)

	// Modify the request if needed (e.g., add headers, modify URL)

	// Proxy the request to the target server
	p.proxy.ServeHTTP(w, r)

	// Modify the response if needed (e.g., add headers)

	log.Println("Request completed:", r.Method, r.URL.Path)
}
