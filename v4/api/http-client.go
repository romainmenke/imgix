package api

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HTTPClient returns an authenticated http Client
func HTTPClient(config *Config) (*http.Client, error) {
	if config == nil {
		return nil, errors.New("Authentication failed")
	}

	if config.APIKey == "" {
		apiKey, err := login(config.Email, config.Password)
		if err != nil {
			return nil, err
		}

		config.APIKey = apiKey
	}

	transport, err := authRoundTripper(http.DefaultTransport, config.APIKey)
	if err != nil {
		return nil, err
	}

	return &http.Client{
		Transport: transport,
	}, nil
}

type roundTripper func(req *http.Request) (*http.Response, error)

func (x roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return x(req)
}

func authRoundTripper(next http.RoundTripper, apiKey string) (http.RoundTripper, error) {
	return roundTripper(func(req *http.Request) (*http.Response, error) {
		req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(apiKey+":"))))

		return next.RoundTrip(req)
	}), nil
}

func login(email string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	payload := map[string]map[string]interface{}{
		"data": map[string]interface{}{
			"type": "users",
			"attributes": map[string]string{
				"email":    email,
				"password": password,
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", errors.New("Authentication failed")
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.imgix.com/v4/users/login", bytes.NewBuffer(body))
	if err != nil {
		return "", errors.New("Authentication failed")
	}

	req.Header.Set("Content-Type", "application/vnd.api+json")

	req = req.Clone(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", errors.New("Authentication failed")
	}

	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("Authentication failed")
	}

	respPayload := map[string]interface{}{}
	err = json.Unmarshal(respData, &respPayload)
	if err != nil {
		return "", errors.New("Authentication failed")
	}

	if data, ok := respPayload["data"].(map[string]interface{}); !ok {
		return "", errors.New("Authentication failed")
	} else if attributes, ok := data["attributes"].(map[string]interface{}); !ok {
		return "", errors.New("Authentication failed")
	} else if auth, ok := attributes["api_key"].(string); !ok {
		return "", errors.New("Authentication failed")
	} else if auth != "" {
		return auth, nil
	}

	return "", errors.New("Authentication failed")
}
