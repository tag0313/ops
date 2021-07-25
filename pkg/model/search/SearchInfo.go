package search

import (
	"ops/pkg/model"
	searchService "ops/proto/search"
)

type SearchOption struct {
	ShowNumber int64 `json:"show_number" binding:"required"`
	PageNumber int64 `json:"page_number" binding:"required"`
}

type SearchContentByIDReq struct {
	Uid       string `json:"uid" binding:"required"`
	Timestamp string `json:"timestamp"`
	SearchOption
}

func (s SearchContentByIDReq) SetValue() *searchService.UserID {
	return &searchService.UserID{
		Uid:        s.Uid,
		Timestamp:  s.Timestamp,
		PageNumber: s.PageNumber,
		ShowNumber: s.ShowNumber,
	}
}

type SearchContentReq struct {
	Content   string `json:"content" binding:"required"`
	Timestamp string `json:"timestamp"  swaggerignore:"true"`
	SearchOption
}

func (s SearchContentReq) SetValue() *searchService.Content {
	return &searchService.Content{
		Content:    s.Content,
		Timestamp:  s.Timestamp,
		PageNumber: s.PageNumber,
		ShowNumber: s.ShowNumber,
	}
}

type SearchContentResp struct {
	model.JSONResult
	Data struct {
		Data        []*searchService.SearchContentResult `json:"data"`
		Timestamp   string                               `json:"timestamp"`
		CurrentPage int64                                `json:"current_page"`
		ShowNumber  int64                                `json:"show_number"`
	} `json:"data"`
}

func (u *SearchContentResp) NewSuccess(value *searchService.SearchContentResults) *SearchContentResp {
	u.JSONResult.NewSuccess()
	u.Data.Data = value.Data
	u.Data.Timestamp = value.Timestamp
	u.Data.CurrentPage = value.CurrentPage
	u.Data.ShowNumber = value.ShowNumber
	return u
}

type SearchUserResp struct {
	model.JSONResult
	Data struct {
		Data        []*searchService.SearchUserResult `json:"data"`
		CurrentPage int64                             `json:"current_page"`
		ShowNumber  int64                             `json:"show_number"`
	} `json:"data"`
}

func (u *SearchUserResp) NewSuccess(value *searchService.SearchUserResults) *SearchUserResp {
	u.JSONResult.NewSuccess()
	u.Data.Data = value.Data
	u.Data.CurrentPage = value.CurrentPage
	u.Data.ShowNumber = value.ShowNumber
	return u
}
