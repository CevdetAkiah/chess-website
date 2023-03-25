package main

import (
	"context"
	"fmt"
	postgres "go-projects/chess/database/postgres"
	"go-projects/chess/service"
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
	lock  sync.Mutex
	conns map[*websocket.Conn]bool
}

func NewWebsocket() *WsServer {
	return &WsServer{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (wss *WsServer) handleWS(wsc *websocket.Conn) {
	wss.lock.Lock()
	wss.conns[wsc] = true
	wss.lock.Unlock()

	wss.readConn(wsc)
}

func (wss *WsServer) readConn(wsc *websocket.Conn) {
	buf := make([]byte, 1024) // TODO: optimize this
	for {
		n, err := wsc.Read(buf) // read frame from conn and put data into the buffer
		if err != nil {
			if err == io.EOF { // break connection if user closes connection
				break
			}
			fmt.Println("read error:", err) // TODO: handle this error better
		}
		msg := buf[:n]

		wss.broadcast(msg)
	}

}

func (wss *WsServer) broadcast(msg []byte) {
	for ws := range wss.conns {
		// send message to each active connection
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(msg); err != nil {
				fmt.Println("Broadcast error: ", err) // TODO: ahandle this error better
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
	wsServer := NewWebsocket()
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
