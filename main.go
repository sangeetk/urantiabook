package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	s "git.urantiatech.com/urantiabook/urantiabook/service"
	"github.com/gorilla/mux"
	h "github.com/urantiatech/kit/transport/http"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Port number")
	flag.Parse()

	log.SetFlags(log.Lshortfile)

	if os.Getenv("PORT") != "" {
		p, err := strconv.ParseInt(os.Getenv("PORT"), 10, 32)
		if err == nil {
			port = int(p)
		}
	}

	var svc s.Page
	svc = s.Page{}

	r := mux.NewRouter()

	r.Handle("/list", h.NewServer(s.MakeListEndpoint(svc), s.DecodeListRequest, s.EncodeResponse))
	r.Handle("/read", h.NewServer(s.MakeReadEndpoint(svc), s.DecodeReadRequest, s.EncodeResponse))
	r.Handle("/search", h.NewServer(s.MakeSearchEndpoint(svc), s.DecodeSearchRequest, s.EncodeResponse))

	http.ListenAndServe(fmt.Sprintf(":%d", port), r)

}
