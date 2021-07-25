package handler

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"ops/pkg/logger"
	"ops/pkg/model/mgodb"
	utils2 "ops/pkg/utils"
	"ops/proto/report"
	"time"
)

type Report struct {
}

type ReportOop struct {
	Oid          string `json:"oid,omitempty" bson:"oid,omitempty"`
	Uid          string `json:"uid,omitempty" bson:"uid,omitempty"`
	Reporting    string `json:"reporting,omitempty" bson:"reporting,omitempty"`
	ReportReason string `json:"report_reason,omitempty" bson:"report_reason,omitempty"`
	CreateTime   int    `json:"create_time,omitempty" bson:"create_time,omitempty"`
}

type ReportUser struct {
	Uid          string `json:"uid,omitempty" bson:"uid,omitempty"`
	Reporting    string `json:"reporting,omitempty" bson:"reporting,omitempty"`
	ReportReason string `json:"report_reason,omitempty" bson:"report_reason,omitempty"`
	CreateTime   int    `json:"create_time,omitempty" bson:"create_time,omitempty"`
}

func (r Report) ReportUser(ctx context.Context, reportUserInfo *pbReport.ReportUserInfo, reportResult *pbReport.ReportResult) error {
	mgoClient := mgodb.NewMgo(utils2.GetConfigStr("mongodb.db.report"), utils2.GetConfigStr("mongodb.collection.report_user"))
	reportUser := ReportUser{
		Uid:          reportUserInfo.Uid,
		Reporting:    reportUserInfo.Reporting,
		ReportReason: reportUserInfo.ReportReason,
		CreateTime:   int(time.Now().Unix()),
	}
	err := mgoClient.InsertOne(reportUser)
	if err != utils2.RECODE_OK {
		logger.Error(err)
		reportResult.Code = err
	}
	return nil
}

func (r Report) ReportOop(ctx context.Context, reportOopInfo *pbReport.ReportOopInfo, reportResult *pbReport.ReportResult) error {
	mgoClient := mgodb.NewMgo(utils2.GetConfigStr("mongodb.db.report"), utils2.GetConfigStr("mongodb.collection.report_oop"))
	reportOop := ReportOop{
		Oid:          reportOopInfo.Oid,
		Uid:          reportOopInfo.Uid,
		Reporting:    reportOopInfo.Reporting,
		ReportReason: reportOopInfo.ReportReason,
		CreateTime:   int(time.Now().Unix()),
	}
	err := mgoClient.InsertOne(reportOop)
	if err != utils2.RECODE_OK {
		logger.Error(err)
		reportResult.Code = err
	}
	return nil
}

func (r Report) GetUserReportTimes(ctx context.Context, UserID *pbReport.UserID, reportResult *pbReport.UserReportedTimes) error {
	mgoClientReportOop := mgodb.NewMgo(utils2.GetConfigStr("mongodb.db.report"), utils2.GetConfigStr("mongodb.collection.report_oop"))
	mgoClientReportUser := mgodb.NewMgo(utils2.GetConfigStr("mongodb.db.report"), utils2.GetConfigStr("mongodb.collection.report_user"))
	reportOopsNum, _ := mgoClientReportOop.Count(bson.M{"reporting": UserID.Uid})
	reportUserNum, _ := mgoClientReportUser.Count(bson.M{"reporting": UserID.Uid})
	reportResult.Times = reportOopsNum + reportUserNum
	return nil
}
