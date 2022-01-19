package ws

import (
	"Go-Chat/internal/protos"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"log"
)

type Client struct{
	Conn *websocket.Conn
	Name string
	Send chan []byte
}

func (c *Client) Read(){
	defer func(){
		MainHub.UnRegister <- c
		c.Conn.Close()
	}()

	for{
		c.Conn.PongHandler()
		_, message ,err := c.Conn.ReadMessage()
		if err != nil{
			log.Println("client read message error",err.Error() )
			break
		}
		msg := &protos.Message{}
		proto.Unmarshal( message,msg )
		if msg.Type == HEAT_BEAT{
			pong := &protos.Message{
				Content: PONG,
				Type: HEAT_BEAT,
			}
			pongByte, err := proto.Marshal( pong )
			if err != nil{
				log.Println("client marshal message error:",err.Error() )
			}
			c.Conn.WriteMessage(websocket.BinaryMessage,pongByte)
		}else{
			MainHub.Broadcast <- message
		}
	}
}

func (c *Client) Write(){
	defer func(){
		c.Conn.Close()
	}()
	for message := range c.Send{
		c.Conn.WriteMessage(websocket.BinaryMessage,message)
	}
}