package main

type SecretResponse struct {
	Data struct {
		Data map[string]interface{} `json:"data"`
	} `json:"data"`
}
