package mysql

import "Go-Chat/internal/model"

func GetUserFriend(uf *model.UserFriend){
	db.First(&uf,"user_id = ? and friend_id = ?",uf.ID,uf.FriendId)
}

func CreateUserFriend(uf *model.UserFriend){
	db.Create( uf )
}
