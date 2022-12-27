package greyrest

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func ApiPost(auth string, endpoint string) []byte {
	// get our body
	//data := url.Values{}
	// create our request
	//request, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	request, err := http.NewRequest("POST", endpoint, strings.NewReader("{ }"))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	request.Header.Add("Authorization", auth)

	/*
	 */
	return (HttpClient(request))
}

func ApiPostString(auth string, endpoint string) string {

	return (string(ApiPost(auth, endpoint)))

}
