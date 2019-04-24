package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"log"
	"math/rand"
	"html/template"
	"sync"
)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type Clients struct {
	m sync.RWMutex
	cs map[*websocket.Conn]bool
}

var (
	clients   = NewClients()
	msgPool = make(chan Message)
	upgrader  = websocket.Upgrader{}
)

func main() {
	http.HandleFunc("/", inRoom)
	http.HandleFunc("/ws", iniWsConn)
	// 静态文件服务，例如js css image
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public", fs))
	// 独立goroutine分发消息
	go handleMsgs()
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 进入聊天室网页
func inRoom(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, nil)
}

// 初始化连接，从HTTP Upgrade到Websocket
func iniWsConn(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer wsConn.Close()

	clients.addClient(wsConn)

	// 为客户端随机分配一个用户名
	m := Message{pickUser(), ""}
	wsConn.WriteJSON(m)

	for {
		var msg Message
		// 从客户端读取消息
		err := wsConn.ReadJSON(&msg)
		if err != nil {
			log.Printf("error : %v", err)
			clients.removeClient(wsConn)
			break
		}
		if len(msg.Message) == 0 {
			continue
		}
		msgPool <- msg
	}
}

func handleMsgs() {
	for {
		msg := <- msgPool
		clients.m.RLock()
		for client := range clients.cs {
			// 向客户端发消息
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error : %v", err)
				client.Close()
				clients.m.RUnlock()
				clients.removeClient(client)
			}
		}
		clients.m.RUnlock()
	}
}

func pickUser() string {
	users := []string{"John", "Root", "Michael", "Jim", "William", "Jack", "Ala",
	"Sofie", "Hawk", "Lily", "Lucy", "Alisa", "Captain"}
	l := len(users)
	n := rand.Intn(l)
	return users[n]
}

func (c *Clients) addClient(conn *websocket.Conn) {
	c.m.Lock()
	c.cs[conn] = true
	c.m.Unlock()
}

func (c *Clients) removeClient(conn *websocket.Conn) {
	c.m.Lock()
	delete(c.cs, conn)
	c.m.Unlock()
}

func NewClients() *Clients {
	c := &Clients{}
	c.cs = make(map[*websocket.Conn]bool)
	return c
}