package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// DoRequest makes a request to the requested endpoint.
func DoRequest(method string, endpointUrl string, params map[string]string, result interface{}) error {
	params["format"] = "json"

	u, err := url.Parse(endpointUrl)
	if err != nil {
		return err
	}

	q := u.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}

	return nil
}

// DoRequest2 makes a request to the requested endpoint.
func DoRequest2(method string, endpointUrl string, headers map[string]string, params map[string]interface{}, result interface{}) error {
	u, err := url.Parse(endpointUrl)
	if err != nil {
		return err
	}

	// If method is GET, add params to query string
	if method == http.MethodGet {
		q := u.Query()
		for key, value := range params {
			switch v := value.(type) {
			case string:
				q.Set(key, v)
			case []string:
				for _, val := range v {
					q.Add(key, val)
				}
			}
		}
		u.RawQuery = q.Encode()
	}

	// Marshal params into JSON if method is POST or PUT
	var jsonBody []byte
	if method == http.MethodPost || method == http.MethodPut {
		jsonBody, err = json.Marshal(params)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	// Set headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}

	return nil
}
