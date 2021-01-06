package models
type TicketTokenReq struct {
	Ticket string `json:"ticket"`
}

type TicketTokenRsp struct {
	Token string `json:"token"`
}

type BindReq struct {
	Mobile string `json:"mobile"`
}