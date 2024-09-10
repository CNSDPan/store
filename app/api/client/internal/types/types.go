// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Mobile   int64  `json:"mobile,omitempty,string"`
	Password string `json:"password,omitempty"`
}

type LoginRes struct {
	UserId int64  `json:"userId,string"`
	Token  string `json:"token"`
}

type RegisterReq struct {
	Mobile   int64  `json:"mobile,omitempty,string"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

type RegisterRes struct {
	UserId int64  `json:"userId,string"`
	Token  string `json:"token"`
}

type Response struct {
	ErrMsg  string `json:"errMsg`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type StoreInfoReq struct {
	StoreId int64 `json:"storeId,omitempty,string"`
}

type StoreInfoRes struct {
	StoreId  int64  `json:"storeId,string"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Contacts int64  `json:"contacts"`
}

type StoreListRes struct {
	Total   int64       `json:"total"`
	Page    int64       `json:"page"`
	Limit   int64       `json:"limit"`
	Offset  int64       `json:"offset"`
	Current int64       `json:"current"`
	Rows    interface{} `json:"rows"`
}

type StoreUsersReq struct {
	StoreId int64 `json:"storeId,omitempty,string"`
	Limit   int64 `json:"limit"`
	Offset  int64 `json:"offset"`
}

type StoreUsersRes struct {
	Total   int64       `json:"total"`
	Page    int64       `json:"page"`
	Limit   int64       `json:"limit"`
	Offset  int64       `json:"offset"`
	Current int64       `json:"current"`
	Rows    interface{} `json:"rows"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
}

type UserInfoReq struct {
	UserId int64  `json:"userId,omitempty,string"`
	Token  string `header:"authorization,"`
}

type UserInfoRes struct {
	UserId int64  `json:"userId,string"`
	Mobile int64  `json:"mobile,omitempty,string"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
