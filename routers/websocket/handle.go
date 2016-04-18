package router

import (
	. "github.com/suboat/go-router"
	"net/http"
)

type WSHandleFunc interface {
	Upgrade(http.ResponseWriter, *http.Request, http.Header) (interface{}, error)
	ReadMessage() (int, []byte, error)
	WriteMessage(int, []byte) error
	Close() error
}

type WSHandle struct {
	Router HTTPRoute
	Func   WSHandleFunc
	Err    error
}

func NewWSHandle(hr HTTPRoute, f WSHandleFunc) (*WSHandle, error) {
	if hr == nil {
		return nil, ErrRouter
	} else if f == nil {
		return nil, ErrHandleFunc
	}
	return &WSHandle{Router: hr, Func: f}, nil
}

func (h *WSHandle) Error() error {
	return h.Err
}

func (h *WSHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.Func == nil {
		h.Err = ErrHandleFunc
		return
	} else if _, err := h.Func.Upgrade(w, r, nil); err != nil {
		h.Err = err
		return
	}
	defer h.Func.Close()
	for {
		if mt, message, err := h.Func.ReadMessage(); err != nil {
			h.Err = err
			break
		} else if req, err := h.HandleRequest(message, r); err != nil {
			h.Err = err
			continue
		} else if resp, err := h.HandleResponse(req, w); err != nil {
			h.Err = err
			continue
		} else if result, err := resp.Bytes(); err != nil {
			h.Err = err
			continue
		} else if err = h.Func.WriteMessage(mt, result); err != nil {
			h.Err = err
			break
		}
	}
}

func (h *WSHandle) HandleRequest(data []byte, r *http.Request) (*WSRequest, error) {
	req, err := BytesToWSRequest(data)
	if err != nil {
		return nil, err
	} else if !req.Valid() {
		return nil, ErrWSRequest
	}
	req.Request = r
	return req, nil
}

func (h *WSHandle) HandleResponse(req *WSRequest, w http.ResponseWriter) (*WSResponse, error) {
	resp, err := NewWSResponse(req, w)
	if err != nil {
		return nil, err
	}
	// ?????????
	return resp, nil
}
