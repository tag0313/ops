package controller

import (
	"context"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"ops/pkg/model/consulreg"
	searchModel "ops/pkg/model/search"
	"ops/pkg/utils"
	pbSearch "ops/proto/search"
)

// SearchID godoc
// @Summary 根据用户id查询该用户的oop。
// @Description 当post接口时，此方法根据用户uid从mongoDB查询该用户oop。
// @ID SearchID
// @tags 搜索
// @Accept  json
// @Produce  json
// @Param token header string true "带上token即可"
// @Param search body pbSearch.UserID true "传入uid page_number show_number timestamp 其中timestamp非必需"
// @Success 0 {object} searchModel.SearchContentResp
// @Header 200 {string} token "token"
// @Failure 4004 {object} model.JSONResult "json的key有错，或者是value不能为空"
// @Failure 4005 {object} model.JSONResult "微服务可能挂了"
// @Router /search/id [post]
func SearchID(ctx *gin.Context) {
	var resp searchModel.SearchContentResp
	var req searchModel.SearchContentByIDReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Error("bind data failed in SearchID", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}
	microClient := pbSearch.NewOperateSearchService("search", consulreg.MicroSer.Client())
	remote, err := microClient.SearchID(context.TODO(), req.SetValue())
	if err != nil {
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		log.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote))
}

// SearchContent godoc
// @Summary 根据content查询包含该内容的oop。
// @Description 当post接口时，此方法根据content从mongoDB查询包含该content的oop。
// @ID SearchContent
// @tags 搜索
// @Accept  json
// @Produce  json
// @Param token header string true "带上token即可"
// @Param Store body pbSearch.Content true "传入content page_number show_number timestamp 其中timestamp非必需"
// @Success 0 {object} searchModel.SearchContentResp
// @Header 200 {string} token "token"
// @Failure 4004 {object} model.JSONResult "json的key有错，或者是value不能为空"
// @Failure 4005 {object} model.JSONResult "微服务可能挂了"
// @Router /search/content [post]
func SearchContent(ctx *gin.Context) {
	var resp searchModel.SearchContentResp
	var req searchModel.SearchContentReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Error("bind data failed in SearchID", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}
	microClient := pbSearch.NewOperateSearchService("search", consulreg.MicroSer.Client())
	remote, err := microClient.SearchContent(context.TODO(), req.SetValue())
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote))
}

// SearchUser godoc
// @Summary 根据content查询ops_account或nick_name包含该内容的user。
// @Description 当post接口时，此方法根据content从mongoDB查询ops_account或nick_name包含该content的user
// @ID SearchUser
// @tags 搜索
// @Accept  json
// @Produce  json
// @Param token header string true "带上token即可"
// @Param Store body searchModel.SearchContentReq true "传入content page_number show_number"
// @Success 0 {object} searchModel.SearchUserResp
// @Header 200 {string} token "token"
// @Failure 4004 {object} model.JSONResult "json的key有错，或者是value不能为空"
// @Failure 4005 {object} model.JSONResult "微服务可能挂了"
// @Router /search/user [post]
func SearchUser(ctx *gin.Context) {
	var resp searchModel.SearchUserResp
	var req searchModel.SearchContentReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Error("bind data failed in SearchID", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}
	microClient := pbSearch.NewOperateSearchService("search", consulreg.MicroSer.Client())
	remote, err := microClient.SearchUser(context.TODO(), req.SetValue())
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote))
}
