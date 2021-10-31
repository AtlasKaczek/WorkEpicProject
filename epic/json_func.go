package epic

import (
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
