package router

import (
	"encoding/json"
	. "github.com/suboat/go-router"
)

type WSResponse struct {
	Tag       string
	RequestId string
	Meta      *Meta
	Method    string
	URL       string
	Data      interface{}
	Error     string
	Success   bool
}

func NewWSResponse(req *WSRequest) (*WSResponse, error) {
	if req == nil {
		return nil, ErrWSRequest
	}
	return &WSResponse{
		Tag:       req.Tag,
		RequestId: req.RequestId,
		Meta:      req.Meta,
		Method:    req.Method,
		URL:       req.URL,
	}, nil
}

func (s *WSResponse) Valid() bool {
	return len(s.Method) != 0 && len(s.URL) != 0
}

func (s *WSResponse) Bytes() ([]byte, error) {
	return json.Marshal(s)
}

func BytesToWSResponse(data []byte) (*WSResponse, error) {
	var r WSResponse
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
