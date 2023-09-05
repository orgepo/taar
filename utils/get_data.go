package utils

import (
	"context"
	"io"
	"log"
	"net/http"
)

func GetDataHTTP(URL string) []byte {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, URL, io.MultiReader())

	response, err := http.DefaultClient.Do(req) //#nosec
	if err != nil {
		log.Println("Unable to get the response")
	}
	defer response.Body.Close()

	responseByte, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Unable to read the response")
	}

	return responseByte
}
