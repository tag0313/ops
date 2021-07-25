package user

import (
	"ops/pkg/model"
)

type LoginRequest struct {
	Pbk  string `json:"pbk" binding:"required"`
	Sign string `json:"sign" binding:"required"`
}

type LoginResponse struct {
	model.JSONResult
	Data Token `json:"data,omitempty"`
}

type Token struct {
	Token string `json:"token" example:"vk73Gqj6mlGGBybBEKQ3FQ3nffECB2fK265ZCmWIuVaSiNhIuJEG9jGxCxTwsK6tbWDNJB9NWenlgIKlilgOwpwyHiujUghcF8Q9t4sWU3xPL49YTbaaZw7zDKRQpZdE"`
}

func (r *LoginResponse) NewError(recode string) *LoginResponse {
	r.JSONResult.NewError(recode)
	return r
}

func (r *LoginResponse) NewSuccess(result string) *LoginResponse {
	r.JSONResult.NewSuccess()
	r.Data.Token = result
	return r
}
