package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	serve()
}

func serve() {
	hosting8080 := make(chan bool)

	go func() {
		director := func(request *http.Request) {
			fmt.Println("Launch Reverse Proxy")
			request.URL.Scheme = "https"
			request.URL.Host = "hidexir.github.io"
			request.Host = "hidexir.github.io"
		}
		rp := &httputil.ReverseProxy{Director: director}
		server := http.Server{
			Addr:    ":8080",
			Handler: rp,
		}
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err.Error())
		}
		hosting8080 <- true
	}()

	// チャンネルの呼び出し
	<-hosting8080
}
