package handler

import (
	"context"
	"encoding/hex"
	"github.com/asim/go-micro/v3/logger"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/sha3"
	"ops/pkg/model/mgodb"
	"ops/pkg/model/rdb"
	"ops/pkg/utils"
	"ops/proto/userRegister"
	"ops/service/userRegister/data"
	"strings"
	"time"
)

type UserRegister struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *UserRegister) GenerateMessage(ctx context.Context, req *pbUserRegister.PublickeyAddr, rsp *pbUserRegister.RandomCode) error {
	logger.Debug("Received UserRegister.GenerateRandomCode request")
	rsp.RandomCode = utils.NewLen(128)
	errCode := rdb.SetS(strings.ToLower(req.PbkAddr), rsp.RandomCode, time.Second*30)
	if errCode == utils.RECODE_STOREDATA_FAILED {
		logger.Error("failed: set random code into redis")
		rsp.RandomCode = errCode
	}
	return nil
}

func (e *UserRegister) GenerateToken(ctx context.Context, value *pbUserRegister.EncryptedValue, token *pbUserRegister.Token) error {
	logger.Debug("Received UserRegister.GetToken request")
	pbk := utils.PbkToAddr(value.PublicKey) //传过来的公钥,转成地址值
	msg, errCode := rdb.Get(pbk)            //使用公钥hash扫描redid中是否存在这个随机码
	if errCode == utils.RECODE_DATAINEXISTENCE {
		token.Token = errCode
	} else if verifyLogin(value.PublicKey, msg, value.Sign) {
		mgoclient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db"), utils.GetConfigStr("mongodb.collection"))
		bytes, err := mgoclient.FindOne(bson.M{"pubkaddr": pbk}).DecodeBytes()
		var userId string
		var userLogin *data.Login
		if err != nil {
			userId = utils.NewLen(20)
			userLogin = &data.Login{
				Uid:      userId,
				Pubkaddr: pbk,
			}
			errCode = mgoclient.InsertOne(userLogin)
			if errCode == utils.RECODE_STOREDATA_FAILED {
				token.Token = errCode
			}

		} else {
			userId = bytes.Lookup("uid").StringValue()
			userLogin = &data.Login{
				Uid:      userId,
				Pubkaddr: pbk,
			}
		}

		logResult, errCode := utils.GeneratedToken(userLogin)
		if errCode == utils.RECODE_GENERATETOKENERR {
			token.Token = errCode
		}
		token.Token = logResult.Token
		token.Uid = userId
	} else {
		token.Token = utils.RECODE_LOGINERR
	}
	return nil
}

func verifyLogin(pbk, msg, sign string) bool {
	keydata, _ := hex.DecodeString(pbk)
	msgdata := sha3.Sum256([]byte(msg))
	msgdata2 := msgdata[:]
	sigdata, _ := hex.DecodeString(sign)
	return secp256k1.VerifySignature(keydata, msgdata2, sigdata[0:64])
}
