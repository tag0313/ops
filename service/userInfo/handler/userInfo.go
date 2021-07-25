package handler

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/asim/go-micro/v3/logger"
	"github.com/golang/protobuf/jsonpb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"ops/pkg/model/azureStorage"
	"ops/pkg/model/mgodb"
	"ops/pkg/model/rdb"
	utils2 "ops/pkg/utils"
	"ops/proto/userInfo"
	"strings"
)

type UserInfo struct {
	Uid            string `json:"uid,omitempty" bson:"uid,omitempty" binding:"required"`
	OpsAccount     string `json:"ops_account,omitempty" bson:"ops_account,omitempty"`
	NickName       string `json:"nick_name,omitempty" bson:"nick_name,omitempty"`
	Description    string `json:"description,omitempty" bson:"description,omitempty"`
	Link           string `json:"link,omitempty" bson:"link,omitempty"`
	RegisterTime   string `json:"register_time,omitempty" bson:"register_time,omitempty"`
	Profile        string `json:"profile,omitempty" bson:"profile,omitempty"`
	Banner         string `json:"banner,omitempty" bson:"banner,omitempty"`
	Location       string `json:"location,omitempty" bson:"location,omitempty"`
	IsOfficialUser bool   `json:"is_official_user,omitempty" bson:"is_official_user"`
	Label          string `json:"label,omitempty" bson:"label,omitempty"`
}

func (u *UserInfo) GetPbkAddrByUid(ctx context.Context, value *pbUserInfo.UidAndPbkaddr, result *pbUserInfo.UidAndPbkaddr) error {
	defer func() {
		logger.Infof("calling GetPbkAddrByUid success,  OpsAccount=%+v,  result=%+v", value, result)
	}()
	mgoClient := mgodb.NewMgo(utils2.GetConfigStr("mongodb.db"), utils2.GetConfigStr("mongodb.collection.user_id"))
	one := mgoClient.FindOne(bson.M{"uid": value.Uid})
	if one.Err() != nil {
		logger.Error(one.Err().Error())
		result.ErrCode = utils2.RECODE_DATAINEXISTENCE
		return nil
	}
	bytes, err := one.DecodeBytes()
	if err != nil {
		logger.Error(err)
		result.ErrCode = utils2.RECODE_DATAINEXISTENCE
		return nil
	}

	result.ErrCode = utils2.RECODE_OK
	result.Pubkaddr = bytes.Lookup("pubkaddr").StringValue()
	return nil
}

func (u *UserInfo) GetUidByOpsAccount(ctx context.Context, accounts *pbUserInfo.OpsAccounts, uid *pbUserInfo.OpsAccounts) error {
	defer func() {
		logger.Infof("calling GetUidByOpsAccount success,  OpsAccount=%+v,  result=%+v", accounts, uid)
	}()
	mgoClient := mgodb.NewMgo(utils2.GetConfigStr("mongodb.db"), utils2.GetConfigStr("mongodb.collection.user_detail"))
	findOptions := options.Find().SetProjection(bson.M{"uid": 1, "ops_account": 1, "_id": 0})
	many, err := mgoClient.FindManyRecords(bson.M{"ops_account": bson.M{"$in": accounts.OpsAccounts}}, findOptions)
	if err != nil {
		logger.Error(err)
		return nil
	}
	for many.Next(context.TODO()) {
		var opsAccountsAndUids = pbUserInfo.OpsAccountsAndUids{}
		opsAccountsAndUids.Uid = many.Current.Lookup("uid").StringValue()
		opsAccountsAndUids.OpsAccount = many.Current.Lookup("ops_account").StringValue()
		uid.OpsAndUid = append(uid.OpsAndUid, &opsAccountsAndUids)
	}
	return nil
}

func (u *UserInfo) StoreUserInfo(ctx context.Context, value *pbUserInfo.UserInfo, result *pbUserInfo.OperateResult) error {
	logger.Debug("Received UserService.StoreInfo request")
	mgoClient := mgodb.NewMgo(utils2.GetConfigStr("mongodb.db"), utils2.GetConfigStr("mongodb.collection.user_detail"))

	one := mgoClient.FindOne(bson.M{"ops_account": value.OpsAccount})
	if one.Err() == nil {
		result.Code = utils2.RECODE_USERONERR
		return nil
	}
	err := u.SetValue(value)
	if err != utils2.RECODE_OK {
		result.Code = err
	}
	err = mgoClient.InsertOne(u)
	if err != utils2.RECODE_OK {
		result.Code = err
	}
	result.Code = utils2.RECODE_OK
	return nil
}

func (u *UserInfo) UpdateUserInfo(ctx context.Context, value *pbUserInfo.UserInfo, result *pbUserInfo.UserInfo) error {
	logger.Debug("Received UserService.UpdateInfo request")
	mgoClient := mgodb.NewMgo(utils2.GetConfigStr("mongodb.db"), utils2.GetConfigStr("mongodb.collection.user_detail"))
	one := mgoClient.FindOne(bson.M{"ops_account": value.OpsAccount})
	if one.Err() == nil {
		result.Uid = utils2.RECODE_USERONERR
		return nil
	}
	err := u.SetValue(value)
	if err != utils2.RECODE_OK {
		result.Uid = err
	}
	updateResult := mgoClient.FindOneAndUpdate(bson.M{"uid": u.Uid}, bson.M{"$set": u})
	bytes, err2 := updateResult.DecodeBytes()
	if err2 != nil {
		return err2
	}

	err2 = jsonpb.UnmarshalString(bytes.String(), result)
	if err2 != nil {
		logger.Debug(err2)
	}
	return nil
}

func (u *UserInfo) DeleteUserInfo(ctx context.Context, value *pbUserInfo.QueryAndDelete, result *pbUserInfo.OperateResult) error {
	logger.Debug("Received UserService.DeleteUserInfo request")
	userInfo := mgodb.NewMgo(utils2.GetConfigStr("mongodb.db"), utils2.GetConfigStr("mongodb.collection.user_detail"))
	bytes, err := userInfo.FindOne(bson.M{"uid": value.Uid}).DecodeBytes()
	if err != nil {
		result.Code = utils2.RECODE_USERNOEXISTSERR
		return nil
	}
	profile := bytes.Lookup("profile").StringValue()
	banner := bytes.Lookup("banner").StringValue()
	if resultCode := deleteImg(profile); resultCode != utils2.RECODE_OK {
		result.Code = resultCode
		return nil
	} else if resultCode = deleteImg(banner); resultCode != utils2.RECODE_OK {
		result.Code = resultCode
		return nil
	} else if resultCode = userInfo.Delete(bson.M{"uid": value.Uid}); resultCode != utils2.RECODE_OK {
		result.Code = resultCode
		return nil
	} else {
		result.Code = resultCode
		return nil
	}
}

func (u *UserInfo) QueryUserInfo(ctx context.Context, value *pbUserInfo.QueryAndDelete, result *pbUserInfo.UserInfo) error {
	logger.Debug("Received UserService.QueryUserInfo request")
	userInfo := mgodb.NewMgo(utils2.GetConfigStr("mongodb.db"), utils2.GetConfigStr("mongodb.collection.user_detail"))
	singleResult := userInfo.FindOne(bson.M{"uid": value.Uid})
	if singleResult.Err() != nil {
		logger.Error(singleResult.Err())
		result.Uid = utils2.RECODE_USERNOEXISTSERR
		return nil
	}
	var userInfoOnMgo UserInfoOnMongo
	err := singleResult.Decode(&userInfoOnMgo)
	if err != nil {
		logger.Error(err)
		result.Uid = utils2.RECODE_DBERR
		return nil
	}
	marshal, err := json.Marshal(&userInfoOnMgo)
	if err != nil {
		logger.Error(err)
		result.Uid = utils2.RECODE_DBERR
		return nil
	}
	err = json.Unmarshal(marshal, &result)
	if err != nil {
		logger.Error(err)
		result.Uid = utils2.RECODE_DBERR
		return nil
	}
	return nil
}

func (u *UserInfo) IsRepeatedOpsAccount(ctx context.Context, account *pbUserInfo.CheckOpsAccount, result *pbUserInfo.OperateResult) error {
	mgoClient := mgodb.NewMgo(utils2.GetConfigStr("mongodb.db"), utils2.GetConfigStr("mongodb.collection.user_detail"))
	one := mgoClient.FindOne(bson.M{"ops_account": account.OpsAccount})
	//用户存在时，返回err是nil，返回用户已存在
	if one.Err() == nil {
		result.Code = utils2.RECODE_USERONERR
		return nil
	} else {
		result.Code = utils2.RECODE_OK
		return nil
	}
}

func (u *UserInfo) OneStepProtect(ctx context.Context, private *pbUserInfo.SetPrivate, result *pbUserInfo.OperateResult) error {
	mgoClient := mgodb.NewMgo(utils2.GetConfigStr("mongodb.db"), utils2.GetConfigStr("mongodb.collection.user_detail"))
	update := mgoClient.FindOneAndUpdate(bson.M{"uid": private.Uid}, bson.M{"$set": bson.M{"is_private": private.IsPrivate}})
	if update.Err() != nil {
		logger.Error(update.Err())
		return errors.New("store failed")
	}

	if private.IsPrivate {
		err := rdb.LPush("private_user", private.Uid)
		if err != nil {
			return errors.New("store private_user to rdb failed")
		}
	} else {
		err := rdb.LRem("private_user", 1, private.Uid)
		if err != utils2.RECODE_OK {
			return errors.New("store private_user to rdb failed")
		}
	}

	result.Code = utils2.RECODE_OK
	return nil
}

func (u *UserInfo) ModifyOopNum(ctx context.Context, value *pbUserInfo.OopNum, result *pbUserInfo.OperateResult) error {
	mgoClient := mgodb.NewMgo(utils2.GetConfigStr("mongodb.db"), utils2.GetConfigStr("mongodb.collection.user_detail"))
	update := mgoClient.UpsertOne(bson.M{"uid": value.Uid}, bson.M{"$inc": bson.M{"oop_number": value.Num}})
	result.Code = update
	return nil
}

func (u *UserInfo) SetPrice(ctx context.Context, value *pbUserInfo.Price, result *pbUserInfo.Price) error {
	mgoClient := mgodb.NewMgo(utils2.GetConfigStr("mongodb.db"), utils2.GetConfigStr("mongodb.collection.user_detail"))
	update := mgoClient.FindOneAndUpdate(bson.M{"uid": value.Uid}, bson.M{"$set": bson.M{"price": value.Price}})
	if update.Err() != nil {
		logger.Error(update.Err())
		result.Uid = utils2.RECODE_DBERR
		return nil
	}
	bytes, err := update.DecodeBytes()
	if err != nil {
		logger.Error(err)
		result.Uid = utils2.RECODE_DATAERR
		return nil
	}
	result.Uid = bytes.Lookup("uid").StringValue()
	result.Price = bytes.Lookup("price").Double()
	return nil
}

func (u *UserInfo) SetTotalReleaseOCardNum(ctx context.Context, value *pbUserInfo.TotalReleaseOCardNum, result *pbUserInfo.OperateResult) error {
	mgoClient := mgodb.NewMgo(utils2.GetConfigStr("mongodb.db"), utils2.GetConfigStr("mongodb.collection.user_detail"))
	logger.Info("uid: ", value.Uid, "total_amount: ", value.ReleaseCardNum)
	update := mgoClient.UpsertOne(bson.M{"uid": value.Uid}, bson.M{"$inc": bson.M{"total_amount": value.ReleaseCardNum}})
	if update != utils2.RECODE_OK {
		result.Code = utils2.RECODE_DBERR
		return nil
	}
	result.Code = update
	return nil
}

func (u *UserInfo) SetValue(info *pbUserInfo.UserInfo) string {
	u.Uid = info.Uid
	u.OpsAccount = info.OpsAccount
	u.NickName = info.NickName
	u.Description = info.Description
	u.Link = info.Link
	u.RegisterTime = info.RegisterTime
	u.Location = info.Location
	u.IsOfficialUser = info.IsOfficialUser
	u.Label = info.Label
	//u.IsPrivate = info.IsPrivate
	//upload new user img
	profile, err := uploadProfile(info.Uid, info.Profile)
	if err == utils2.RECODE_OK || err == utils2.RECODE_USEDEFAULTIMG {
		u.Profile = profile
	}
	banner, err := uploadBanner(info.Uid, info.Banner)
	if err == utils2.RECODE_OK || err == utils2.RECODE_USEDEFAULTIMG {
		u.Banner = banner
	}

	//delete old img
	if err = deleteImg(info.OldProfileLink); err != utils2.RECODE_OK {
		logger.Error("delete profile failed ", u.Uid)
	}
	if err = deleteImg(info.OldBannerLink); err != utils2.RECODE_OK {
		logger.Error("delete banner failed ", u.Uid)
	}
	return err
}

func uploadProfile(uid string, profile string) (string, string) {
	if len(profile) == 0 {
		return "", utils2.RECODE_USEDEFAULTIMG
	}

	toByte, err := base64ToByte(profile)
	if err == utils2.RECODE_DECODEIMGERR {
		logger.Error("Decode img from base64 to byte is failed.", uid)
		return "", err
	} else if len(toByte) >= 500*1024 {
		return "", utils2.RECODE_PROFILEERR
	}

	result, err := uploadImg(utils2.NewLen(20)+".png", toByte)
	if err != utils2.RECODE_OK {
		return "", err
	}

	return result, err
}

func uploadBanner(uid string, profile string) (string, string) {
	if len(profile) == 0 {
		return "", utils2.RECODE_USEDEFAULTIMG
	}

	toByte, err := base64ToByte(profile)
	if err == utils2.RECODE_DECODEIMGERR {
		logger.Error("Decode img from base64 to byte is failed.", uid)
		return "", err
	} else if len(toByte) >= 1000*1024 {
		return "", utils2.RECODE_BANNERERR
	}

	result, err := uploadImg(utils2.NewLen(20)+".png", toByte)
	if err != utils2.RECODE_OK {
		return "", err
	}
	return result, err
}

func deleteImg(img string) string {
	if img == "" {
		return utils2.RECODE_OK
	}
	blobName := strings.Split(img, "/")[3]
	blobURL := azureStorage.AzureStorage.Container.NewBlockBlobURL(blobName)
	_, err := blobURL.Delete(context.Background(), azblob.DeleteSnapshotsOptionNone, azblob.BlobAccessConditions{})
	if err != nil {
		logger.Error(err.Error())
		return utils2.RECODE_DELETEBLOBERR
	}
	return utils2.RECODE_OK
}

func uploadImg(opsAccount string, img []byte) (string, string) {
	const url = "https://storage.opsnft.net/"
	err := azureStorage.UploadFile(opsAccount, img)
	if err != utils2.RECODE_OK {
		logger.Error("Profile upload failed: ", opsAccount, err)
		return "", err
	}
	return url + opsAccount, err
}

func base64ToByte(imgBase64 string) ([]byte, string) {
	decodeString, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return nil, utils2.RECODE_DECODEIMGERR
	}
	return decodeString, utils2.RECODE_OK
}
