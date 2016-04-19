package router

import (
	"encoding/json"
	. "github.com/suboat/go-router"
	"net/http"
)

type WSResponse struct {
	ResponseWriter http.ResponseWriter `json:"-"`
	Tag            string
}

func NewWSResponse(req *WSRequest, w http.ResponseWriter) (*WSResponse, error) {
	if req == nil {
		return nil, ErrWSRequest
	}
	return &WSResponse{
		ResponseWriter: w,
		Tag:            req.Tag,
	}, nil
}

func (s *WSResponse) IsHTTP() bool {
	return false
}

func (s *WSResponse) Valid() bool {
	return len(s.Tag) != 0
}

func (s *WSResponse) Bytes() ([]byte, error) {
	return json.Marshal(s)
}

func (s *WSResponse) FromBytes(data []byte) error {
	if err := json.Unmarshal(data, s); err != nil {
		return err
	}
	return nil
}

func (s *WSResponse) GetResponseWriter() http.ResponseWriter {
	return s.ResponseWriter
}
