package tools

import "strconv"

type Keys struct {

}
func (k Keys)ApiToken(token string) string{
	return "apiToken:" + token
}
func (k Keys)BdApiTokenRsp() string{
	return "bdApiTokenRsp"
}
//上车
func (k Keys)GetOn( ymd string, teacherId int) string{
	return "getOn:" + ymd + ":" + strconv.Itoa(teacherId)
}
//记录刷脸时间
func (k Keys)SwipeAt(ymd string, lineId int) string{
	return "swipeAt:" + ymd + ":" + strconv.Itoa(lineId)
}
//记录正在刷脸的时间
func (k Keys)SwipeNow() string{
	return "swipeNow"
}
//记录正在刷脸的时间
func (k Keys)SwipeNowRecord(teacherId int) string{
	return "swipeNowRecord:" + strconv.Itoa(teacherId)
}
//下车
func (k Keys)GetOff(ymd string, teacherId int) string{
	return "getOff:" + ymd + ":" + strconv.Itoa(teacherId)
}
