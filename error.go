package gorouter

import "errors"

var (
	ErrNil        error = errors.New("gorouter: DON`T supoort nil")
	ErrRouter           = errors.New("gorouter: DON`T supoort the router or is nil")
	ErrHandle           = errors.New("gorouter: DON`T supoort the handle or is nil")
	ErrHandleFunc       = errors.New("gorouter: DON`T supoort the handle function or is nil")
	ErrWSRequest        = errors.New("gorouter: DON`T supoort the websocket request or is nil")
)
