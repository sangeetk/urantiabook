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

	var ub = s.UrantiaBook{}

	r := mux.NewRouter()

	r.Handle("/index", h.NewServer(s.MakeIndexEndpoint(ub), s.DecodeIndexRequest, s.EncodeResponse))
	r.Handle("/parts", h.NewServer(s.MakePartsEndpoint(ub), s.DecodePartsRequest, s.EncodeResponse))
	r.Handle("/paper", h.NewServer(s.MakePaperEndpoint(ub), s.DecodePaperRequest, s.EncodeResponse))

	http.ListenAndServe(fmt.Sprintf(":%d", port), r)

}
