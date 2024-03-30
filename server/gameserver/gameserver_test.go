package gameserver

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/anthdm/hollywood/remote"
	"github.com/gorilla/websocket"
)

type TestGameServer struct {
	ctx *actor.Context
}

type testMessage struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

const (
	endpoint = "ws://localhost:4000/ws"
)

func setUpServer() error {
	// set up server
	serverEngine, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {

		return err
	}
	serverRemote := remote.New("localhost:4000", remote.NewConfig())
	serverEngine.Spawn(NewGameServer, "server")
	// set up session

	go serverRemote.Start(serverEngine)
	// set up websocket request

	defer func() {
		serverRemote.Stop()
	}()
	return nil
}

func setUpWebsocket() (*websocket.Conn, error) {
	dialer := websocket.Dialer{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, _, err := dialer.Dial(endpoint, nil)
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	return conn, nil
}

// func TestGameServerReceiveJoin(t *testing.T) {
// 	err := setUpServer()
// 	if err != nil {
// 		t.Errorf("Server engine failed to start: %b", err)
// 	}
// 	var conn *websocket.Conn
// 	conn, err = setUpWebsocket()
// 	if err != nil {
// 		t.Fatalf("websocket dialing broke with error: %s ", err.Error())
// 	}

// 	jsonData := []byte(`{"name": "test_user", "uniqueID": "123"}`)
// 	request := testMessage{
// 		Type: "join",
// 		Data: json.RawMessage(jsonData),
// 	}

// 	conn.WriteJSON(request)

// }

func TestPlayerState(t *testing.T) {
	serverEngine, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {

		t.Fatal(err)
	}
	serverRemote := remote.New("localhost:4000", remote.NewConfig())
	serverEngine.Spawn(NewGameServer, "server")
	// set up session

	go serverRemote.Start(serverEngine)
	// set up websocket request

	defer func() {
		serverRemote.Stop()
	}()
	var conn *websocket.Conn
	conn, err = setUpWebsocket()
	if err != nil {
		t.Fatalf("websocket dialing broke with error: %s ", err.Error())
	}

	jsonData := []byte(`{"uniqueID": "123"}`)
	request := testMessage{
		Type: "join",
		Data: json.RawMessage(jsonData),
	}

	conn.WriteJSON(request)
	request = testMessage{
		Type: "playerState",
		Data: json.RawMessage(jsonData),
	}
	conn.WriteJSON(request)

	var conn2 *websocket.Conn
	jsonData2 := []byte(`{"uniqueID": "456"}`)

	request2 := testMessage{
		Type: "join",
		Data: json.RawMessage(jsonData2),
	}
	conn2.WriteJSON(request2)

	request2 = testMessage{
		Type: "playerState",
		Data: json.RawMessage(jsonData2),
	}
	conn2.WriteJSON(request2)

	time.Sleep(1 * time.Second)
}

// func TestPlayerSessionJoin(t *testing.T) {
// 	// set up server
// 	serverEngine, err := actor.NewEngine(actor.NewEngineConfig())
// 	if err != nil {
// 		t.Errorf("Server engine failed to start: %b", err)
// 	}
// 	serverRemote := remote.New("http://localhost:4000", remote.NewConfig())
// 	spid := serverEngine.Spawn(NewGameServer, "server")

// 	go serverRemote.Start(serverEngine)

// 	defer func() {
// 		serverRemote.Stop()
// 	}()
// 	dialer := websocket.Dialer{
// 		ReadBufferSize:  1024,
// 		WriteBufferSize: 1024,
// 	}

// 	conn, _, err := dialer.Dial(endpoint, nil)
// 	if err != nil {
// 		t.Fatalf("websocket dialing broke with error: %s ", err.Error())
// 	}
// 	playerSessionSPID := serverEngine.Spawn(newPlayerSession(spid, 123, conn), "playerSession")
// 	// playerSession.
// 	jsonData := []byte(`{"name": "test_user", "uniqueID": "123"}`)
// 	request := testMessage{
// 		Type: "join",
// 		Data: json.RawMessage(jsonData),
// 	}
// 	serverEngine.Send(spid, &request)
// 	fmt.Println(playerSessionSPID)
// 	time.Sleep((time.Second * 1))
// }
