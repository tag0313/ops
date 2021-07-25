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
	"ops/pkg/model/userInfo"
	"ops/pkg/utils"
	pbContract "ops/proto/contract"
	pbProperty "ops/proto/property"
	pbUserInfo "ops/proto/userInfo"
)

// StoreUserInfo godoc
// @Summary 储存用户信息。
// @Description 当post接口时，此方法会将用户数据存入mongodb。
// @ID StoreUserInfo
// @tags 用户信息操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid，所以不需要传别的值"
// @Param Store body userInfo.StoreReq true "传入store里边的全部信息"
// @Success 0 {object} model.JSONResult
// @Header 200 {string} token "token"
// @Failure 4004 {object} model.JSONResult "json的key有错，或者是value不能为空"
// @Failure 4005 {object} model.JSONResult "微服务可能挂了"
// @Failure 4006 {object} model.JSONResult "存入mongodb失败"
// @Failure 4124 {object} model.JSONResult "用户已经存在"
// @Router /user/store [post]
func StoreUserInfo(ctx *gin.Context) {
	var req userInfo.StoreReq
	var resp model.JSONResult
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	if err := ctx.ShouldBind(&req); err != nil {
		log.Error("bind data failed in StoreUserInfo", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}

	microClient := pbUserInfo.NewOperateUserInfoService("userInfo", consulreg.MicroSer.Client())

	//Get user info at the very first time, if no user exists, good, otherwise throw an error.
	result1, err1 := microClient.QueryUserInfo(context.TODO(), &pbUserInfo.QueryAndDelete{Uid: claims.Uid})
	if err1 != nil {
		log.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
		log.Error(err1)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	}
	if result1.Uid != utils.RECODE_USERNOEXISTSERR {
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_USERDUPLICATEDERR))
		return
	}

	remote, err := microClient.StoreUserInfo(context.TODO(), req.SetValue(claims))
	if err != nil || remote.Code != utils.RECODE_OK {
		log.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
		log.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	}
	opspoint, err := OperateOpspoint(claims.Uid, utils.GetConfigFloat64("ops_point_gift"))
	if err != nil {
		logger.Error("can't find the remote service in consul", utils.GetConfigStr("mirco.addr"))
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if opspoint != utils.RECODE_OK {
		logger.Error("can't find the remote service in consul", utils.GetConfigStr("mirco.addr"))
		ctx.JSON(http.StatusOK, resp.NewError(opspoint))
		return
	}

	//get the user pubkaddr by uid
	result, err := microClient.GetPbkAddrByUid(context.TODO(), &pbUserInfo.UidAndPbkaddr{Uid: claims.Uid})
	if err != nil {
		logger.Error("can't find the remote service in consul", utils.GetConfigStr("mirco.addr"))
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if result.ErrCode != utils.RECODE_OK {
		logger.Error("GetPbkAddrByUid problem", utils.GetConfigStr("mirco.addr"))
		ctx.JSON(http.StatusOK, resp.NewError(result.ErrCode))
		return
	}

	//注册成功之后调用智能合约自动分发空投币
	contractMicroClient := pbContract.NewContractService("contract", consulreg.MicroSer.Client())
	_, err = contractMicroClient.DonateCoin(ctx, &pbContract.DonateCoinRequest{
		UserWalletAddress: result.Pubkaddr,
	})
	if err != nil {
		logger.Error("Abnormal contract airdrop")
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_UNKNOWERR))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess())
}

// UpdateUserInfo godoc
// @Summary 更新用户信息。
// @Description 当post接口时，此方法会将更新mongodb中的用户信息,并返回更新后的数据。
// @ID UpdateUserInfo
// @tags 用户信息操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid，所以不需要传别的值"
// @Param update body userInfo.UpdateReq true "传入需要更新的任意数量信息"
// @Success 0 {object} userInfo.UserResp
// @Header 200 {string} token "token"
// @Failure 4004 {object} model.JSONResult "json的key有错，或者是value不能为空"
// @Failure 4005 {object} model.JSONResult "微服务可能挂了"
// @Failure 4006 {object} model.JSONResult "存入mongodb失败"
// @Failure 4104 {object} model.JSONResult "该ops_account已经存在"
// @Router /user/update [put]
func UpdateUserInfo(ctx *gin.Context) {
	var req userInfo.UpdateReq
	var resp userInfo.UserResp
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	if err := ctx.ShouldBind(&req); err != nil {
		log.Error("bind data failed in StoreUserInfo", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}

	microClient := pbUserInfo.NewOperateUserInfoService("userInfo", consulreg.MicroSer.Client())
	remote, err := microClient.UpdateUserInfo(context.TODO(), req.SetValue(claims))
	if err != nil {
		log.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
		log.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if remote.NickName == "" {
		ctx.JSON(http.StatusOK, resp.NewError(remote.Uid))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote))
}

// DeleteUserInfo godoc
// @Summary 删除用户信息。
// @Description 当delete接口时，此方法会将删除用户信息。
// @ID DeleteUserInfo
// @tags 用户信息操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid，所以不需要传别的值"
// @Success 0 {object} model.JSONResult
// @Header 200 {string} token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiJtY2NBc0tOOVlZamN3RjZDbzNQcSIsImV4cCI6MTYyMTY3OTgwMiwiaXNzIjoib3BzbmZ0X2p3dCIsIm5iZiI6MTYyMDgxNDgwMn0.1p55LDpZhfg2aCV-cPxV8P1Bea96MB8a3B73aVOZliI"
// @Failure 4005 {object} model.JSONResult "微服务可能挂了"
// @Failure 4006 {object} model.JSONResult "存入mongodb失败"
// @Failure 4121 {object} model.JSONResult "用户不存在"
// @Failure 4120 {object} model.JSONResult "用户的point未导出"
// @Router /user/delete [delete]
func DeleteUserInfo(ctx *gin.Context) {
	var resp model.JSONResult
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	//whether the user still have opspoint
	RpcPropertyClient := pbProperty.NewOperatePropertyService("property", consulreg.MicroSer.Client())
	point, err := RpcPropertyClient.CheckOpsPoint(context.TODO(), &pbProperty.OpsPoint{Uid: claims.Uid})
	if err != nil {
		logger.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
		logger.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if point.OpsPoint > 0 {
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_EXISTSPOINTERR))
		return
	}
	//whether the user still have ocard on ops

	//delete user info
	RpcUserInfoClient := pbUserInfo.NewOperateUserInfoService("userInfo", consulreg.MicroSer.Client())
	remote, err := RpcUserInfoClient.DeleteUserInfo(context.TODO(), &pbUserInfo.QueryAndDelete{Uid: claims.Uid})
	if err != nil {
		logger.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
		logger.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if remote.Code != utils.RECODE_OK {
		ctx.JSON(http.StatusOK, resp.NewError(remote.Code))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess())
}

// QueryUserInfo godoc
// @Summary 查询用户信息。
// @Description 当post接口时，此方法会查询出post来的对应用户uid的相关信息。
// @ID QueryUserInfo
// @tags 用户信息操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid，所以不需要传别的值"
// @Param query body userInfo.QuesryReq true "传入uid，返回对应的数据"
// @Success 0 {object} userInfo.UserResp
// @Header 200 {string} token
// Failure 4004 {object} userInfo.UserResp "json的key有错，或者是value不能为空"
// @Failure 4005 {object} userInfo.UserResp "微服务可能挂了"
// @Failure 4121 {object} userInfo.UserResp "用户不存在"
// @Router /user/query [post]
func QueryUserInfo(ctx *gin.Context) {
	var req userInfo.QuesryReq
	var resp userInfo.UserResp
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	if err := ctx.ShouldBind(&req); err == nil {
		//logger.Error("bind data failed in QueryUserInfo",)
		//ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
	}
	if req == (userInfo.QuesryReq{}) {

		req.Uid = claims.Uid
	}
	microClient := pbUserInfo.NewOperateUserInfoService("userInfo", consulreg.MicroSer.Client())
	remote, err := microClient.QueryUserInfo(context.TODO(), &pbUserInfo.QueryAndDelete{Uid: req.Uid})
	if err != nil {
		log.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
		log.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if remote.Uid != req.Uid {
		ctx.JSON(http.StatusOK, resp.NewError(remote.Uid))
		return
	} else {
		ctx.JSON(http.StatusOK, resp.NewSuccess(remote))
		return
	}

}

// IsRepeatedOpsAccount godoc
// @Summary 查询ops_account是否重复。
// @Description 当post接口时，此方法会查询出当前ops_account是否重复, 该ops_account不存在返回0，如果该用户存在返回4104。
// @ID IsRepeatedOpsAccount
// @tags 用户信息操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid，所以不需要传别的值"
// @Param query body pbUserInfo.CheckOpsAccount true "传入uid，返回对应的数据"
// @Success 0 {object} model.JSONResult "用户不存在返回0"
// @Success 4104 {object} model.JSONResult "用户存在返回4104"
// @Header 200 {string} token
// Failure 4004 {object} model.JSONResult "json的key有错，或者是value不能为空"
// @Failure 4005 {object} model.JSONResult "微服务可能挂了"
// @Router /user/query/ops-account [post]
func IsRepeatedOpsAccount(ctx *gin.Context) {
	var req pbUserInfo.CheckOpsAccount
	var resp model.JSONResult
	if err := ctx.ShouldBind(&req); err != nil {
		log.Error("bind data failed in QueryUserInfo", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
	}
	microClient := pbUserInfo.NewOperateUserInfoService("userInfo", consulreg.MicroSer.Client())
	remote, err := microClient.IsRepeatedOpsAccount(context.TODO(), &req)
	if err != nil {
		log.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
		log.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if remote.Code != utils.RECODE_OK {
		ctx.JSON(http.StatusOK, resp.NewError(remote.Code))
		return
	} else {
		ctx.JSON(http.StatusOK, resp.NewSuccess())
		return
	}
}

// SetPrivate godoc
// @Summary 设置是否为私有账户。
// @Description 当post接口时，此方法会将用户设置的是否为私有账户状态，如果true用户开启私有，数据存入mongo和redis，如果是false只存入mongo。如果用户从true切换到false，redis列表里的数据同样也会删除。
// @ID SetPrivate
// @tags 用户信息操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid，所以不需要传别的值"
// @Param query body pbUserInfo.SetPrivate true "传入is_private，boolean类型，true或false"
// @Success 0 {object} model.JSONResult "设置成功返回0"
// @Header 200 {string} token
// @Failure 4004 {object} model.JSONResult "json的key有错，或者是value不能为空"
// @Failure 4005 {object} model.JSONResult "微服务错误"
// @Router /user/private [post]
func SetPrivate(ctx *gin.Context) {
	var req pbUserInfo.SetPrivate
	var resp model.JSONResult
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	if err := ctx.ShouldBind(&req); err != nil {
		log.Error("bind data failed in QueryUserInfo", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
	}
	req.Uid = claims.Uid
	microClient := pbUserInfo.NewOperateUserInfoService("userInfo", consulreg.MicroSer.Client())
	remote, err := microClient.OneStepProtect(context.TODO(), &req)
	if err != nil {
		log.Error("The remote server may have some problem", utils.GetConfigStr("mirco.addr"))
		log.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if remote.Code != utils.RECODE_OK {
		ctx.JSON(http.StatusOK, resp.NewError(remote.Code))
		return
	} else {
		ctx.JSON(http.StatusOK, resp.NewSuccess())
		return
	}
}

// SetPrice godoc
// @Summary 设置OCard出售单价。
// @Description 当post接口时，该接口会设置重新设置用户所输入的OCard的单价。
// @ID SetPrice
// @tags 用户信息操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid，所以不需要传别的值"
// @Param set body pbUserInfo.Price true "传入price，double"
// @Success 0 {object} userInfo.PriceResp "设置成功返回0，还有设置后的价格"
// @Header 200 {string} token
// @Failure 4001 {object} userInfo.PriceResp "数据库查询错误，没有找到对应的uid。"
// @Failure 4004 {object} userInfo.PriceResp "从mgo取出数据后，需要进行序列化操作，这个错误意思是序列化出问题了。"
// @Failure 4501 {object} userInfo.PriceResp "微服务错误，或者发生了未知错误"
// @Router /user/price [post]
func SetPrice(ctx *gin.Context) {
	var req pbUserInfo.Price
	var resp userInfo.PriceResp
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	if err := ctx.ShouldBind(&req); err != nil {
		log.Error("bind data failed in QueryUserInfo", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
	}
	req.Uid = claims.Uid
	microClient := pbUserInfo.NewOperateUserInfoService("userInfo", consulreg.MicroSer.Client())
	remote, err := microClient.SetPrice(context.TODO(), &req)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if remote.Uid == utils.RECODE_DBERR {
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DBERR))
		return
	} else if remote.Uid == utils.RECODE_DATAERR {
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote))
}

// GetUidByOpsAccount godoc
// @Summary 通过用户ops_account获取用户uid。（批量）
// @Description 当post接口时，该接口会使用ops_account精确匹配uid。
// @ID GetUidByOpsAccount
// @tags 用户信息操作
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid，所以不需要传别的值"
// @Param set body pbUserInfo.OpsAccounts true "传入ops_accounts<------注意这个是有s的，因为是个数组"
// @Success 0 {object} userInfo.GetOpsAccountByUidResp "设置成功返回0，还有设置后的价格"
// @Header 200 {string} token
// @Failure 4001 {object} userInfo.GetOpsAccountByUidResp "数据库查询错误，没有找到对应的uid。"
// @Failure 4004 {object} userInfo.GetOpsAccountByUidResp "从mgo取出数据后，需要进行序列化操作，这个错误意思是序列化出问题了。"
// @Failure 4005 {object} userInfo.GetOpsAccountByUidResp "微服务错误，或者发生了未知错误"
// @Router /user/get-uid-by-account [post]
func GetUidByOpsAccount(ctx *gin.Context) {
	var (
		statusCode int
		req        pbUserInfo.OpsAccounts
		reqData    userInfo.GetOpsAccountByUidData
		response   userInfo.GetOpsAccountByUidResp
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils.RECODE_DATAERR, "bind data failed in GetOpsAccountByUid", err)
		statusCode = http.StatusBadRequest
		return
	}
	logger.Info(req)
	microClient := pbUserInfo.NewOperateUserInfoService("userInfo", consulreg.MicroSer.Client())
	remote, err := microClient.GetUidByOpsAccount(context.TODO(), &req)
	if err != nil {
		response.SetError(utils.RECODE_MICROERR, "the micro service may have problem", err)
		statusCode = http.StatusBadRequest
		return
	} else if len(remote.OpsAndUid) == 0 {
		statusCode = http.StatusOK
		response.NewSuccess()
		return
	}
	for index := range remote.OpsAndUid {
		reqData.OpsAccount = remote.OpsAndUid[index].OpsAccount
		reqData.Uid = remote.OpsAndUid[index].Uid
		response.Data = append(response.Data, reqData)
	}
	statusCode = http.StatusOK
	response.NewSuccess()
}
