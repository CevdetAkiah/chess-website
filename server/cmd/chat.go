package main

import (
	"encoding/json"
	"fmt"
	"go-projects/chess/service"
	"sync"

	"golang.org/x/net/websocket"
)

// websocket server
type WsServer struct {
	lock     sync.Mutex
	conns    map[*websocket.Conn]bool
	DBAccess service.DatabaseAccess
}

type UserMsg struct {
	Username string `json:"name"`
	Message  string `json:"message"`
}

func NewWebsocket(DBA service.DatabaseAccess) *WsServer {
	return &WsServer{
		conns:    make(map[*websocket.Conn]bool),
		DBAccess: DBA,
	}
}

func (wss *WsServer) handleWS(wsc *websocket.Conn) {
	wss.lock.Lock()
	wss.conns[wsc] = true
	wss.lock.Unlock()
	fmt.Println("connection opened: ", wsc.LocalAddr())

	wss.readConn(wsc)
	defer wsc.Close()

}

func (wss *WsServer) readConn(wsc *websocket.Conn) {
	var outgoingMessage []byte
	outgoingMessage = []byte(fmt.Sprint("Hi from GOLANG"))
	var msg []byte
	for {
		wsc.Read(msg)
		msgReceived := string(msg)
		fmt.Println(msgReceived)
		if msgReceived == "msg received" {
			return
		}
		wss.broadcast(outgoingMessage)
	}
	// var username []byte
	// buf := make([]byte, 1024) // TODO: optimize this

	// if util.CheckLogin(wsc.Request(), wss.DBAccess) {
	// username = util.GetUserName(wsc.Request(), wss.DBAccess)
	// }

	// n, err := wsc.Read(buf) // read frame from conn and put data into the buffer
	// if err != nil {
	// if err == io.EOF { // break connection if user closes connection
	// break
	// }
	// fmt.Println("read error:", err) // TODO: handle this error better
	// }
	// outgoingMessage = encodeUserMsg(username, buf[:n])

}

func (wss *WsServer) broadcast(msg []byte) {
	for ws := range wss.conns {
		// send message to each active connection
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(msg); err != nil {
				fmt.Println("Broadcast error: ", err) // TODO: handle this error better
			}
			delete(wss.conns, ws)
		}(ws)
	}
}

// encode this as JSON instead
func encodeUserMsg(uname, msg []byte) []byte {
	userMessage := &UserMsg{Username: string(uname), Message: string(msg)}
	outgoingMessage, err := json.Marshal(userMessage)
	if err != nil {
		fmt.Println(err) // TODO: handle this error better
	}
	return outgoingMessage
}
