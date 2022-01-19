package mysql

import (
	"Go-Chat/internal/model"
	"errors"
)

func GetMessageResponseList(user *model.User,friend *model.User)( []model.MessageResponse ){
	var messages []model.MessageResponse
	db.Raw("SELECT m.id, m.from_user_id, m.to_user_id, m.content, m.content_type, m.url, m.created_at, u.username AS from_username, u.avatar, to_user.username AS to_username  FROM messages AS m LEFT JOIN users AS u ON m.from_user_id = u.id LEFT JOIN users AS to_user ON m.to_user_id = to_user.id WHERE from_user_id IN (?, ?) AND to_user_id IN (?, ?)",
		user.Id, friend.Id, user.Id, friend.Id).Scan(&messages)
	return messages
}

func CreateMessage( message *model.Message){
	db.Create( message )
}

func GetGroupMessage(uuid string)( []model.MessageResponse,error ){
	var group model.Group
	db.First( &group,"uuid = ?",uuid )
	if group.ID <= 0{
		return nil,errors.New("群组不存在")
	}

	var messages []model.MessageResponse
	db.Raw("SELECT m.id, m.from_user_id, m.to_user_id, m.content, m.content_type, m.url, m.created_at, u.username AS from_username, u.avatar FROM messages AS m LEFT JOIN users AS u ON m.from_user_id = u.id WHERE m.message_type = 2 AND m.to_user_id = ?",
		group.ID).Scan(&messages)

	return messages, nil
}