package main

import (
	"log"
	"net/http"
	"strings"
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

func stringToSlice(valString string) []string {
	replacer := strings.NewReplacer("[", "", "]", "", "\"", "")
	cleaned := replacer.Replace(valString)
	values := strings.Split(cleaned, ",")
	if (len(values) == 1) && (strings.Compare(values[0], "") == 0) {
		return []string{}
	}
	return values
}
