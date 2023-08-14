package utils

import (
	"io"
	"log"
	"net/http"
)

func GetDataHTTP(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Unable to get the response")
	}

	responseByte, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Unable to read the response")
	}

	return responseByte
}
