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