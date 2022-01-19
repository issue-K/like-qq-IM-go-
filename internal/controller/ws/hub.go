package ws

import (
	"Go-Chat/config"
	"Go-Chat/internal/dao/mysql"
	"Go-Chat/internal/model"
	"Go-Chat/internal/protos"
	"Go-Chat/internal/service"
	util "Go-Chat/pkg/utils"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"strings"
)

var MainHub = NewHub()

type Hub struct{
	Clients map[string]*Client
	Broadcast chan []byte
	Register chan *Client
	UnRegister chan *Client
}

func NewHub() *Hub{
	return &Hub{
		make( map[string]*Client ),
		make( chan []byte ),
		make(chan *Client),
		make(chan *Client),
	}
}

func receive(data []byte){
	MainHub.Broadcast <- data
}
func (h *Hub) Start(){
	log.Println("start Hub.......")
	for{
		select{
		case client := <-h.Register:
			h.Clients[client.Name] = client
		case client := <-h.UnRegister:
			if _,ok := h.Clients[client.Name]; ok{
				close( client.Send )
				delete( h.Clients,client.Name )
			}
		case message := <-h.Broadcast:
			msg := &protos.Message{}
			proto.Unmarshal( message,msg )
			if msg.To != ""{
				if msg.ContentType >= TEXT && msg.ContentType<=VIDEO{
					if _, exits := h.Clients[msg.From]; exits{
						saveMessage(msg)
					}
					if msg.MessageType == MESSAGE_TYPE_USER{
						client, ok := h.Clients[msg.To]
						if ok {
							msgByte, err := proto.Marshal(msg)
							if err == nil {
								client.Send <- msgByte
							}
						}
					}else if msg.MessageType == MESSAGE_TYPE_GROUP{
						send_GroupMessage( msg,h )
					}
				}
			}
		}
	}
}

func send_GroupMessage(msg *protos.Message,h *Hub){
	//需要发送给每一个群成员
	fromUser := new( model.User )
	fromUser.Uuid = msg.From
	mysql.GetUserByUuid( fromUser )

	users := service.GroupService.GetUserIdByGroupUuid( msg.To )
	for _,user := range users{
		if msg.From == user.Uuid{
			continue
		}
		log.Println("发送给用户 ",user.Username )
		client ,ok:= h.Clients[user.Uuid]
		if !ok{
			continue
		}
		log.Println("没问题")

		msgSend := &protos.Message{
			Avatar: fromUser.Avatar,
			FromUsername: fromUser.Username,
			From: msg.To,
			To: msg.From,
			Content: msg.Content,
			ContentType: msg.ContentType,
			Type: msg.Type,
			MessageType:  msg.MessageType,
			Url: msg.Url,
		}
		msgByte, err := proto.Marshal(msgSend)
		if err==nil {
			log.Println("发出去了")
			client.Send <- msgByte
		}
	}
}


// 保存消息，如果是文本消息直接保存，如果是文件，语音等消息，保存文件后，保存对应的文件路径
func saveMessage(message *protos.Message){
	// 如果上传的是base64字符串文件，解析文件保存
	log.Println("message type: ",message.ContentType )
	if message.ContentType == 3 {  //save the image
		// 普通的文件二进制上传
		//fileSuffix := util.GetFileType(message.File)
		//nullStr := ""
		//if nullStr == fileSuffix {
		//	log.Println("fileSuffix: ",message.FileSuffix )

		fileSuffix := strings.ToLower(message.FileSuffix)
		log.Println("message FileSuffix: ",message.FileSuffix)

		contentType := util.GetContentTypeBySuffix(fileSuffix)
		url := uuid.New().String() + "." + fileSuffix
		err := ioutil.WriteFile(config.AppCF.Static_url+url, message.File, 0666)
		if err != nil {
			log.Println("write file error", err.Error() )
			return
		}
		message.Url = url
		message.File = nil
		message.ContentType = contentType
	}
	service.WsService.SaveMessage( *message )
}