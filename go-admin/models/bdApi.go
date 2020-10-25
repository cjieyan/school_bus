package models
type BdApiGetTokenRsp struct {
	RefreshToken string `json:"refresh_token"`
	ExpiresIn int `json:"expires_in"`
	Scope string `json:"scope"`
	SessionKey string `json:"session_key"`
	AccessToken string `json:"access_token"`
	SessionSecret string `json:"session_secret"`

}
type BdApiFacesetAddReq struct {
	Image string `json:"image"`
	ImageType string `json:"image_type"`
	GroupId string `json:"group_id"`
	UserId string `json:"user_id"`
	ActionType string `json:"action_type"`
}
type BdApiFacesetMutilSearchReq struct {
	Image string `json:"image"`
	ImageType string `json:"image_type"`
	GroupIdList string `json:"group_id_list"`
	MaxFaceNum int `json:"max_face_num"`
	MatchThreshold int `json:"match_threshold"`
	QualityControl string `json:"quality_control"`
	LivenessControl string `json:"liveness_control"`
	MaxUserNum uint32 `json:"max_user_num"`
}

type BdApiFacesetMutilSearchRsp struct {
	Cached    int64    `json:"cached"`
	ErrorCode int64    `json:"error_code"`
	ErrorMsg  string   `json:"error_msg"`
	LogID     int64    `json:"log_id"`
	Result    BdApiFacesetMutilSearchRspResult `json:"result"`
	Timestamp int64    `json:"timestamp"`
}

type BdApiFacesetMutilSearchRspResult struct {
	FaceList []BdApiFacesetMutilSearchRspResultFaceList `json:"face_list"`
	FaceNum  int64      `json:"face_num"`
}

type BdApiFacesetMutilSearchRspResultFaceList struct {
	FaceToken string     `json:"face_token"`
	Location  BdApiFacesetMutilSearchRspResultFaceListLocation   `json:"location"`
	UserList  []BdApiFacesetMutilSearchRspResultFaceListUserList `json:"user_list"`
}

type BdApiFacesetMutilSearchRspResultFaceListUserList struct {
	GroupID  string  `json:"group_id"`
	Score    float64 `json:"score"`
	UserID   string  `json:"user_id"`
	UserInfo string  `json:"user_info"`
}

type BdApiFacesetMutilSearchRspResultFaceListLocation struct {
	Height   int64   `json:"height"`
	Left     float64 `json:"left"`
	Rotation int64   `json:"rotation"`
	Top      float64 `json:"top"`
	Width    int64   `json:"width"`
}

type BdApiFacesetAddRsp struct {
	ErrorCode int `json:"error_code"`
	ErrorMsg string `json:"error_msg"`
	LogId int64 `json:"log_id"`
	Timestamp int `json:"timestamp"`
	Cached int `json:"cached"`
	Result BdApiFacesetAddRspResult
}
type BdApiFacesetAddRspResult struct {
	FaceToken string `json:"face_token"`
	Location BdApiFacesetAddRspResultLocation `json:"location"`
}
type BdApiFacesetAddRspResultLocation struct {
	LogId string
	Left float64
	Top float64
	Width float64
	Height float64
	Rotation int64
}