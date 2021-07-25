package report

import (
	"ops/pkg/model"
	reportService "ops/proto/report"
)

type ReportBase struct {
	Reporting    string `json:"reporting" binding:"required"`
	ReportReason string `json:"report_reason" binding:"required"`
}

type ReportOop struct {
	Oid string `json:"oid" binding:"required"`
	Uid string `json:"uid" binding:"required" swaggerignore:"true"`
	ReportBase
}

func (r ReportOop) SetValue() *reportService.ReportOopInfo {
	return &reportService.ReportOopInfo{
		Oid:          r.Oid,
		Uid:          r.Uid,
		Reporting:    r.Reporting,
		ReportReason: r.ReportReason,
	}
}

type ReportUser struct {
	Uid string `json:"uid" swaggerignore:"true"`
	ReportBase
}

func (r ReportUser) SetValue() *reportService.ReportUserInfo {
	return &reportService.ReportUserInfo{
		Uid:          r.Uid,
		Reporting:    r.Reporting,
		ReportReason: r.ReportReason,
	}
}

type UserID struct {
	Uid string `json:"uid" binding:"required"`
}

func (u UserID) SetValue() *reportService.UserID {
	return &reportService.UserID{
		Uid: u.Uid,
	}
}

type ReportResp struct {
	model.JSONResult
}

func (u *ReportResp) NewSuccess(value *reportService.ReportResult) *ReportResp {
	u.JSONResult.NewSuccess()
	return u
}

type ReportNumResp struct {
	model.JSONResult
	Times int64 `json:"times"`
}

func (u *ReportNumResp) NewSuccess(value *reportService.UserReportedTimes) *ReportNumResp {
	u.JSONResult.NewSuccess()
	u.Times = value.Times
	return u
}
