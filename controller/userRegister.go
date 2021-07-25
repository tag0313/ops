package controller

import (
	"context"
	"github.com/asim/go-micro/v3/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"ops/pkg/model/consulreg"
	"ops/pkg/model/user"
	"ops/pkg/utils"
	pbUserRegister "ops/proto/userRegister"
)

// PostMessage godoc
// @Summary 获取校验所需要的消息。
// @Description 当post接口时，此方法会返回一个128位的随机字符串，用于用户登录与注册时进行的校验行为。
// @ID PostMessage
// @tags 用户注册与登录
// @Accept  json
// @Produce  json
// @Param pbkaddr body user.GenerateMsgRequest true "传入public key address"
// @Success 0 {object} user.GenerateMsgResponse
// @Header 200 {header} string
// @Failure 4004 {object} user.GenerateMsgResponse "json的key有错，或者是value不能为空"
// @Failure 4005 {object} user.GenerateMsgResponse "微服务可能挂了"
// @Failure 4006 {object} user.GenerateMsgResponse "msg存入redis失败"
// @Router /login/message [post]
func PostMessage(ctx *gin.Context) {
	var req user.GenerateMsgRequest
	var resp user.GenerateMsgResponse
	if err := ctx.ShouldBind(&req); err != nil {
		logger.Error("bind data failed in PostMessage", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}
	microClient := pbUserRegister.NewUserRegisterService("userRegister", consulreg.MicroSer.Client())
	remote, err := microClient.GenerateMessage(context.TODO(), &pbUserRegister.PublickeyAddr{PbkAddr: req.PbkAddr})
	if err != nil {
		logger.Error("can't find the remote service in consul", utils.GetConfigStr("mirco.addr"))
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if remote.RandomCode == utils.RECODE_STOREDATA_FAILED {
		ctx.JSON(http.StatusOK, resp.NewError(remote.RandomCode))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote.RandomCode))
}

// PostToken godoc
// @Summary 校验pbk、msg、sign是否正确。
// @Description 当post接口时，此方法校验pbk、msg、sign是否正确，用于用户登录进行的校验行为。
// @ID PostToken
// @tags 用户注册与登录
// @Accept  json
// @Produce  json
// @Param pbk,sign body user.LoginRequest true "传入publickey和 sign"
// @Success 0 {object} user.LoginResponse
// @Header 200 {header} string
// @Failure 4004 {object} user.LoginResponse "json的key有错，或者是value不能为空"
// @Failure 4005 {object} user.LoginResponse "微服务可能挂了"
// @Failure 4006 {object} user.LoginResponse "存入mongodb失败"
// @Failure 4007 {object} user.LoginResponse "在redis里边没有找到对应的msg"
// @Failure 4102 {object} user.LoginResponse "签名验证失败"
// @Failure 4114 {object} user.LoginResponse "Token生成失败"
// @Router /login/token [post]
func PostToken(ctx *gin.Context) {
	var req user.LoginRequest
	var resp user.LoginResponse
	if err := ctx.ShouldBind(&req); err != nil {
		logger.Error("bind data failed in PostToken", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}

	microClient := pbUserRegister.NewUserRegisterService("userRegister", consulreg.MicroSer.Client())
	remote, err := microClient.GenerateToken(context.TODO(), &pbUserRegister.EncryptedValue{
		PublicKey: req.Pbk,
		Sign:      req.Sign,
	})

	if err != nil {
		logger.Error("can't find the remote service in consul", utils.GetConfigStr("mirco.addr"))
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	} else if remote.Token == utils.RECODE_DATAINEXISTENCE {
		ctx.JSON(http.StatusOK, resp.NewError(remote.Token))
		return
	} else if remote.Token == utils.RECODE_STOREDATA_FAILED {
		ctx.JSON(http.StatusOK, resp.NewError(remote.Token))
		return
	} else if remote.Token == utils.RECODE_GENERATETOKENERR {
		ctx.JSON(http.StatusOK, resp.NewError(remote.Token))
		return
	} else if remote.Token == utils.RECODE_LOGINERR {
		ctx.JSON(http.StatusOK, resp.NewError(remote.Token))
		return
	}

	ctx.JSON(http.StatusOK, resp.NewSuccess(remote.Token))
}
