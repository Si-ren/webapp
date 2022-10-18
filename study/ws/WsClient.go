package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"time"
)

func main() {
	dialer := &websocket.Dialer{
		NetDial:           nil,
		NetDialContext:    nil,
		NetDialTLSContext: nil,
		Proxy:             nil,
		TLSClientConfig:   nil,
		HandshakeTimeout:  0,
		ReadBufferSize:    0,
		WriteBufferSize:   0,
		WriteBufferPool:   nil,
		Subprotocols:      nil,
		EnableCompression: false,
		Jar:               nil,
	}
	conn, resp, err := dialer.Dial("ws://localhost:8191", nil)
	defer conn.Close()
	if err != nil {
		fmt.Println("ws dial err: ", err)
		fmt.Println(resp.StatusCode)
		msg, _ := io.ReadAll(resp.Body)
		fmt.Println(string(msg))
		return
	}
	for k, v := range resp.Header {
		fmt.Printf("header %s : %s\n", k, v)
	}
	conn.WriteMessage(websocket.TextMessage, []byte("zxcvbn"))
	time.Sleep(1 * time.Second)
}
