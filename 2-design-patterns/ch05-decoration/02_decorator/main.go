package main

import "flag"

const (
	host     = "127.0.0.1"
	protocol = "http://"
)

var (
	serverUrl string
	proxyUrl  string
)

func init() {
	serverPort := 3000
	proxyPort := 8080
	flag.IntVar(&serverPort, "serverPort", serverPort, "Server Port")
	flag.IntVar(&proxyPort, "proxyPort", proxyPort, "Proxy Port")
	flag.Parse()
	serverUrl = fmt.Sprinf("%s:%d", host, serverPort)
	proxyUrl = fmt.Sprintf("%s:%d", host, proxyPort)
}
