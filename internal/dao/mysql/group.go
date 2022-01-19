package mysql

import "Go-Chat/internal/model"

func GetGroupByName(name string) *model.Group {
	var queryGroup *model.Group
	db.Model(&model.Group{}).Where("name = ?",name ).First(&queryGroup)
	return queryGroup
}

func CreateGroup(group *model.Group){
	db.Create( group )
}

func GetGroupByUuid(uuid string) *model.Group {
	group := new( model.Group )
	db.First( group,"uuid = ?",uuid )
	return group
}

func GetGroupResponseByUseruuid(userId int32) ([]model.GroupResponse,error ){
	var groups []model.GroupResponse
	db.Raw("SELECT g.id AS group_id, g.uuid, g.created_at, g.name, g.notice FROM group_members AS gm LEFT JOIN `groups` AS g ON gm.group_id = g.id WHERE gm.user_id = ?",
		userId).Scan(&groups)

	return groups, nil
}


func GetUserListByGroupID(groupID int32) []model.User{
	var users []model.User
	db.Raw("SELECT u.uuid, u.avatar, u.username FROM `groups` AS g JOIN group_members AS gm ON gm.group_id = g.id JOIN users AS u ON u.id = gm.user_id WHERE g.id = ?",groupID).Scan(&users)
	return users
}