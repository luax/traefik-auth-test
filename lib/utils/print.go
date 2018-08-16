package utils

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func PrintEnvVariables() {
	log.Println(">>> Environment variables <<<")
	for _, e := range os.Environ() {
		log.Println(e)
	}
	log.Println(">>> end <<<")
}

func PrintRequest(req *http.Request) {
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(">>> Request <<<")
	log.Println(string(requestDump))
	log.Println(">>> end <<<")
}
