package greyrest

import (
	"fmt"
	"net/http"
	"os"
)

func ApiDelete(auth string, endpoint string) []byte {
	// create our request
	request, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} // add headers
	// Auth
	request.Header.Add("Authorization", auth)
	// Accept content types
	request.Header.Add("Accept", "application/json")

	return (HttpClient(request))
}

func ApiDeleteString(auth string, endpoint string) string {

	return (string(ApiDelete(auth, endpoint)))

}
