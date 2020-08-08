package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// GetRandomNumber gets random number from back-service
func GetRandomNumber() (order map[string]interface{}, err error) {
	log.Printf("Back server requested GetRandomNumber")

	resp, err := http.Get(fmt.Sprintf("%s/randomNumber", os.Getenv("BACK_SVC_ADDRESS")))
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &order)
	if err != nil {
		return
	}

	log.Printf("Back server GetRandomNumber answer %+v", resp)
	return
}
