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

var node2Conn *websocket.Conn // WebSocket连接到Node2

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

	// 向Node2发送数据
	if node2Conn != nil {
		node2Conn.WriteMessage(websocket.TextMessage, []byte("Hello, this is Node1!"))

		// 接收来自Node2的响应
		_, response, err := node2Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// 输出Node2的响应
		fmt.Println("Received response from Node2:", string(response))
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/ws", handleConnection)
	http.HandleFunc("/send", handleSendRequest) // 添加处理POST请求的路由

	// 创建WebSocket连接到Node2
	node2Conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:5001/ws", nil)
	if err != nil {
		log.Println(err)
	}
	defer node2Conn.Close()

	fmt.Println("Node1 is running on :5000")
	http.ListenAndServe(":5000", nil)
}
