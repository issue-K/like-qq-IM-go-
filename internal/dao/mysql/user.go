package mysql

import "Go-Chat/internal/model"

// 创建一个用户
func CreateUser( user *model.User) error {
	return db.Create( user ).Error
}

func GetUserCountByName( username string ) int64 {
	var count int64
	count = 0
	db.Model(&model.User{}).Where("username",username).Count(&count)
	return count
}

func GetUserByName( username string ) *model.User{
	var queryUser *model.User
	db.Model(&model.User{}).Where("username = ?",username ).First(&queryUser)
	return queryUser
}

func GetUserByUuid( user *model.User ){
	db.First(&user,"uuid = ?",user.Uuid )
}

func UpdUserAvatar(user *model.User){
	db.Updates(user)
}

func GetUserIdListByUuid(uuid string) []int{
	var IdList []int
	db.Raw("select friend_id from user_friends where user_id = ?",uuid).Scan( &IdList )
	return IdList
}

func GetUserListByIdList(IdList []int)( UserList [] model.User){
	for _,id := range IdList{
		var user model.User
		db.Model(&model.User{}).Where("id = ?",id ).First( &user )
		if user.Id != 0{
			UserList = append( UserList,user )
		}
	}
	return
}