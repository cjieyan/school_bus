package models

const UserId string = "userId"
type XcxLoginReq struct {
	Phone string `json:"phone"`
	Password string `json:"password"`
}
type XcxLoginRsp struct {
	Token string `json:"token"`
}
type SwipeReq struct {
	StudentId int `json:"student_id"`
}
type SwipeRsp struct {
	Status int `json:"status"`
}
//刷脸时间
type SwipeAt struct {
	Status int `json:"status"`
	Time int 	`json:"time"`
}

type LineFinishReq struct {
	StudentId int `json:"student_id"`
}