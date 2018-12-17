package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func setupHandler(cfg map[string]float64) {
	http.HandleFunc("/setup", func(w http.ResponseWriter, r *http.Request) {
	    switch r.Method {
        case "GET":
            setupBytes, _ := json.Marshal(cfg)
            w.Write(setupBytes)
    		return
        case "POST":
            errP := r.ParseForm()
    		if errP != nil {
    			log.Println(errP)
    			w.Write([]byte(""))
    			return
    		}
    		setupStr, setupPres := r.Form["Setup"]
    		if setupPres == false {
    			log.Println("Setup not present")
    			w.Write([]byte(""))
    			return
    		}
    		// Must pass explicitly by reference in this case //
    		errJSON := json.Unmarshal([]byte(setupStr[0]), &cfg)
    		if errJSON != nil {
    			log.Println(errJSON)
    			w.Write([]byte(""))
    			return
    		}
    		errSave := saveConfigToFile("setup.cfg", cfg)
    		if errSave != nil {
    		    log.Println(errSave)
    		    w.Write([]byte(""))
    			return
    		}
    		log.Println("Setup updated.")
    		w.Write([]byte(""))
    		return
        default:
            log.Println("Only GET and POST methods are supported.")
            w.Write([]byte(""))
			return
	    }
	})
}
