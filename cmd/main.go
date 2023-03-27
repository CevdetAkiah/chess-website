package main

import (
	"context"
	"encoding/json"
	"fmt"
	postgres "go-projects/chess/database/postgres"
	"go-projects/chess/service"
	"go-projects/chess/util"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

// websocket server
type WsServer struct {
	lock     sync.Mutex
	conns    map[*websocket.Conn]bool
	DBAccess service.DbService
}

func NewWebsocket(DBA service.DbService) *WsServer {
	return &WsServer{
		conns:    make(map[*websocket.Conn]bool),
		DBAccess: DBA,
	}
}

func (wss *WsServer) handleWS(wsc *websocket.Conn) {
	wss.lock.Lock()
	wss.conns[wsc] = true
	wss.lock.Unlock()

	wss.readConn(wsc)
}

type UserMsg struct {
	Username string `json:"name"`
	Message  string `json:"message"`
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

func (wss *WsServer) readConn(wsc *websocket.Conn) {
	var outgoingMessage []byte
	var username []byte
	buf := make([]byte, 1024) // TODO: optimize this

	if util.CheckLogin(wsc.Request(), wss.DBAccess) {
		username = wss.getUserName(wsc.Request())
	}

	for {
		n, err := wsc.Read(buf) // read frame from conn and put data into the buffer
		if err != nil {
			if err == io.EOF { // break connection if user closes connection
				break
			}
			fmt.Println("read error:", err) // TODO: handle this error better
		}
		outgoingMessage = encodeUserMsg(username, buf[:n])

		wss.broadcast(outgoingMessage)
	}

}

// getUserName returns the username for the message giver.
// TODO: refactor this into a general function. Doesn't need to be limited to websocket server
func (wss *WsServer) getUserName(r *http.Request) []byte {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err) // TODO: handle this error better
	}
	session, err := wss.DBAccess.SessionService.SessionByUuid(cookie.Value)
	if err != nil {
		fmt.Println(err) // TODO: handle this error better
	}
	user, err := wss.DBAccess.UserService.UserByEmail(session.Email)
	if err != nil {
		fmt.Println(err) // TODO: handle this error better
	}
	username := user.Name + ": "
	return []byte(username)
}

func (wss *WsServer) broadcast(msg []byte) {
	for ws := range wss.conns {
		// send message to each active connection
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(msg); err != nil {
				fmt.Println("Broadcast error: ", err) // TODO: handle this error better
			}
		}(ws)
	}

}

func main() {
	// test database connection
	err := postgres.Db.Ping()
	if err != nil {
		err = fmt.Errorf("cannot connect to database with error: %w", err)
		log.Fatalln(err)
	}
	fmt.Println("connected to database chess")

	// set up the database service/access
	DBAccess := service.NewDbService(
		postgres.Db,
		postgres.UserAccess{},
		postgres.SessionAccess{},
		log.New(os.Stdout, "database-api ", log.LstdFlags),
	)
	wsServer := NewWebsocket(DBAccess)
	mux := NewMux(DBAccess, wsServer)

	// set up server
	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	DBAccess.Println("Connected to port :8080")
	go func() { // go routine so the enclosed doesn't block
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// sigChan signals when the interrupt or kill signal is received from the OS.
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	DBAccess.Printf("Received terminate message ", sig)

	// Graceful shutdown. Users are given 2 minutes to finish their game if the server needs to restart for any reason
	t := time.Now().Add(time.Second * 120)
	tc, _ := context.WithDeadline(context.Background(), t)
	server.Shutdown(tc)
}
