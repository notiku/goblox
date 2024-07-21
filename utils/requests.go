package utils

import (
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
