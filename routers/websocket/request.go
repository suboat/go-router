package router

import (
	"encoding/json"
	"net/http"
)

type WSRequest struct {
	Request *http.Request `json:"-"`
	Tag     string
}

func NewWSRequest(r *http.Request) *WSRequest {
	return &WSRequest{Request: r}
}

func (s *WSRequest) IsHTTP() bool {
	return false
}

func (s *WSRequest) Valid() bool {
	return len(s.Tag) != 0
}

func (s *WSRequest) Bytes() ([]byte, error) {
	return json.Marshal(s)
}

func (s *WSRequest) FromBytes(data []byte) error {
	if err := json.Unmarshal(data, s); err != nil {
		return err
	}
	return nil
}

func (s *WSRequest) GetRequest() *http.Request {
	return s.Request
}
