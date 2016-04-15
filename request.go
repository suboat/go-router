package gorouter

type WsRequest struct {
	Tag           string
	RequestId     string
	Header        *Header
	Meta          *Meta
	Method        string
	URL           string
	Content       []byte
	ContentLength int64
	Data          interface{}
	Ignore        bool
}
