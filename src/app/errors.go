package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
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

func checkAndConvertToInt(variable string, form map[string][]string) (int, error) {
	var err error = nil
	varSlice, varPresent := form[variable]
	if (varPresent == false) || (len(varSlice) == 0) {
		err = errors.New(variable + " not defined")
		return 0, err
	}
	varValue, errConvert := strconv.Atoi(varSlice[0])
	if errConvert != nil {
		err = errors.New(variable + " unable to be converted to integer")
		return 0, err
	}
	return varValue, nil
}
