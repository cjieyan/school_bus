package models

const UserId string = "userId"

type XcxLoginReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
type XcxLoginRsp struct {
	Token string `json:"token"`
}
type SwipeReq struct {
	StudentId   int    `json:"student_id"`
	LineId      int    `json:"line_id"`
	CarId      int    `json:"car_id"`
	SiteId      int    `json:"site_id"`
	QqLongitude string `json:"qq_longitude"`
	QqLatitude  string `json:"qq_latitude"`
}
type SwipeRsp struct {
	Status int `json:"status"`
}

//刷脸时间
type SwipeAt struct {
	Status int `json:"status"` //-1 未上车 0 上车 1 下车
	Time   int `json:"time"`
}
type LineStartReq struct {
	LineId int `json:"line_id"`
	CarId      int    `json:"car_id"`
}
type LineFinishReq struct {
	LineId int `json:"line_id"`
	CarId      int    `json:"car_id"`
	Force	int `json:"force"`
}
type LineFinishRsp struct {
	IsFinished int `json:"is_finished"`
	GetOn int `json:"get_on"`
}
type LineInfoReq struct {
	LineId int `json:"line_id"`
	CarId      int    `json:"car_id"`
}
type LineStudentsReq struct {
	LineId int `json:"line_id"`
	CarId      int    `json:"car_id"`
}
type LineCheckReq struct {
	LineId int `json:"line_id"`
	CarId      int    `json:"car_id"`
}
type LineCheckRsp struct {
	StartAt int      `json:"start_at"`
	Line    ScbLines `json:"line"`
}
type FollowRecordReq struct {
	PageSize  int `json:"page_size"`
	PageIndex int `json:"page_index"`
}

type FaceSwipeReq struct {
	Image       string `json:"image"`
	LineId      int    `json:"line_id"`
	CarId      int    `json:"car_id"`
	SiteId      int    `json:"site_id"`
	QqLongitude string `json:"qq_longitude"`
	QqLatitude  string `json:"qq_latitude"`
}
type FaceSwipeRsp struct {
	Num           int                         `json:"num"`
	StudentStatus []FaceSwipeRspStudentStatus `json:"students"`
	IsFinished    bool                        `json:"isFinished"`
}
type FaceSwipeRspStudentStatus struct {
	StudentId int `json:"student_id"`
	Status    int `json:"status"` //-1 未上车 0 上车 1 下车
}
type StudentInfoReq struct {
	LineId    int `json:"line_id"`
	StudentId int `json:"student_id"`
}
type StudentInfoRsp struct {
	Student ScbStudents `json:"student"`
}
