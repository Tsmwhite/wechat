package socket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"time"
	"wechat/core/encrypt/token"
	model "wechat/models"
)

func sendMsg(room *model.Room, user *model.User) {
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8011", Path: "/ws"}
	log.Printf("connecting to %s", u.String())
	reqHeader := http.Header{}
	reqHeader = make(map[string][]string)
	reqHeader.Add("Sec-Websocket-Protocol", token.New(user.Uuid).ToString())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), reqHeader)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err := c.WriteJSON(map[string]interface{}{
				"type":        200,
				"second_type": 400,
				"recipient":   room.Uuid,
				"sender":      user.Uuid,
				"send_time":   time.Now().Unix(),
				"content":     "打卡",
				//"content":     "hi, i'm " + user.Name + ", now the time is " + time.Now().Format("2006/01/02 15:04:05"),
			})
			if err != nil {
				log.Println("WriteJSON Err: ", err)
			}
		}
	}
}

func Run() {
	var rooms []*model.Room
	model.DB().Table("rooms").Where("type = 1").Limit(100).Find(&rooms)
	for _, r := range rooms {
		members := r.GetMembers()
		//go sendMsg(r, model.GetUserByUuid(members[0]))
		for i, m := range members {
			if i == 1 {
				continue
			}
			go sendMsg(r, model.GetUserByUuid(m))
		}
	}
	for {
	}
}
