package models

const UserId string = "userId"
type XcxLoginReq struct {
	Phone string `json:"phone"`
	Password string `json:"password"`
}
type XcxLoginRsp struct {
	Token string `json:"token"`
}