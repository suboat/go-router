package gorouter

import "encoding/json"

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
