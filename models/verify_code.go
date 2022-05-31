package model

type VerifyCode struct {
	Id         int64  `json:"id"`
	Account    string `json:"account"`
	Code       string `json:"code"`
	Type       int    `json:"type"`
	Status     int64  `json:"status"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	ExpireTime int64  `json:"expire_time"`
}

const VerifyCodeStatusActive = 0
const VerifyCodeStatusUsed = 1