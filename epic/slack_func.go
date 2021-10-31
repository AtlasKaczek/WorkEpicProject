package epic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendSlackMessege(msg string, url string) (*[]byte, error) {

	postBody, _ := json.Marshal(map[string]string{
		"text": msg,
	})

	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		fmt.Printf("SendSlackMessege[1]: An error occured %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("SendSlackMessege[2]: An error occured %v", err)
	}

	return &body, nil
}
