package utility

import (
	"bytes"
	"cb-ldp-backend/models/response"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

func HttpRequest(methodType string, url string, payload interface{}, headers map[string]string) (*http.Response, error) {
	var reqBody []byte
	client := &http.Client{}

	req, err := http.NewRequest(methodType, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ParseResponseBody(response *http.Response) (map[string]string, error) {
	defer response.Body.Close()
	var responseBody map[string]string
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(respBody, &responseBody)
	return responseBody, nil
}

func ProcessZohoResponse(rawJsonString string) (response.ZohoResponse, error) {
	var final response.ZohoResponse
	jsonResponseMessage := gjson.Get(rawJsonString, "response.message")
	log.Println("Zoho Api Response Message : ", jsonResponseMessage)
	if jsonResponseMessage.String() == "Data fetched successfully" {
		jsonT := gjson.Get(rawJsonString, "response.result.0")
		if len(jsonT.String()) == 0 {
			err := errors.New("Parsing Zoho Response Issue")
			log.Println("Error :", err)
			return final, err
		}
		jsonT.ForEach(func(key, value gjson.Result) bool {
			_ = json.Unmarshal([]byte(value.String()), &final)
			return false // keep iterating
		})
		return final, nil
	} else {
		return final, errors.New("User Verification Failed : No Zoho Record Found")
	}

}
