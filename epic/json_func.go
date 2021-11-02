package epic

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetJSON(URL string) (*[]byte, error) {

	res, getErr := http.Get(URL)
	if getErr != nil {
		fmt.Printf("GetJson 1: An error occured: %v", getErr)
		return nil, getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	if res.StatusCode != 200 {
		err := fmt.Sprintf("GetJson 2: Bad StatusCode: %d", res.StatusCode)
		return nil, errors.New(err)
	}

	body, _ := ioutil.ReadAll(res.Body)

	return &body, nil
}

func ParseJSON(url string) (Games, error) {
	var freeGame Games

	res, err := GetJSON(url)
	if err != nil {
		fmt.Printf("ParseJSON 1: An error occured: %v", err)
		return Games{}, err
	}

	jsonErr := json.Unmarshal(*res, &freeGame)
	if jsonErr != nil {
		fmt.Printf("ParseJSON 1: An error occured: %v", jsonErr)
		return Games{}, jsonErr
	}
	return freeGame, nil
}
