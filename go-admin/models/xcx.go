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
	LineId int `json:"line_id"`
	SiteId int `json:"site_id"`
	QqLongitude string `json:"qq_longitude"`
	QqLatitude string `json:"qq_latitude"`
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
	LineId int `json:"line_id"`
}
type LineInfoReq struct {
	LineId int `json:"line_id"`
}
type LineStudentsReq struct{
	LineId int `json:"line_id"`
}
type LineCheckReq struct{
	LineId int `json:"line_id"`
}
type LineCheckRsp struct{
	StartAt int `json:"start_at"`
	Line ScbLines `json:"line"`
}
type FollowRecordReq struct {
	PageSize int `json:"page_size"`
	PageIndex int `json:"page_index"`
}