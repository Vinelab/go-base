package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// GET issues a GET request to the given "url",
// parses the response into a JSON structure
// according to the given "responseMap"
func GET(url string, responseMap interface{}) {
	c := &http.Client{
		Timeout: 10 * time.Minute,
	}

	res, err := c.Get(url)
	if err != nil {
		err := errors.New(fmt.Sprintf("%v in url: %s", err, url))

		fmt.Println(err)
		return
	}

	body, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(body, responseMap)

	if res.StatusCode != 200 {
		showErrorResponse(body, url)
	}

	res.Body.Close()
}

// POST issues a POST request to the given "url",
// parses the response into a JSON structure
// according to the given "responseMap"
func POST(url string, content map[string]interface{}, responseMap interface{}) {

	payload, _ := json.Marshal(content)

	c := &http.Client{
		Timeout: 10 * time.Minute,
	}

	res, err := c.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(fmt.Sprintf("%v in url: %s", err, url))
		return
	}

	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, responseMap)

	if res.StatusCode != 200 {
		showErrorResponse(body, url)
	}

	res.Body.Close()
}

func showErrorResponse(body []byte, url string) {
	fmt.Println(fmt.Sprintf("%v in url: %s", string(body), url))
}
