package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var node1Conn *websocket.Conn // WebSocket连接到Node1

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			return
		}
	}
}

func handleSendRequest(w http.ResponseWriter, r *http.Request) {
	// 处理发送消息的逻辑，您可以在此处使用WebSocket向客户端发送消息
	fmt.Println("Received a POST request at /send")
	// 读取请求的JSON数据并处理
	// ...

	// 向Node1发送数据
	if node1Conn != nil {
		node1Conn.WriteMessage(websocket.TextMessage, []byte("Hello, this is Node2!"))
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/ws", handleConnection)
	http.HandleFunc("/send", handleSendRequest) // 添加处理POST请求的路由

	// 创建WebSocket连接到Node1
	node1Conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:5000/ws", nil)
	if err != nil {
		log.Println(err)
	}
	defer node1Conn.Close()

	fmt.Println("Node2 is running on :5001")
	http.ListenAndServe(":5001", nil)
}
