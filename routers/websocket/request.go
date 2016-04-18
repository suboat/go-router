package router

import (
	"encoding/json"
	. "github.com/suboat/go-router"
)

type WSRequest struct {
	Tag       string
	RequestId string
	Header    *Header
	Meta      *Meta
	Method    string
	URL       string
	Data      interface{}
	Ignore    bool
}

func (s *WSRequest) Valid() bool {
	return len(s.Method) != 0 && len(s.URL) != 0
}

func (s *WSRequest) Bytes() ([]byte, error) {
	return json.Marshal(s)
}

func BytesToWSRequest(data []byte) (*WSRequest, error) {
	var r WSRequest
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
