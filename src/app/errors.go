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

func absenceResponse(presence bool, ID string, w *http.ResponseWriter) {
	if presence == false {
		log.Println(ID + " not present in post request")
		(*w).Write([]byte(""))
	}
	return
}

func invalidValueResponse(warning string, w *http.ResponseWriter) {
	log.Println(warning)
	(*w).Write([]byte(""))
	return
}
