package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	. "github.com/YoungsoonLee/funtional-programming-in-go/2-design-patterns/ch05-decoration/02_decorator/decorator"
)

const (
	host     = "127.0.0.1"
	protocol = "http://"
)

var (
	serverUrl string
	proxyUrl  string
)

func init() {
	InitLog("trace-log.txt", ioutil.Discard, os.Stdout, os.Stderr)

	serverPort := 3000
	proxyPort := 8080
	flag.IntVar(&serverPort, "serverPort", serverPort, "Server Port")
	flag.IntVar(&proxyPort, "proxyPort", proxyPort, "Proxy Port")
	flag.Parse()
	serverUrl = fmt.Sprintf("%s:%d", host, serverPort)
	proxyUrl = fmt.Sprintf("%s:%d", host, proxyPort)

	Info.Printf("Metrics server listening on %s ", serverUrl)
	go func() {
		log.Fatal(easy_metrics.Serve(serverUrl))
	}()
	time.Sleep(1 * time.Second)

	req, err := http.NewRequest(http.MethodGet, protocol+serverUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}

	Info.Printf("Proxy listening on %s", proxyUrl)
	proxyUrl, _ := url.Parse(proxyUrl)
	tr := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	tr.TLSNextProto = make(map[string]func(string, *tls.Conn) http.RoundTripper)
	proxyTimeoutClient := &http.Client{Transport: tr, Timeout: 1 * time.Second}

	client := Decorate(
		proxyTimeoutClient,
		Authorization("mysecretpassword"),
		LoadBalancing(RoundRobin(0, "web01:3000", "web02:3000", "web03:3000")),
		Logging(log.New(InfoHandler, "client: ", log.Ltime)),
		FaultTolerance(2, time.Second),
	)

	job := &dec.Job{
		Client:       client,
		Request:      req,
		NumRequests:  10,
		IntervalSecs: 10,
	}

	start := time.Now()
	job.Run()
	dec.Info.Printf("\n>> It took %s", time.Since(start))

	dec.Info.Printf("metrics")
	err = easy_metrics.DisplayResults(serverUrl)
	if err != nil {
		log.Fatalln(err)
	}

	dec.Info.Printf("CTRL+C to exit")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
