package main

import (
	"log"
	"net/http"
	"net/url"
	"time"
)

// Drugula Market URL
var webUrl string = "http://drugula44brpin5w2xk6j3adll35pp3zdihqqkndhg336j55pwdtufyd.onion/"

// Tor proxy
var torProxy string = "socks5://127.0.0.1:9050"

func main() {
	for {
		go statusCheck()
		time.Sleep(time.Second * 60)
	}
}

func statusCheck() {
	// Parse Tor proxy URL string to a URL type
	torProxyUrl, err := url.Parse(torProxy)
	if err != nil {
		log.Print("Error parsing Tor proxy URL:", torProxy, ".", err)
	}

	// Set up a custom HTTP transport to use the proxy and create the client
	torTransport := &http.Transport{Proxy: http.ProxyURL(torProxyUrl)}
	client := &http.Client{Transport: torTransport, Timeout: time.Second * 10}

	// Make request
	resp, err := client.Get(webUrl)
	if err != nil {
		log.Print("Error making GET request.", err)
	}
	defer resp.Body.Close()

	log.Println("Drugula Market status code:", resp.StatusCode)
}
