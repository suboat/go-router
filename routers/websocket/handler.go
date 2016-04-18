package router

import (
	. "github.com/suboat/go-router"
	//"net/http"
)

type WSHandler struct {
	Router HTTPRoute
}

//func NewWSHandler(r HTTPRoute) *WSRouter {
//	ws := newWSRouter(r)
//	if r == nil {
//		ws.err = ErrRouter
//	}
//	return ws
//}
//
//func (h *WSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	c, err := upgrader.Upgrade(w, r, nil)
//	if err != nil {
//		logS("upgrade:", err)
//		exit <- err
//		return
//	}
//	defer c.Close()
//	for {
//		mt, message, err := c.ReadMessage()
//		if err != nil {
//			logS("read:", err)
//			exit <- err
//			break
//		}
//		logS(fmt.Sprintf("recv: %s", message))
//		err = c.WriteMessage(mt, message)
//		if err != nil {
//			logS("write:", err)
//			exit <- err
//			break
//		}
//	}
//}
