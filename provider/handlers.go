package provider

import (
	"bytes"
	"log"
	"net/http"
	"strings"
)

func postHandler(input []byte, auth string, instance string, endpoint string) int {

	bearer := "Bearer " + auth
	req, _ := http.NewRequest(
		"POST",
		instance+endpoint,
		bytes.NewBuffer(input),
	)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	res, err := http.DefaultClient.Do(req) // send an HTTP using `req` object
	if err != nil {                        // check for response error
		log.Fatal("Error:O", err)
	}
	defer res.Body.Close()
	return res.StatusCode
}

func deleteHandler(auth string, instance string, endpoint string) error {

	bearer := "Bearer " + auth
	req, _ := http.NewRequest(
		"DELETE",
		instance+endpoint,
		strings.NewReader(""),
	)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	res, err := http.DefaultClient.Do(req) // send an HTTP using `req` object
	if err != nil {                        // check for response error
		log.Fatal("Error:O", err)
	}
	defer res.Body.Close()
	log.Printf("POST status: %d\n", res.StatusCode)
	return err
}

func (jgs *JiraGroupState) appendGroupstate(val int) {
	jgs.ReturnCode = val
}

func (jps *JiraProjectState) appendProjectstate(val int) {
	jps.ReturnCode = val
}
