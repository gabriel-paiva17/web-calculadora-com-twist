package main

import (
	"log"
	"net/http"
	"time"
)

type MainHandler struct {
}

func (h MainHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	
	if !isMethodAndPathValid(req) {
	
		res.Write(MessageToJson(TableMessages[404]))
		return

	}
	
	ResultController[req.Method][req.URL.Path](res, req)
	
}

func CreateServer() {

	s := &http.Server{
		Addr:         "172.22.51.31:8080",
		Handler:      MainHandler{},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	
	log.Fatal(s.ListenAndServe())

}


func isMethodAndPathValid(req *http.Request)bool {

	return ResultController[req.Method][req.URL.Path] != nil

}

