package api

import (
	"encoding/json"
	"io"
)

type ServerMessage struct {
	Message              string        `json:"message"`
	MessageArgs          []interface{} `json:"messageArgs"`
	IsHangarAPIException bool          `json:"isHangarApiException"`
	HTTPError            struct {
		StatusCode   int    `json:"statusCode"`
		StatusPhrase string `json:"statusPhrase"`
	} `json:"httpError"`
}

func decodeServerMessage(r io.Reader) (ServerMessage, error) {
	var data ServerMessage
	err := json.NewDecoder(r).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
