package router

import (
	"encoding/json"
	. "github.com/suboat/go-router"
)

type WebSocketRequest struct {
	Tag       string
	RequestId string
	Header    *Header
	Meta      *Meta
	Method    string
	URL       string
	Data      interface{}
	Ignore    bool
}

func (s *WebSocketRequest) Bytes() ([]byte, error) {
	return json.Marshal(s)
}

func BytesToWSRequest(data []byte) (*WebSocketRequest, error) {
	var r WebSocketRequest
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
