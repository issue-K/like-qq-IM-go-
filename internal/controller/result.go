package controller

var(
	ErrorServerBusy = "服务器繁忙"
	ErrorExistUsername = "用户名已存在"
)

type Result struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(data interface{} ) Result {
	return Result{
		Code: 0,
		Msg:"success",
		Data:data,
	}
}

func Failed(msg string ) Result {
	return Result{
		Code:-1,
		Msg:msg,
		Data:nil,
	}
}