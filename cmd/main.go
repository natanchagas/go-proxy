package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"gocloud.dev/server"
)

func main() {
	cert, err := tls.LoadX509KeyPair("localhost.crt", "localhost.key")
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		net.Dial("tcp", r.Host)

		log.Println("Received request:", r.Method, r.URL)

		r.RequestURI = ""
		response, err := http.DefaultClient.Do(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		io.Copy(w, response.Body)
		return
	})

	srv := server.NewDefaultDriver()
	srv.Server = http.Server{
		Addr:      ":8080",
		Handler:   mux,
		TLSConfig: config,
	}

	err = srv.ListenAndServeTLS(srv.Server.Addr, "localhost.crt", "localhost.key", mux)
	if err != nil {
		panic(err)
	}
}
