package service

import (
	"Go-Chat/internal/dao/mysql"
	"Go-Chat/internal/model"
	"errors"
	"github.com/google/uuid"
	"log"
)

var GroupService = new( groupService )

type groupService struct{

}

func (g *groupService)CreateGroup(userUuid string,groupName string) (err error) {
	user := new( model.User )
	user.Uuid = userUuid
	mysql.GetUserByUuid( user )
	if user.Id <= 0{
		log.Println("该用户不在数据库中")
		return
	}
	temp := mysql.GetGroupByName(groupName)
	if temp.ID != 0{  //保证之前们没有这个群
		return errors.New("已存在该群名称")
	}

	group := &model.Group{
		UserId: user.Id,
		Uuid: uuid.New().String(),
		Name:groupName,
	}
	mysql.CreateGroup( group )

	return g.JoinGroup( group.Uuid,user.Uuid )
}

func (g *groupService)JoinGroup(groupUuid string,userUuid string ) (error) {
	user := new( model.User )
	user.Uuid = userUuid
	mysql.GetUserByUuid(user)
	if user.Id<=0{
		return errors.New("用户不存在")
	}

	group := mysql.GetGroupByUuid(groupUuid)
	if group.ID <= 0{
		return errors.New("该群组不存在")
	}

	if mysql.ExistGroupMember(user.Id,group.ID){
		return errors.New("已经加入了该群聊")
	}
	groupMember := &model.GroupMember{
		UserId: user.Id,
		GroupId: group.ID,
		Nickname: user.Username,
		Mute: 0,
	}
	mysql.CreateGroupMember(groupMember)
	return nil
}

func (g *groupService)GetGroup(userUuid string) ([] model.GroupResponse,error){
	user := new(model.User)
	user.Uuid = userUuid
	mysql.GetUserByUuid( user )
	if user.Id<=0{
		return nil,errors.New("用户不存在")
	}

	return mysql.GetGroupResponseByUseruuid( user.Id )
}

func (g *groupService)GetUserIdByGroupUuid(groupUuid string) []model.User{
	group := mysql.GetGroupByUuid(groupUuid)

	if group.ID <= 0{
		return nil
	}
	return mysql.GetUserListByGroupID( group.ID )
}