package mysql

import "Go-Chat/internal/model"

func ExistGroupMember(userID int32,groupID int32) bool{
	var count int64
	count = 0
	db.Model( &model.GroupMember{}).Where("user_id = ? and group_id = ?",userID,groupID).Count(&count)
	return count!=0
}

func CreateGroupMember(member *model.GroupMember){
	db.Create( member )
}
