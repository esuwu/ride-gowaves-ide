package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	gowavesCompiler = "GowavesCompiler"
	scalaCompiler   = "WavesCompiler"
)

type Script struct {
	Code     string `json:"code"`
	Compiler string `json:"compiler"`
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/app/compile", func(w http.ResponseWriter, r *http.Request) {
		var script Script
		lala := r.URL.Query().Get("code")
		fmt.Println(lala)
		err := json.NewDecoder(r.Body).Decode(&script)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

		if script.Compiler == gowavesCompiler {
			w.WriteHeader(http.StatusOK)
			result := "Gowaves compiler is not supported yet"
			w.Write([]byte(result))
		}

		if script.Compiler == scalaCompiler {
			req, err := http.NewRequest("POST", "http://mainnet-statehash-aws-fr-1.wavesnodes.com/utils/script/compileCode",
				strings.NewReader(script.Code))
			if err != nil {
				log.Printf("failed to could not create request: %s\n", err)
			}
			req.Header.Set("Content-Type", "text/plain")
			req.Header.Set("accept", "application/json")
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Printf("failed to send an http request: %s\n", err)
			}

			resBody, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Printf("client: could not read response body: %s\n", err)
			}
			w.WriteHeader(http.StatusOK)
			log.Println(res.StatusCode)
			w.Write(resBody)
		}

	})

	log.Println("Starting server at port 8085")
	if err := http.ListenAndServe(":8085", mux); err != nil {
		log.Fatal(err)
	}
}
