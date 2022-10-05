package main

import (
	"encoding/json"
	"fmt"
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

type ScriptResponse struct {
	Script string `json:"script"`
}

type ScriptResponseError struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
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

			if res.StatusCode != http.StatusOK {
				scriptError := ScriptResponseError{}
				err := json.NewDecoder(res.Body).Decode(&scriptError)
				if err != nil {
					log.Printf("failed to decode response body into scriptError: %s\n", err)
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(scriptError.Message))
				return
			}

			scriptResponse := ScriptResponse{}
			err = json.NewDecoder(res.Body).Decode(&scriptResponse)
			if err != nil {
				log.Printf("failed to decode response body into scriptResponse: %s\n", err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			scriptResponse.Script = strings.Replace(scriptResponse.Script, "base64:", "", 1)
			w.Write([]byte(scriptResponse.Script))

		}

	})

	log.Println("Starting server at port 8085")
	if err := http.ListenAndServe(":8085", mux); err != nil {
		log.Fatal(err)
	}
}
