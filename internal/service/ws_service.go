package service

import (
	"Go-Chat/internal/dao/mysql"
	"Go-Chat/internal/model"
	"Go-Chat/internal/protos"
	"errors"
	"log"
)

const (
	HEAT_BEAT = "heatbeat"
	PONG      = "pong"

	// 消息类型，单聊或者群聊
	MESSAGE_TYPE_USER  = 1
	MESSAGE_TYPE_GROUP = 2

	// 消息内容类型
	TEXT         = 1
	FILE         = 2
	IMAGE        = 3
	AUDIO        = 4
	VIDEO        = 5
	AUDIO_ONLINE = 6
	VIDEO_ONLINE = 7

	// 消息队列类型
	GO_CHANNEL = "gochannel"
	KAFKA      = "kafka"
)


type wsService struct{

}
var WsService = new(wsService)

func(w *wsService) GetMessages(message model.MessageRequest)( []model.MessageResponse,error){
	if message.MessageType == MESSAGE_TYPE_USER{
		queryUser := new(model.User)
		queryUser.Uuid = message.Uuid
		mysql.GetUserByUuid(queryUser)
		if queryUser.Id == 0{
			return nil,errors.New("用户不存在")
		}
		var friend *model.User
		friend = mysql.GetUserByName(message.FriendUsername)
		if friend.Id == 0{
			return nil,errors.New("用户不存在")
		}

		return mysql.GetMessageResponseList( queryUser,friend ),nil

	}else if message.MessageType == MESSAGE_TYPE_GROUP{
		return mysql.GetGroupMessage( message.Uuid )
	}
	return nil,errors.New("不支持的查询类型")
}

func (w *wsService)SaveMessage( message protos.Message ){
	user := new( model.User )
	user.Uuid = message.From
	mysql.GetUserByUuid( user )

	if user.Id == 0{
		log.Println("SaveMessage not find from user")
		return
	}
	var toUserId int32 = 0
	if message.MessageType == MESSAGE_TYPE_USER{
		toUser := new( model.User )
		toUser.Uuid = message.To
		mysql.GetUserByUuid( toUser )

		if toUser.Id == 0{
			log.Println("SaveMessage not find from toUser")
			return
		}
		toUserId = toUser.Id
	}
	if message.MessageType == MESSAGE_TYPE_GROUP{
		group := mysql.GetGroupByUuid( message.To )
		if group.ID == 0{
			return
		}
		toUserId = group.ID
	}
	saveMessage := &model.Message{
		FromUserId: user.Id,
		ToUserId: toUserId,
		Content: message.Content,
		ContentType: int16(message.ContentType),
		MessageType: int16(message.MessageType),
		Url:         message.Url,
	}
	mysql.CreateMessage( saveMessage )
}