package model

type Friendparam struct{
	Uuid           string `json:"uuid" binding:"required"`
	FriendUsername string `json:"friendUsername" binding:"required"`
}

type MessageRequest struct {
	MessageType    int32  `json:"messageType"`
	Uuid           string `json:"uuid"`
	FriendUsername string `json:"friendUsername"`
}
