package heartbeat

import (
	"net/http"
	"log"
	"time"
	"encoding/json"
)

var startTime time.Time

type Message struct {
	UntilNow string `json:"until_now"`
}

func init() {
	startTime = time.Now()
}

func Heartbeat(addr string) error {
	h := &Handler{}
	err := http.ListenAndServe(addr, h)
	if err != nil {
		return err
	}
	log.Println("start listening http on: ", addr)
	return nil
}

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := time.Since(startTime).String()
	m := NewMessage(s)
	encoded, err := json.Marshal(m)
	if err != nil {
		log.Printf("fail to encode message, err: %s", err.Error())
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(encoded)
}

func NewMessage(u string) *Message {
	return &Message{
		UntilNow: u,
	}
}
