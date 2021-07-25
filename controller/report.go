package controller

import (
	"context"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"ops/pkg/model/consulreg"
	myjwt "ops/pkg/model/jwt"
	reportModel "ops/pkg/model/report"
	"ops/pkg/utils"
	pbReport "ops/proto/report"
)

// ReportOop godoc
// @Summary 举报用户oop
// @Description 当post接口时, 存储该举报内容。
// @ID ReportOop
// @tags 举报
// @Accept  json
// @Produce  json
// @Param token header string true "带上token即可, 举报人的uid从token中取出来"
// @Param report body reportModel.ReportOop true "传入oid reporting report_reason 全部参数必须"
// @Success 0 {object} reportModel.ReportResp
// @Header 200 {string} token "token"
// @Failure 4004 {object} model.JSONResult "json的key有错，或者是value不能为空"
// @Failure 4005 {object} model.JSONResult "微服务可能挂了"
// @Failure 4006 {object} model.JSONResult "存入mongodb失败"
// @Router /report/oop [post]
func ReportOop(ctx *gin.Context) {
	var resp reportModel.ReportResp
	var req reportModel.ReportOop
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req.Uid = claims.Uid
	if err := ctx.ShouldBind(&req); err != nil {
		log.Error("bind data failed in SearchID", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}
	microClient := pbReport.NewOperateReportService("report", consulreg.MicroSer.Client())
	remote, err := microClient.ReportOop(context.TODO(), req.SetValue())
	if err != nil {
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		log.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote))
}

// ReportUser godoc
// @Summary 举报用户本身
// @Description 当post接口时, 存储该举报内容。
// @ID ReportUser
// @tags 举报
// @Accept  json
// @Produce  json
// @Param token header string true "带上token即可, 举报人的uid从token中取出来"
// @Param report body reportModel.ReportUser true "传入reporting report_reason 全部参数必须"
// @Success 0 {object} reportModel.ReportResp
// @Header 200 {string} token "token"
// @Failure 4004 {object} model.JSONResult "json的key有错，或者是value不能为空"
// @Failure 4005 {object} model.JSONResult "微服务可能挂了"
// @Failure 4006 {object} model.JSONResult "存入mongodb失败"
// @Router /report/user [post]
func ReportUser(ctx *gin.Context) {
	var resp reportModel.ReportResp
	var req reportModel.ReportUser
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req.Uid = claims.Uid
	if err := ctx.ShouldBind(&req); err != nil {
		log.Error("bind data failed in SearchID", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}
	microClient := pbReport.NewOperateReportService("report", consulreg.MicroSer.Client())
	remote, err := microClient.ReportUser(context.TODO(), req.SetValue())
	if err != nil {
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		log.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote))
}

// ReportTimes godoc
// @Summary 返回用户被举报总计次数
// @Description 当post接口时, 返回该用户合计oop 本身被举报的次数。
// @ID ReportTimes
// @tags 举报
// @Accept  json
// @Produce  json
// @Param token header string true "带上token即可"
// @Param report body pbReport.UserID true "传入uid 参数必须"
// @Success 0 {object} reportModel.ReportNumResp
// @Header 200 {string} token "token"
// @Failure 4004 {object} model.JSONResult "json的key有错，或者是value不能为空"
// @Failure 4005 {object} model.JSONResult "微服务可能挂了"
// @Router /report/times [post]
func ReportTimes(ctx *gin.Context) {
	var resp reportModel.ReportNumResp
	var req reportModel.UserID
	if err := ctx.ShouldBind(&req); err != nil {
		log.Error("bind data failed in SearchID", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}
	microClient := pbReport.NewOperateReportService("report", consulreg.MicroSer.Client())
	remote, err := microClient.GetUserReportTimes(context.TODO(), req.SetValue())
	if err != nil {
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		log.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote))
}
