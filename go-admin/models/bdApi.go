package models
type BdApiGetTokenRsp struct {
	RefreshToken string
	ExpiresIn int
	Scope string
	SessionKey string
	AccessToken string
	SessionSecret string

}
type BdApiFacesetAddRsp struct {
	FaceToken string
	Location BdApiFacesetAddRspLocation

}
type BdApiFacesetAddRspLocation struct {
	LogId string
	Left float64
	Top float64
	Width float64
	Height float64
	Rotation int64
}