package models

const UserId string = "userId"
type XcxLoginReq struct {
	Phone string `json:"phone"`
	Password string `json:"password"`
}
type XcxLoginRsp struct {
	Token string `json:"token"`
}
type XcxGetOn struct {
	StudentId int `json:"student_id"`
}
type XcxGetOff struct {
	StudentId int `json:"student_id"`
}