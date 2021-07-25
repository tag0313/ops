package handler

import (
	"context"
	"github.com/asim/go-micro/v3/logger"
	"github.com/golang/protobuf/jsonpb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"ops/pkg/model/mgodb"
	pbFollower "ops/proto/follower"
	"ops/pkg/utils"
	"strconv"
	"time"
)

type Follower struct {
}

func (f Follower) WhoFollowedMe(ctx context.Context, req *pbFollower.RelationReq, resp *pbFollower.RelationResp) error {
	defer func() {
		logger.Infof("calling WhoFollowedMe,  card=%+v,  result=%+v", req, resp)
	}()
	follow := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.user_info"), utils.GetConfigStr("mongodb.collection.following"))
	many, errCode := follow.FindMany(bson.M{"following": req.Uid, "uid": bson.M{"$in": req.Relationship}})
	if errCode == utils.RECODE_DATAINEXISTENCE {
		resp.Code = errCode
		return nil
	}
	var currentUid string
	resultList := make([]string, many.RemainingBatchLength())
	for i := 0; many.Next(context.TODO()); i++ {
		currentUid = many.Current.Lookup("uid").StringValue()
		resultList[i] = currentUid
	}
	resp.Code = utils.RECODE_OK
	resp.Uid = resultList
	return nil
}

func (f Follower) WhoFollowingMe(ctx context.Context, req *pbFollower.RelationReq, resp *pbFollower.RelationResp) error {
	defer func() {
		logger.Infof("calling WhoFollowingMe,  card=%+v,  result=%+v", req, resp)
	}()
	follow := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.user_info"), utils.GetConfigStr("mongodb.collection.following"))
	many, errCode := follow.FindMany(bson.M{"uid": req.Uid, "following": bson.M{"$in": req.Relationship}})
	if errCode == utils.RECODE_DATAINEXISTENCE {
		resp.Code = errCode
		return nil
	}
	var currentUid string
	resultList := make([]string, many.RemainingBatchLength())
	for i := 0; many.Next(context.TODO()); i++ {
		currentUid = many.Current.Lookup("following").StringValue()
		resultList[i] = currentUid
	}
	resp.Code = utils.RECODE_OK
	resp.Uid = resultList
	return nil
}

func (f Follower) Follow(ctx context.Context, follower *pbFollower.Follower, num *pbFollower.FollowNum) error {
	follow := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.user_info"), utils.GetConfigStr("mongodb.collection.following"))

	//查询当前用户是否已关注该账户，如果已关注则无需再重复关注
	originData := follow.FindOne(bson.M{"uid": follower.Uid, "following": follower.Following})
	if originData.Err() == nil{
		num.Uid = utils.RECODE_REPEAT_FOLLOW
		return nil
	}

	userDetail := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.user_info"), utils.GetConfigStr("mongodb.collection.user_detail"))
	//查询被关注者是否存在
	own := userDetail.FindOne(bson.M{"uid": follower.Following})
	if own.Err() != nil{
		num.Uid = utils.RECODE_FOLLOWING_NOT_EXIST
		return nil
	}

	one := follow.InsertOne(bson.M{"uid": follower.Uid, "following": follower.Following, "create_time": getRegisterTime()})
	update := userDetail.FindOneAndUpdate(bson.M{"uid": follower.Uid}, bson.M{"$inc": bson.M{"following": 1}})  //更新并返回关注后的关注数
	updateOne := userDetail.UpsertOne(bson.M{"uid": follower.Following}, bson.M{"$inc": bson.M{"followed": 1}}) //更新被关注人的关注数，不返回

	if one == utils.RECODE_STOREDATA_FAILED || update.Err() != nil || updateOne == utils.RECODE_STOREDATA_FAILED {
		num.Uid = utils.RECODE_STOREDATA_FAILED
	}
	return nil
}

func (f Follower) CanalFollow(ctx context.Context, follower *pbFollower.Follower, num *pbFollower.FollowNum) error {
	follow := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.user_info"), utils.GetConfigStr("mongodb.collection.following"))
	userDetail := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.user_info"), utils.GetConfigStr("mongodb.collection.user_detail"))
	follow.Delete(bson.M{"uid": follower.Uid, "following": follower.Following})
	update := userDetail.FindOneAndUpdate(bson.M{"uid": follower.Uid}, bson.M{"$inc": bson.M{"following": -1}}) //更新并返回去取消关注后的关注数
	userDetail.UpsertOne(bson.M{"uid": follower.Following}, bson.M{"$inc": bson.M{"followed": -1}})             //更新被关注人的关注数，不返回

	bytes, err := update.DecodeBytes()
	if err != nil {
		return err
	}
	jsonpb.UnmarshalString(bytes.String(), num)
	return nil
}

func (f Follower) QueryFollowListAll(ctx context.Context, req *pbFollower.OopFollowListReq, resp *pbFollower.OopFollowListResp) error {
	followClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.user_info"), utils.GetConfigStr("mongodb.collection.following"))
	followOptions := options.Find()
	followOptions.SetProjection(bson.M{"_id": 0}).SetSort(bson.M{"create_time": -1})
	var filter bson.M
	if req.Timestamp == "" {
		filter = bson.M{"uid": req.Uid}
	} else {
		filter = bson.M{"uid": req.Uid, "create_time": bson.M{"$lte": req.Timestamp}}
	}
	many, errCode := followClient.FindMany(filter, followOptions)
	if errCode == utils.RECODE_DATAINEXISTENCE {
		resp.Code = errCode
		return nil
	}

	uidList := make([]string, many.RemainingBatchLength())
	for i := 0; many.Next(context.TODO()); i++ {
		uidList[i] = many.Current.Lookup("following").StringValue()
	}
	resp.Code = utils.RECODE_OK
	resp.Uid = uidList
	return nil
}

func (f Follower) QueryFollowingList(ctx context.Context, req *pbFollower.FollowListReq, resp *pbFollower.FollowListResp) error {
	followClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.user_info"), utils.GetConfigStr("mongodb.collection.following"))
	followOptions := options.Find()
	followOptions.SetProjection(bson.M{"_id": 0})

	var filter bson.M
	if req.Timestamp == "" {
		filter = bson.M{"uid": req.Uid}
	} else {
		filter = bson.M{"uid": req.Uid, "create_time": bson.M{"$lte": req.Timestamp}}
	}
	documents, errCode := followClient.CollectionDocuments(req.ShowNumber*req.PageNumber-req.ShowNumber,
		req.ShowNumber,
		bson.M{"creat_time": -1},
		filter,
	)
	if errCode == utils.RECODE_DATAERR {
		resp.Code = errCode
		return nil
	}

	followingList := make([]string, 0)
	for documents.Next(context.Background()) {
		followingList = append(followingList, documents.Current.Lookup("following").StringValue())
	}

	resp.Code = utils.RECODE_OK
	resp.Uid = RemoveRepeatedElement(followingList)
	return nil
}

func (f Follower) QueryFollowedList(ctx context.Context, req *pbFollower.FollowListReq, resp *pbFollower.FollowListResp) error {
	followClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.user_info"), utils.GetConfigStr("mongodb.collection.following"))
	followOptions := options.Find()
	followOptions.SetProjection(bson.M{"_id": 0})

	var filter bson.M
	if req.Timestamp == "" {
		filter = bson.M{"following": req.Uid}
	} else {
		filter = bson.M{"following": req.Uid, "create_time": bson.M{"$lte": req.Timestamp}}
	}
	documents, errCode := followClient.CollectionDocuments(req.ShowNumber*req.PageNumber-req.ShowNumber,
		req.ShowNumber,
		bson.M{"creat_time": -1},
		filter)
	if errCode == utils.RECODE_DATAERR {
		resp.Code = errCode
		return nil
	}

	followedList := make([]string, 0)
	for documents.Next(context.Background()) {
		followedList = append(followedList, documents.Current.Lookup("uid").StringValue())
	}

	resp.Code = utils.RECODE_OK
	resp.Uid = RemoveRepeatedElement(followedList)
	return nil
}

func getRegisterTime() string {
	timestamp := time.Now().Unix()
	return strconv.FormatInt(timestamp, 10)
}

func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
