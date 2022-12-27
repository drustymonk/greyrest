package greyrest

import (
	"fmt"
	"net/http"
	"os"
)

func ApiGet(auth string, endpoint string) []byte {
	// create our request
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// add headers
	// Auth
	request.Header.Add("Authorization", auth)
	// Accept content types
	request.Header.Add("Accept", "application/json")

	return (HttpClient(request))
}

func ApiGetString(auth string, endpoint string) string {

	return (string(ApiGet(auth, endpoint)))

}
