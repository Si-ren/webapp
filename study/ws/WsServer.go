package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"strconv"
	"time"
)

type WsServer struct {
	Listener net.Listener
	Addr     string
	Upgrade  *websocket.Upgrader
}

func NewWsServer(port int) *WsServer {
	ws := new(WsServer)
	ws.Addr = "0.0.0.0:" + strconv.Itoa(port)
	ws.Upgrade = &websocket.Upgrader{
		HandshakeTimeout: 5 * time.Second,
		ReadBufferSize:   4096,
		WriteBufferSize:  4096,
	}
	return ws
}
func (s *WsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := s.Upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("升级失败 ：%v\n", err)
		return
	}
	fmt.Printf("已经与 %s 建立了链接\n", r.RemoteAddr)
	go s.HandleOneConnection(conn)

}

func (s *WsServer) HandleOneConnection(conn *websocket.Conn) {
	defer conn.Close()
	for {

		MessageType, request, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("ws read err : ", err)
			return
		}

		fmt.Printf("MessageType %d , message %s\n", MessageType, request)
		err = conn.WriteMessage(websocket.TextMessage, []byte("This is response message."))
		if err != nil {
			fmt.Println("ws write err : ", err)
			return
		}
	}
}

func (s *WsServer) Start() error {
	var err error
	s.Listener, err = net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	if err != nil {
		fmt.Println("ws listen err : ", err)
		return err
	}
	if err = http.Serve(s.Listener, s); err != nil {
		fmt.Println("ws server start err :", err)
		return err
	}
	return nil
}

func main() {
	ws := NewWsServer(8191)
	ws.Start()
}
