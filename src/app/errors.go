package main

import (
	"log"
	"net/http"
)

func errorResponse(err error, w *http.ResponseWriter) {
	if err != nil {
		log.Println(err)
		(*w).Write([]byte(""))
	}
	return
}
