package Main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// submit http request and handle response
func HttpClient(request *http.Request) []byte {
	myClient := &http.Client{}
	response, err := myClient.Do(request)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	// make sure we close the client
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error while reading response: ", err)
	}
	if response.StatusCode != 200 {
		fmt.Println("Got response ", response.StatusCode)
		fmt.Println(string(responseData))
		os.Exit(1)
	}
	return (responseData)
}
