package models
type BdApiGetTokenRsp struct {
	RefreshToken string `json:"refresh_token"`
	ExpiresIn int `json:"expires_in"`
	Scope string `json:"scope"`
	SessionKey string `json:"session_key"`
	AccessToken string `json:"access_token"`
	SessionSecret string `json:"session_secret"`

}
type BdApiFacesetAddRsp struct {
	FaceToken string `json:"face_token"`
	Location BdApiFacesetAddRspLocation `json:"location"`

}
type BdApiFacesetAddRspLocation struct {
	LogId string
	Left float64
	Top float64
	Width float64
	Height float64
	Rotation int64
}