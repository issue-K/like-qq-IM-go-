package service

import (
	"Go-Chat/internal/dao/mysql"
	"Go-Chat/internal/model"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
)

type userService struct{

}
var UserService = new( userService )

func (u *userService) Register(user *model.User) (err error) {
	if mysql.GetUserCountByName(user.Username)>0{
		return errors.New("用户名已存在")
	}

	user.Uuid = uuid.New().String()
	return mysql.CreateUser(user)
}

func (u *userService)Login(user *model.User)(bool){
	queryUser := mysql.GetUserByName(user.Username)

	if queryUser.Password == user.Password{
		*user = *queryUser
		return true
	}
	return false
}

func (u *userService)GetUserOrGroupByName(name string) model.SearchResponse{
	queryUser := mysql.GetUserByName(  name )

	queryGroup := mysql.GetGroupByName( name )
	return model.SearchResponse{
		*queryUser,
		*queryGroup,
	}
}

func (u *userService)ModifyUserAvatar(avatar string,uuid string) error{
	queryUser := new( model.User )
	queryUser.Uuid = uuid
	mysql.GetUserByUuid(queryUser)
	if queryUser.Id == 0{
		return errors.New("user not exist")
	}
	queryUser.Avatar = avatar
	mysql.UpdUserAvatar(queryUser)
	return nil
}

func (u *userService)GetUserDetails( uuid string ) model.User{
	queryUser := new( model.User )
	queryUser.Uuid = uuid
	mysql.GetUserByUuid(queryUser)
	return *queryUser
}

func (u *userService) AddFriend( friendParam *model.Friendparam) error{
	queryUser := new( model.User )
	queryUser.Uuid = friendParam.Uuid
	mysql.GetUserByUuid( queryUser )

	if queryUser.Id == 0{
		return errors.New("user not exist")
	}

	friend := mysql.GetUserByName( friendParam.FriendUsername )
	log.Println("friendusername = ",friendParam.FriendUsername )
	if friend.Id == 0{
		return errors.New("friend not exist")
	}

	userFriend := model.UserFriend{
		UserId:queryUser.Id,
		FriendId: friend.Id,
	}
	mysql.GetUserFriend(&userFriend)
	if userFriend.ID != 0{
		return errors.New("userFriend relation exist")
	}
	if friend.Uuid == queryUser.Uuid{
		return errors.New("不能添加自己为好友")
	}

	mysql.CreateUserFriend( &userFriend )
	userFriend1 := model.UserFriend{  //opposite
		UserId:friend.Id,
		FriendId: queryUser.Id,
	}
	mysql.GetUserFriend(&userFriend1)
	mysql.CreateUserFriend( &userFriend1 )
	return nil
}

func (u *userService)GetUserList(uuid string)( []model.User){
	var user model.User
	user.Uuid = uuid
	mysql.GetUserByUuid(&user)
	var IdList []int
	IdList = mysql.GetUserIdListByUuid( fmt.Sprint(user.Id) )
	var UserList []model.User
	UserList = mysql.GetUserListByIdList( IdList )
	return UserList
}