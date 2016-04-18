package router

import (
	"encoding/json"
	. "github.com/suboat/go-router"
	"net/http"
)

type WSResponse struct {
	ResponseWriter http.ResponseWriter `json:"-"`
	Tag            string
	RequestId      string
	Meta           *Meta
	Methods        []string
	URL            string
	Data           interface{}
	Error          string
	Success        bool
}

func NewWSResponse(req *WSRequest, w http.ResponseWriter) (*WSResponse, error) {
	if req == nil {
		return nil, ErrWSRequest
	}
	return &WSResponse{
		ResponseWriter: w,
		Tag:            req.Tag,
		RequestId:      req.RequestId,
		Meta:           req.Meta,
		Methods:        req.Methods,
		URL:            req.URL,
	}, nil
}

func (s *WSResponse) Valid() bool {
	return len(s.Methods) != 0 && len(s.URL) != 0
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
