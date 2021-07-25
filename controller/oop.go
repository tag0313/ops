package controller

import (
	"context"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"ops/pkg/model"
	"ops/pkg/model/consulreg"
	myjwt "ops/pkg/model/jwt"
	"ops/pkg/model/oopResp"
	"ops/pkg/utils"
	pbOop "ops/proto/oop"
	pbUserInfo "ops/proto/userInfo"
)

// StoreOop godoc
// @Summary 储存用户oop。
// @Description 当post接口时，此方法会将用户oop存入mongodb,并返回存储后的结果，并且用户oop总数量+1。
// @ID StoreOop
// @tags 用户oop操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid"
// @Param Store body pbOop.Oop true "传个content就行"
// @Success 0 {object} oopResp.Resp "直接就返回用户要存的数据，省得再去查询一次，但是如果在数据库中的shard、comment、like_times的三个字段是0的话，那么返回的数据连这个字段都会没有，因为返回0是没有意义的。"
// @Header 200 {string} token "token"
// @Failure 4004 {object} oopResp.Resp "json的key有错，或者是value不能为空"
// @Failure 4005 {object} oopResp.Resp "微服务可能挂了"
// @Failure 4006 {object} oopResp.Resp "存入mongodb失败"
// @Router /oop/store [post]
func StoreOop(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req := &pbOop.Oop{}
	resp := &oopResp.Resp{}
	req.Uid = claims.Uid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error("bind data failed in StoreOop", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}

	microClient := pbOop.NewOperateOopService("oop", consulreg.MicroSer.Client())
	remote, err := microClient.StoreOop(context.TODO(), req)
	if err != nil {
		logger.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"), ": StoreOop")
		logger.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if remote.Code != utils.RECODE_OK {
		logger.Error("Error of remote", utils.RecodeTest(remote.Code))
		ctx.JSON(http.StatusOK, resp.NewError(remote.Code))
		return
	}
	userInfo := pbUserInfo.NewOperateUserInfoService("userInfo", consulreg.MicroSer.Client())
	result, err := userInfo.ModifyOopNum(context.TODO(), &pbUserInfo.OopNum{Uid: claims.Uid, Num: 1})
	if err != nil {
		logger.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"), ": ModifyOopNum")
		logger.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if result.Code != remote.Code {
		logger.Error("Error of remote", utils.RecodeTest(result.Code))
		ctx.JSON(http.StatusOK, resp.NewError(result.Code))
		return
	}

	ctx.JSON(http.StatusOK, resp.NewSuccess(remote.Data))
}

// UpdateOop godoc
// @Summary 更新用户oop。
// @Description 当put接口时，此方法会更新用户oop存入mongodb，并返回存储后的结果。
// @ID UpdateOop
// @tags 用户oop操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid"
// @Param update body pbOop.Oop true "传oid和content，oid为oop的唯一id，content是修改后的内容"
// @Success 0 {object} oopResp.Resp "直接就返回用户要存的数据，省得再去查询一次，但是如果在数据库中的shard、comment、like_times的三个字段是0的话，那么返回的数据连这个字段都会没有，因为返回0是没有意义的。"
// @Header 200 {string} token "token"
// @Failure 4004 {object} oopResp.Resp "json的key有错，或者是value不能为空"
// @Failure 4005 {object} oopResp.Resp "微服务可能挂了"
// @Failure 4006 {object} oopResp.Resp "存入mongodb失败"
// @Router /oop/update [put]
func UpdateOop(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req := &pbOop.Oop{}
	resp := &oopResp.Resp{}
	req.Uid = claims.Uid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error("bind data failed in UpdateOop", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}

	microClient := pbOop.NewOperateOopService("oop", consulreg.MicroSer.Client())
	remote, err := microClient.UpdateOop(context.TODO(), req)
	if err != nil {
		logger.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
		logger.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if remote.Code != utils.RECODE_OK {
		logger.Error("Error of remote", utils.RecodeTest(remote.Code))
		ctx.JSON(http.StatusOK, resp.NewError(remote.Code))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote.Data))
}

// DeleteOop godoc
// @Summary 删除用户oop。
// @Description 当delete接口时，此方法会删除用户oop,并且用户oop总数量-1。
// @ID DeleteOop
// @tags 用户oop操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid"
// @Param delete body pbOop.Delete true "传一个oid就行，oop的唯一id"
// @Success 0 {object} model.JSONResult "直接就返回用户要存的数据，省得再去查询一次"
// @Header 200 {string} token "token"
// @Failure 4004 {object} model.JSONResult "json的key有错，或者是value不能为空"
// @Failure 4005 {object} model.JSONResult "微服务可能挂了"
// @Failure 4006 {object} model.JSONResult "删除失败"
// @Router /oop/delete [delete]
func DeleteOop(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req := &pbOop.Delete{}
	resp := &model.JSONResult{}
	req.Uid = claims.Uid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error("bind data failed in DeleteOop", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}

	microClient := pbOop.NewOperateOopService("oop", consulreg.MicroSer.Client())
	remote, err := microClient.DeleteOop(context.TODO(), req)
	if err != nil {
		logger.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
		logger.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if remote.Code != utils.RECODE_OK {
		logger.Error("Error of remote", utils.RecodeTest(remote.Code))
		ctx.JSON(http.StatusOK, resp.NewError(remote.Code))
		return
	}

	userInfo := pbUserInfo.NewOperateUserInfoService("userInfo", consulreg.MicroSer.Client())
	result, err := userInfo.ModifyOopNum(context.TODO(), &pbUserInfo.OopNum{Uid: claims.Uid, Num: -1})
	if err != nil {
		logger.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"), ": ModifyOopNum")
		logger.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if result.Code != remote.Code {
		logger.Error("Error of remote", utils.RecodeTest(result.Code))
		ctx.JSON(http.StatusOK, resp.NewError(result.Code))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess())
}

// QueryOop godoc
// @Summary 查询用户oop。
// @Description 当query接口时，此方法会分页返回用户自己50条固定的oop，不区分公私。
// @ID QueryOop
// @tags 用户oop操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid，"
// @Param query body pbOop.Query true "uid不用传，page_number初始要传个1过来，给前端返回多少条，后端这边是写死的返回50条，如果返回的数据不足50条，或者为空的话，那就是说明那一页后边没数据了。timestamp发string过来，格式是一个精确到秒的数据。例子：1621935650 = 2021-05-25 19:40:50，前端那边解析完等于这个时间咱们就对上了"
// @Success 0 {object} oopResp.QueryResp "返回用户所有的oop"
// @Header 200 {string} token "token"
// @Failure 4004 {object} oopResp.QueryResp "json的key有错，或者是value不能为空"
// @Failure 4005 {object} oopResp.QueryResp "微服务可能挂了"
// @Failure 4006 {object} oopResp.QueryResp "存入mongodb失败"
// @Router /oop/query [post]
func QueryOop(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req := &pbOop.Query{}
	resp := &oopResp.QueryResp{}
	req.Uid = claims.Uid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error("bind data failed in QueryOop", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}
	req.ShowNumber = utils.CommonShowNumber
	microClient := pbOop.NewOperateOopService("oop", consulreg.MicroSer.Client())
	remote, err := microClient.QueryOwnerOop(context.TODO(), req)
	if err != nil {
		logger.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
		logger.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if remote.Code != utils.RECODE_OK {
		logger.Error("Error of remote", utils.RecodeTest(remote.Code))
		ctx.JSON(http.StatusOK, resp.NewError(remote.Code))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote))
}

// QueryOwnerOop godoc
// @Summary 查询用户oop。
// @Description 当query接口时，此方法会分页返回用户自己50条固定的oop，不区分公私。
// @ID QueryOwnerOop
// @tags 用户oop操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid，"
// @Param query body pbOop.Query true "uid不用传，page_number初始要传个1过来，给前端返回多少条，后端这边是写死的返回50条，如果返回的数据不足50条，或者为空的话，那就是说明那一页后边没数据了。timestamp发string过来，格式是一个精确到秒的数据。例子：1621935650 = 2021-05-25 19:40:50，前端那边解析完等于这个时间咱们就对上了"
// @Success 0 {object} oopResp.QueryResp "返回用户所有的oop"
// @Header 200 {string} token "token"
// @Failure 4004 {object} oopResp.QueryResp "json的key有错，或者是value不能为空"
// @Failure 4005 {object} oopResp.QueryResp "微服务可能挂了"
// @Failure 4006 {object} oopResp.QueryResp "存入mongodb失败"
// @Router /oop/owner [post]
func QueryOwnerOop(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req := &pbOop.Query{}
	resp := &oopResp.QueryResp{}
	req.Uid = claims.Uid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error("bind data failed in QueryOop", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}
	req.ShowNumber = utils.CommonShowNumber
	microClient := pbOop.NewOperateOopService("oop", consulreg.MicroSer.Client())
	remote, err := microClient.QueryOwnerOop(context.TODO(), req)
	if err != nil {
		logger.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
		logger.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if remote.Code != utils.RECODE_OK {
		logger.Error("Error of remote", utils.RecodeTest(remote.Code))
		ctx.JSON(http.StatusOK, resp.NewError(remote.Code))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote))
}

// QueryOtherOop godoc
// @Summary 查询用户oop。
// @Description 当query接口时，此方法会分页返回用户自己50条固定的oop，只返回用户公开的oop。
// @ID QueryOtherOop
// @tags 用户oop操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid，"
// @Param query body pbOop.Query true "uid不用传，page_number初始要传个1过来，给前端返回多少条，后端这边是写死的返回50条，如果返回的数据不足50条，或者为空的话，那就是说明那一页后边没数据了。timestamp发string过来，格式是一个精确到秒的数据。例子：1621935650 = 2021-05-25 19:40:50，前端那边解析完等于这个时间咱们就对上了"
// @Success 0 {object} oopResp.QueryResp "返回用户所有的oop"
// @Header 200 {string} token "token"
// @Failure 4004 {object} oopResp.QueryResp "json的key有错，或者是value不能为空"
// @Failure 4005 {object} oopResp.QueryResp "微服务可能挂了"
// @Failure 4006 {object} oopResp.QueryResp "存入mongodb失败"
// @Router /oop/other [post]
func QueryOtherOop(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req := &pbOop.Query{}
	resp := &oopResp.QueryResp{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error("bind data failed in QueryOop", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}
	req.ShowNumber = utils.CommonShowNumber
	req.SelfUid = claims.Uid
	microClient := pbOop.NewOperateOopService("oop", consulreg.MicroSer.Client())
	remote, err := microClient.QueryOtherOop(context.TODO(), req)
	if err != nil {
		logger.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
		logger.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if remote.Code != utils.RECODE_OK {
		logger.Error("Error of remote", utils.RecodeTest(remote.Code))
		ctx.JSON(http.StatusOK, resp.NewError(remote.Code))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote))
}

// SquareOop godoc
// @Summary 广场oop。
// @Description 当调用Square接口时，此方法会分页返回50条固定的oop。！！！注意！！！这个接口只会返回oop权限为公开的oop，现在oop新加了字段is_private来判断oop是否为公开，如果这个值是false，那么这条oop就会返回。
// @ID SquareOop
// @tags 用户oop操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid，"
// @Param query body pbOop.Query true "uid不用传，page_number初始要传个1过来，给前端返回多少条，后端这边是写死的返回50条，如果返回的数据不足50条，或者为空的话，那就是说明那一页后边没数据了。timestamp发string过来，格式是精确到秒的数据。例子：1621935650 = 2021-05-25 19:40:50，前端那边解析完等于这个时间咱们就对上了"
// @Success 0 {object} oopResp.QueryResp "返回用户所有的oop"
// @Header 200 {string} token "token"
// @Failure 4004 {object} oopResp.QueryResp "json的key有错，或者是value不能为空"
// @Failure 4005 {object} oopResp.QueryResp "微服务可能挂了"
// @Failure 4006 {object} oopResp.QueryResp "存入mongodb失败"
// @Router /oop/square [post]
func SquareOop(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req := &pbOop.Query{}
	resp := &oopResp.QueryResp{}
	req.Uid = claims.Uid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error("bind data failed in QueryOop", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}
	req.ShowNumber = utils.CommonShowNumber
	microClient := pbOop.NewOperateOopService("oop", consulreg.MicroSer.Client())
	remote, err := microClient.SquareOop(context.TODO(), req)
	if err != nil {
		logger.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
		logger.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if remote.Code != utils.RECODE_OK {
		logger.Error("Error of remote", utils.RecodeTest(remote.Code))
		ctx.JSON(http.StatusOK, resp.NewError(remote.Code))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote))
}

// LikeOop godoc
// @Summary 点赞oop。
// @Description 如果用户已点赞该oop则为修改
// @ID Like
// @tags 用户oop操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid，"
// @Param query body oop.Like true "uid不用传，如果用户已点赞再次调用则为取消点赞"
// @Success 0 {object} oopResp.Resp "返回用户所有的oop"
// @Header 200 {string} token "token"
// @Failure 4004 {object} oopResp.QueryResp "json的key有错，或者是value不能为空"
// @Failure 4005 {object} oopResp.QueryResp "微服务可能挂了"
// @Failure 4006 {object} oopResp.QueryResp "存入mongodb失败"
// @Router /oop/like [post]
func LikeOop(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req := &pbOop.Like{}
	resp := &oopResp.Resp{}
	req.Uid = claims.Uid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error("bind data failed in StoreOop", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}

	microClient := pbOop.NewOperateOopService("oop", consulreg.MicroSer.Client())
	remote, err := microClient.LikeOop(context.TODO(), req)
	if err != nil {
		logger.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"), ": StoreOop")
		logger.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if remote.Code != utils.RECODE_OK {
		logger.Error("Error of remote", utils.RecodeTest(remote.Code))
		ctx.JSON(http.StatusOK, resp.NewError(remote.Code))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(nil))
}

// MyLikeOop godoc
// @Summary 我的点赞记录。
// @Description 查询用户的点赞记录
// @ID MyLike
// @tags 用户oop操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid，"
// @Param query body oop.MyLike true "uid不用传，page_number初始要传个1过来"
// @Success 0 {object} oopResp.MyLikeResp "返回所有点赞过的oop"
// @Header 200 {string} token "token"
// @Failure 4004 {object} oopResp.QueryResp "json的key有错，或者是value不能为空"
// @Failure 4005 {object} oopResp.QueryResp "微服务可能挂了"
// @Failure 4006 {object} oopResp.QueryResp "存入mongodb失败"
// @Router /oop/my-like-list [post]
func MyLikeOop(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req := &pbOop.MyLike{}
	resp := &oopResp.MyLikeResp{}
	req.Uid = claims.Uid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error("bind data failed in StoreOop", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}

	microClient := pbOop.NewOperateOopService("oop", consulreg.MicroSer.Client())
	remote, err := microClient.MyLikeOop(context.TODO(), req)
	if err != nil {
		logger.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"), ": StoreOop")
		logger.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if remote.Code != utils.RECODE_OK {
		logger.Error("Error of remote", utils.RecodeTest(remote.Code))
		ctx.JSON(http.StatusOK, resp.NewError(remote.Code))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote))
}

// CancelLikeOop godoc
// @Summary 取消点赞。
// @Description 取消电站
// @ID CancelLike
// @tags 用户oop操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid，"
// @Param query body oop.CancelLike true "uid不用传"
// @Success 0 {object} model.JSONResult "取消点赞是否成功"
// @Header 200 {string} token "token"
// @Failure 4004 {object} oopResp.QueryResp "json的key有错，或者是value不能为空"
// @Failure 4005 {object} oopResp.QueryResp "微服务可能挂了"
// @Failure 4006 {object} oopResp.QueryResp "存入mongodb失败"
// @Router /oop/cancel-like [post]
func CancelLikeOop(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req := &pbOop.CancelLike{}
	resp := &model.JSONResult{}
	req.Uid = claims.Uid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error("bind data failed in StoreOop", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}

	microClient := pbOop.NewOperateOopService("oop", consulreg.MicroSer.Client())
	remote, err := microClient.CancelLikeOop(context.TODO(), req)
	if err != nil {
		logger.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"), ": StoreOop")
		logger.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if remote.Code != utils.RECODE_OK {
		logger.Error("Error of remote", utils.RecodeTest(remote.Code))
		ctx.JSON(http.StatusOK, resp.NewError(remote.Code))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess())
}

