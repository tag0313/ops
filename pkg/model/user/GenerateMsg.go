package user

import (
	"ops/pkg/model"
)

type GenerateMsgRequest struct {
	PbkAddr string `json:"pbkaddr" binding:"required" example:"0xDD685E8Ec061A3827191c0bFdE431ECa5Fc44ba5"`
}

type GenerateMsgResponse struct {
	model.JSONResult
	Data Message `json:"data"`
}

type Message struct {
	Message string `json:"message" example:"vk73Gqj6mlGGBybBEKQ3FQ3nffECB2fK265ZCmWIuVaSiNhIuJEG9jGxCxTwsK6tbWDNJB9NWenlgIKlilgOwpwyHiujUghcF8Q9t4sWU3xPL49YTbaaZw7zDKRQpZdE"`
}

func (r *GenerateMsgResponse) NewSuccess(result string) *GenerateMsgResponse {
	r.JSONResult.NewSuccess()
	r.Data.Message = result
	return r
}
