package tools

type Keys struct {

}
func (k Keys)ApiToken(token string) string{
	return "apiToken:" + token
}
