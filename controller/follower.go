package controller

import (
	"context"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"ops/pkg/model/consulreg"
	followerResp "ops/pkg/model/follower"
	myjwt "ops/pkg/model/jwt"
	"ops/pkg/utils"
	pbFollower "ops/proto/follower"
	pbOop "ops/proto/oop"
	pbUserInfo "ops/proto/userInfo"
)

// Following godoc
// @Summary 当前用户关注其他用户时调用
// @Description 当put接口时，此方法会添加一条关注数据到mongodb，并且在当前用户信息中的关注数+1，同时被关注人的关注数也会+1。
// @ID Following
// @tags Follower操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid"
// @Param Store body pbFollower.Follower true "传following就行，它的意思是被关注用户的uid"
// @Success 0 {object} followerResp.Resp "返回关注成功后，当前用户的关注数与被关注数"
// @Header 200 {string} token "token"
// @Failure 4004 {object} followerResp.Resp "json的key有错，或者是value不能为空"
// @Failure 4005 {object} followerResp.Resp "微服务可能挂了"
// @Failure 4006 {object} followerResp.Resp "存入mongodb失败"
// @Router /follow/following [put]
func Following(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req := &pbFollower.Follower{}
	resp := &followerResp.Resp{}
	req.Uid = claims.Uid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error("bind data failed in Following", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}

	microClient := pbFollower.NewOperateFollowService("follower", consulreg.MicroSer.Client())
	remote, err := microClient.Follow(context.TODO(), req)
	if err != nil {
		logger.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
		logger.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote))
}

// CancelFollow godoc
// @Summary 当前用户取消关注其他用户时调用
// @Description 当DELETE接口时，此方法会删除一条mongodb中的关注数据，并且在当前用户信息中的关注数-1，同时被关注人的关注数也会-1。
// @ID CancelFollow
// @tags Follower操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid"
// @Param Store body pbFollower.Follower true "传following，它的意思是被关注用户的uid"
// @Success 0 {object} followerResp.Resp "返回取消关注后，当前用户的关注数与被关注数"
// @Header 200 {string} token "token"
// @Failure 4004 {object} followerResp.Resp "json的key有错，或者是value不能为空"
// @Failure 4005 {object} followerResp.Resp "微服务可能挂了"
// @Router /follow/cancel [delete]
func CancelFollow(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req := &pbFollower.Follower{}
	resp := &followerResp.Resp{}
	req.Uid = claims.Uid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error("bind data failed in Following", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}

	microClient := pbFollower.NewOperateFollowService("follower", consulreg.MicroSer.Client())
	remote, err := microClient.CanalFollow(context.TODO(), req)
	if err != nil {
		log.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
		log.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote))
}

// QueryFollowOop godoc
// @Summary 查询用户关注了的其他用户的oop
// @Description 当post接口时，此方法会返回当前用户关注的其他用户的oop。
// @ID QueryFollowOop
// @tags Follower操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid"
// @Param followerResp.OopReq body followerResp.OopReq true "如果传过来的uid是空的，那该接口会自动使用token中的uid，如果timestamp字段为0或者没有传，那么就会返回这些用户的最新的50条oop"
// @Success 0 {object} followerResp.OopResp "返回当前用户关注的其他用户的oop列表，一次最多返回50条"
// @Header 200 {string} token "token"
// @Failure 4004 {object} followerResp.OopResp "json的key有错，或者是value不能为空"
// @Failure 4005 {object} followerResp.OopResp "微服务可能挂了"
// @Router /follow/oop [post]
func QueryFollowOop(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req := &followerResp.OopReq{}
	resp := &followerResp.OopResp{}
	if err := ctx.ShouldBind(&req); err == nil {
		logger.Error("bind data failed in QueryFollowOop", err)
		//ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		//return
		if req.Uid == "" {
			req.Uid = claims.Uid
		}
	}

	RpcFollowClient := pbFollower.NewOperateFollowService("follower", consulreg.MicroSer.Client())
	remote, err := RpcFollowClient.QueryFollowListAll(context.TODO(), &pbFollower.OopFollowListReq{Uid: req.Uid, Timestamp: req.Timestamp})
	if err != nil {
		log.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
		log.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if remote.Code != utils.RECODE_OK {
		ctx.JSON(http.StatusOK, resp.NewError(remote.Code))
		return
	}
	logger.Info(remote)

	//add the user uid into the list
	followOop := &pbOop.FollowOop{}
	followOop.Uid = append(remote.Uid, claims.Uid)
	followOop.Timestamp = req.Timestamp
	followOop.ShowNumber = utils.CommonShowNumber
	followOop.PageNumber = req.PageNumber
	followOop.SelfUid = claims.Uid
	RpcOopClient := pbOop.NewOperateOopService("oop", consulreg.MicroSer.Client())
	queryFollowOop, err := RpcOopClient.QueryFollowOop(context.TODO(), followOop)
	if err != nil {
		log.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
		log.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if queryFollowOop.Code != utils.RECODE_OK {
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(queryFollowOop))
	logger.Debug(resp.NewSuccess(queryFollowOop))
}

// QueryFollowingList godoc
// @Summary 查询用户关注列表中的其他用户的详细信息
// @Description 当post接口时，此方法会返回当前用户关注列表中的其他用户的信息。
// @ID QueryFollowingList
// @tags Follower操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid"
// @Param Store body pbFollower.FollowListReq true "如果传过来的uid是空的，那该接口会自动使用token中的uid。"
// @Success 0 {object} followerResp.FollowListResp "返回当前用户关注的其他用户信息"
// @Header 200 {string} token "token"
// @Failure 4004 {object} followerResp.FollowListResp "json的key有错，或者是value不能为空"
// @Failure 4005 {object} followerResp.FollowListResp "微服务可能挂了"
// @Router /follow/following-list [post]
func QueryFollowingList(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req := &pbFollower.FollowListReq{}
	resp := &followerResp.FollowListResp{}
	if err := ctx.ShouldBind(&req); err == nil {
		if req.Uid == "" {
			req.Uid = claims.Uid
		}
		//logger.Error("bind data failed in QueryUserInfo",)
		//ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
	}
	req.ShowNumber = utils.CommonShowNumber
	RpcFollowListClient := pbFollower.NewOperateFollowService("follower", consulreg.MicroSer.Client())
	followList, err := RpcFollowListClient.QueryFollowingList(context.TODO(), req)
	if err != nil {
		logger.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
		logger.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if followList.Code != utils.RECODE_OK {
		logger.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(followList.Code))
		return
	}
	userInfoList := make([]*pbUserInfo.UserInfo, len(followList.Uid))
	for index, value := range followList.Uid {
		RpcUserInfoClient := pbUserInfo.NewOperateUserInfoService("userInfo", consulreg.MicroSer.Client())
		info, err := RpcUserInfoClient.QueryUserInfo(context.TODO(), &pbUserInfo.QueryAndDelete{Uid: value})
		if err != nil {
			log.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
			log.Error(err)
			ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
			return
		} else if info.Uid != value {
			ctx.JSON(http.StatusOK, resp.NewError(info.Uid))
			return
		}
		userInfoList[index] = info
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(userInfoList))
}

// QueryFollowedList godoc
// @Summary 查询用户被关注列表中的其他用户的详细信息
// @Description 当post接口时，此方法会返回当前用户被关注列表中的其他用户的信息。
// @ID QueryFollowedList
// @tags Follower操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid"
// @Param Store body pbFollower.FollowListReq true "如果传过来的uid是空的，那该接口会自动使用token中的uid。"
// @Success 0 {object} followerResp.FollowListResp "返回当前用户被关注的用户信息"
// @Header 200 {string} token "token"
// @Failure 4004 {object} followerResp.FollowListResp "json的key有错，或者是value不能为空"
// @Failure 4005 {object} followerResp.FollowListResp "微服务可能挂了"
// @Router /follow/followed-list [post]
func QueryFollowedList(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req := &pbFollower.FollowListReq{}
	resp := &followerResp.FollowListResp{}
	if err := ctx.ShouldBind(&req); err == nil {
		if req.Uid == "" {
			req.Uid = claims.Uid
		}
		//logger.Error("bind data failed in QueryUserInfo",)
		//ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
	}
	req.ShowNumber = utils.CommonShowNumber
	RpcFollowListClient := pbFollower.NewOperateFollowService("follower", consulreg.MicroSer.Client())
	followList, err := RpcFollowListClient.QueryFollowedList(context.TODO(), req)
	if err != nil {
		logger.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
		logger.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if followList.Code != utils.RECODE_OK {
		logger.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(followList.Code))
		return
	}
	userInfoList := make([]*pbUserInfo.UserInfo, len(followList.Uid))
	for index, value := range followList.Uid {
		RpcUserInfoClient := pbUserInfo.NewOperateUserInfoService("userInfo", consulreg.MicroSer.Client())
		info, err := RpcUserInfoClient.QueryUserInfo(context.TODO(), &pbUserInfo.QueryAndDelete{Uid: value})
		if err != nil {
			logger.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
			logger.Error(err)
			ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
			return
		} else if info.Uid != value {
			ctx.JSON(http.StatusOK, resp.NewError(info.Uid))
			return
		}
		userInfoList[index] = info
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(userInfoList))
}

// WhoFollowingMe godoc
// @Summary 查询一批用户是否与当前用户有关注关系
// @Description 当post接口时，会返回一批用户是否与当前用户有关注关系
// @ID WhoFollowingMe
// @tags Follower操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid"
// @Param Store body pbFollower.RelationReq true "如果传过来的uid是空的，那该接口会自动使用token中的uid。"
// @Success 0 {object} followerResp.RelationResp "返回当前用户被关注的用户信息"
// @Header 200 {string} token "token"
// @Failure 4004 {object} followerResp.RelationResp "json的key有错，或者是value不能为空"
// @Failure 4005 {object} followerResp.RelationResp "微服务可能挂了"
// @Router /follow/who-following-list [post]
func WhoFollowingMe(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req := &pbFollower.RelationReq{}
	resp := &followerResp.RelationResp{}
	if err := ctx.ShouldBind(&req); err == nil {
		if req.Uid == "" {
			req.Uid = claims.Uid
		}
		//logger.Error("bind data failed in QueryUserInfo",)
		//ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
	}
	RpcFollowListClient := pbFollower.NewOperateFollowService("follower", consulreg.MicroSer.Client())
	list, err := RpcFollowListClient.WhoFollowingMe(context.TODO(), req)
	if err != nil {
		logger.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
		logger.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if list.Code == utils.RECODE_DATAINEXISTENCE {
		ctx.JSON(http.StatusOK, resp.NewError(list.Code))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(list.Uid))
}

// WhoFollowedMe godoc
// @Summary 查询一批用户是否与当前用户有被关注关系
// @Description 当post接口时，会返回一批用户是否与当前用户有被关注关系
// @ID WhoFollowedMe
// @tags Follower操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid"
// @Param Store body pbFollower.RelationReq true "如果传过来的uid是空的，那该接口会自动使用token中的uid。"
// @Success 0 {object} followerResp.RelationResp "返回当前用户被关注的用户信息"
// @Header 200 {string} token "token"
// @Failure 4004 {object} followerResp.RelationResp "json的key有错，或者是value不能为空"
// @Failure 4005 {object} followerResp.RelationResp "微服务可能挂了"
// @Router /follow/who-followed-list [post]
func WhoFollowedMe(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req := &pbFollower.RelationReq{}
	resp := &followerResp.RelationResp{}
	if err := ctx.ShouldBind(&req); err == nil {
		if req.Uid == "" {
			req.Uid = claims.Uid
		}
		//logger.Error("bind data failed in QueryUserInfo",)
		//ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
	}

	RpcFollowListClient := pbFollower.NewOperateFollowService("follower", consulreg.MicroSer.Client())
	list, err := RpcFollowListClient.WhoFollowedMe(context.TODO(), req)
	if err != nil {
		logger.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
		logger.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if list.Code == utils.RECODE_DATAINEXISTENCE {
		ctx.JSON(http.StatusOK, resp.NewError(list.Code))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(list.Uid))
}
