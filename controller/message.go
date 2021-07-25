package controller

import (
	"context"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/gin-gonic/gin"
	"net/http"
	baseModel "ops/pkg/model"
	"ops/pkg/model/consulreg"
	myjwt "ops/pkg/model/jwt"
	model "ops/pkg/model/message"
	"ops/pkg/utils"
	pbMessage "ops/proto/message"
)

// GetAllMessage godoc
// @Summary 取出用户消息队列里面全部消息
// @Description 当post接口时，此方法从redis队列取出所有消息
// @ID GetAllMessage
// @tags 消息推送
// @Accept  json
// @Produce  json
// @Param token header string true "带上token即可"
// @Header 200 {string} token "token"
// @Failure 4004 {object} baseModel.JSONResult "json的key有错，或者是value不能为空"
// @Failure 4005 {object} baseModel.JSONResult "微服务可能挂了"
// @Router /message/get-all [post]
func GetAllMessage(ctx *gin.Context) {
	var resp model.MessageResp
	var req pbMessage.UserID
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req.Uid = claims.Uid
	microClient := pbMessage.NewOperateMessageService("message", consulreg.MicroSer.Client())
	remote, err := microClient.GetAll(context.TODO(), &req)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote))
}

// GetMessage godoc
// @Summary 取出用户消息队列里面指定数量消息
// @Description 当post接口时，此方法从redis队列取出指定数量消息
// @ID GetMessage
// @tags 消息推送
// @Accept  json
// @Produce  json
// @Param token header string true "带上token即可"
// @Param num body model.GetNumMessageReq true "要取出的消息数量"
// @Header 200 {string} token "token"
// @Failure 4004 {object} baseModel.JSONResult "json的key有错，或者是value不能为空"
// @Failure 4005 {object} baseModel.JSONResult "微服务可能挂了"
// @Router /message/get [post]
func GetMessage(ctx *gin.Context) {
	var resp model.MessageResp
	var req model.GetNumMessageReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.Error("bind data failed in SearchID", err.Error())
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		return
	}
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req.Uid = claims.Uid
	microClient := pbMessage.NewOperateMessageService("message", consulreg.MicroSer.Client())
	remote, err := microClient.Get(context.TODO(), req.SetValue())
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_MICROERR))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote))
}

// PushMessage godoc
// @Summary 推一个消息进队列(测试接口)
// @Description 当post接口时，此方法从向redis队列里面推送一个消息
// @ID PushMessage
// @tags 消息推送
// @Accept  json
// @Produce  json
// @Param token header string true "带上token即可"
// @Param num body model.GetNumMessageReq true "要推进去的消息 from_user可能为空，根据具体消息来"
// @Success 0 {object} baseModel.JSONResult "from_user可能为空，根据具体消息来"
// @Header 200 {string} token "token"
// @Failure 4004 {object} baseModel.JSONResult "json的key有错，或者是value不能为空"
// @Failure 4005 {object} baseModel.JSONResult "微服务可能挂了"
// @Router /message/push [post]
func PushMessage(ctx *gin.Context) {
	var resp baseModel.JSONResult
	// claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	// utils.SendCreateCardMsgPending("0xf0c993c2d541dd96ad6ca196151c464d9a00339596b5c92ebb11e45590277ee6", "213123131231", 60)
	// utils.SendCreateCardMsgFailed("0xf0c993c2d541dd96ad6ca196151c464d9a00339596b5c92ebb11e45590277ee6", "1231313131", 60)
	// utils.SendCreateCardMsgSuccess("0xf0c993c2d541dd96ad6ca196151c464d9a00339596b5c92ebb11e45590277ee6", "1231313131", 60)
	// utils.SendTradeMsgPending("0x46c98c95e56b349ea9a8aa64f5df98a2b8804c02de964c1ba6b7a91a6bf80b2c", "to_address", "from_address", claims.Uid, 213.123)
	utils.SendTradeMsgFailed("0x46c98c95e56b349ea9a8aa64f5df98a2b8804c02de964c1ba6b7a91a6bf80b2c", "to_address", "from_address", "", 213.123)
	// utils.SendTradeMsgSuccess("0x46c98c95e56b349ea9a8aa64f5df98a2b8804c02de964c1ba6b7a91a6bf80b2c", "to_address", "from_address", claims.Uid, 213.123)
	// utils.SendCardNotEnoughMsg(claims.Uid)
	// utils.SendPrivateMsg(claims.Uid, "SendPrivateMsg", "MsgSenderUid")z
	// utils.SendCommentMsg(claims.Uid, "oid", "content", "contentOid", "CommentSenderUid")
	// utils.SendForwordMsg(claims.Uid, "forworderUid", "oid")
	// utils.SendFollowMsg(claims.Uid, "followerUid")
	// utils.SendLikeMsg(claims.Uid, "oid", "uid", 1)
	ctx.JSON(http.StatusOK, resp.NewSuccess())
}
