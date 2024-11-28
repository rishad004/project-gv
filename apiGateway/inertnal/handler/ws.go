package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/rishad004/project-gv/apiGateway/inertnal/domain"
)

type Socket struct {
	Conns map[string][]*websocket.Conn
}

func (h *ApiHanlder) UserChat(w http.ResponseWriter, r *http.Request) {
	channel := r.URL.Query().Get("channel")
	if channel == "" {
		http.Error(w, "Channel required", http.StatusBadRequest)
		return
	}

	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	h.WsConn.Conns[channel] = append(h.WsConn.Conns[channel], conn)
	index := len(h.WsConn.Conns[channel]) - 1

	defer func() {
		h.WsConn.Conns[channel] = append(h.WsConn.Conns[channel][:index], h.WsConn.Conns[channel][index+1:]...)
		conn.Close()
	}()

	for {
		var chat domain.Chat

		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}

		if err = json.Unmarshal(msg, &chat); err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			break
		}

		log.Println(chat)

		if chat.Type == "superchat" {
			amount, _ := strconv.Atoi(chat.Amount)
			if err = h.Superchat(amount, int(r.Context().Value("Id").(uint))); err != nil {
				log.Println(err)
				continue
			}
		}

		h.WsConn.sendToEveryViewer(channel, msg)
	}
}

func (s *Socket) sendToEveryViewer(channel string, message []byte) {

	for _, v := range s.Conns[channel] {
		v.WriteMessage(websocket.TextMessage, message)
	}

}
