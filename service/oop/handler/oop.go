package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/asim/go-micro/v3/util/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/encoding/protojson"
	"ops/pkg/logger"
	"ops/pkg/model/mgodb"
	"ops/pkg/model/rdb"
	"ops/pkg/utils"
	pbOop "ops/proto/oop"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Oop struct {
	Oid          string `json:"oid,omitempty" bson:"oid,omitempty"`
	Uid          string `json:"uid,omitempty" bson:"uid,omitempty"`
	Content      string `json:"content,omitempty" bson:"content,omitempty"`
	CreateTime   string `json:"create_time,omitempty" bson:"create_time,omitempty"`
	ShardTimes   int32  `json:"shard_times,omitempty" bson:"shard_times,omitempty"`
	CommentTimes int32  `json:"comment_times,omitempty" bson:"comment_times,omitempty"`
	LikeTimes    int32  `json:"like_times,omitempty" bson:"like_times,omitempty"`
	IsPrivate    bool   `json:"is_private" bson:"is_private"`
	AtUsers      string `json:"at_users,omitempty" bson:"at_users,omitempty"`
}

type Like struct {
	Oid        string `json:"oid,omitempty" bson:"oid,omitempty"`
	Uid        string `json:"uid,omitempty" bson:"uid,omitempty"`
	Star       int64  `json:"star,omitempty" bson:"star,omitempty"`
	CreateTime string `json:"create_time,omitempty" bson:"create_time,omitempty"`
}

func (o *Oop) QueryOwnerOop(ctx context.Context, in *pbOop.Query, result *pbOop.ManyResult) error {
	defer func() {
		logger.Infof("calling QueryOop success, request=%+v,  response=%+v", in, result)
	}()
	mgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db"), utils.GetConfigStr("mongodb.collection.oop"))
	var filter bson.M
	if in.Timestamp == "" {
		filter = bson.M{"uid": in.Uid}
	} else {
		filter = bson.M{"uid": in.Uid, "create_time": bson.M{"$lte": in.Timestamp}}
	}
	documents, errCode := mgoClient.CollectionDocuments(in.ShowNumber*in.PageNumber-in.ShowNumber,
		in.ShowNumber,
		bson.M{"create_time": -1},
		filter)
	if errCode == utils.RECODE_DATAERR {
		result.Code = errCode
		return nil
	} else {
		var oops []*Oop
		var oops2 []*pbOop.Oop
		err := documents.All(context.TODO(), &oops)
		if err != nil {
			logger.Error("binding data failed", err)
			return err
		}
		marshal, err := json.Marshal(&oops)
		if err != nil {
			return err
		}
		err = json.Unmarshal(marshal, &oops2)

		assembleLikeData(oops2, in.Uid)

		if err != nil {
			return err
		}
		result.Code = utils.RECODE_OK
		result.Data = oops2
		result.ShowNumber = in.ShowNumber
		result.CurrentPage = in.PageNumber
		result.Timestamp = in.Timestamp
		return nil
	}
}

func (o *Oop) QueryOtherOop(ctx context.Context, in *pbOop.Query, result *pbOop.ManyResult) error {
	defer func() {
		logger.Infof("calling QueryOop success, request=%+v,  response=%+v", in, result)
	}()
	mgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db"), utils.GetConfigStr("mongodb.collection.oop"))
	var filter bson.M
	if in.Timestamp == "" {
		filter = bson.M{"uid": in.Uid, "is_private": bson.M{"$ne": true}}
	} else {
		filter = bson.M{"uid": in.Uid, "is_private": bson.M{"$ne": true}, "create_time": bson.M{"$lte": in.Timestamp}}
	}
	documents, errCode := mgoClient.CollectionDocuments(in.ShowNumber*in.PageNumber-in.ShowNumber,
		in.ShowNumber,
		bson.M{"create_time": -1},
		filter)
	if errCode == utils.RECODE_DATAERR {
		result.Code = errCode
		return nil
	} else {
		var oops []*Oop
		var oops2 []*pbOop.Oop
		err := documents.All(context.TODO(), &oops)
		if err != nil {
			logger.Error("binding data failed", err)
			return err
		}
		marshal, err := json.Marshal(&oops)
		if err != nil {
			return err
		}
		err = json.Unmarshal(marshal, &oops2)
		if err != nil {
			return err
		}

		assembleLikeData(oops2, in.SelfUid)
		result.Code = utils.RECODE_OK
		result.Data = oops2
		result.ShowNumber = in.ShowNumber
		result.CurrentPage = in.PageNumber
		result.Timestamp = in.Timestamp
		return nil
	}
}

func (o *Oop) StoreOop(ctx context.Context, in *pbOop.Oop, result *pbOop.OneResult) error {
	defer func() {
		logger.Infof("calling StoreOop success, request=%+v,  response=%+v", in, result)
	}()
	if !isGtContentLimit(in.Content) {
		result.Code = utils.RECODE_OOP_CONTENT
		return nil
	}
	in.Oid = utils.NewLen(20)
	in.CreateTime = setCreateTime()
	mgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db"), utils.GetConfigStr("mongodb.collection.oop"))
	marshal, err := protojson.Marshal(in)
	if err != nil {
		logger.Error("proto marshal to json failed", err)
		return err
	}
	oopObj := new(Oop)
	err = json.Unmarshal(marshal, oopObj)
	if err != nil {
		logger.Error("json unmarshal failed", err)
		return err
	}
	logger.Info("存入数据库的数据为=%+v", oopObj)
	insertOne := mgoClient.InsertOne(oopObj)
	if insertOne == utils.RECODE_STOREDATA_FAILED {
		result.Code = insertOne
		return nil
	} else {
		result.Code = utils.RECODE_OK
		result.Data = in
		//发送通知信息
		logger.Info("oop中提及的人=%+v", in.AtUsers)
		if in.AtUsers != "" {
			mention := strings.Split(in.AtUsers, ",")
			for _, s := range mention {
				log.Info("通知被@用户=%+v", s)
				err = utils.SendMentionMsg(s, in.Uid, in.Oid, in.Content, in.AtUsers)
			}
			logger.Info("通知执行完成")
		}
		return nil
	}
}

func (o *Oop) UpdateOop(ctx context.Context, in *pbOop.Oop, result *pbOop.OneResult) error {
	if !isGtContentLimit(in.Content) {
		result.Code = utils.RECODE_OOP_CONTENT
		return nil
	}
	in.CreateTime = setCreateTime()
	mgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db"), utils.GetConfigStr("mongodb.collection.oop"))
	updateOne := mgoClient.FindOneAndUpdate(bson.M{"uid": in.GetUid(), "oid": in.GetOid()}, bson.M{"$set": bson.M{"content": in.GetContent(), "create_time": in.GetCreateTime(), "is_private": in.IsPrivate}})
	if updateOne.Err() != nil {
		logger.Error(updateOne.Err())
		result.Code = utils.RECODE_STOREDATA_FAILED
		return nil
	} else {
		result.Code = utils.RECODE_OK
		err := updateOne.Decode(o)
		if err != nil {
			logger.Error(err)
			return nil
		}
		var oopProto *pbOop.Oop
		marshal, err := json.Marshal(o)
		if err != nil || marshal == nil {
			logger.Error(err)
		}
		err = json.Unmarshal(marshal, &oopProto)
		if err != nil {
			logger.Error(err)
		}
		result.Data = oopProto
		//发送通知信息
		if in.AtUsers != "" {
			mention := strings.Split(in.AtUsers, ",")
			for _, s := range mention {
				err = utils.SendMentionMsg(s, in.Uid, in.Oid, in.Content, in.AtUsers)
			}
		}
		return nil
	}
}

func (o *Oop) DeleteOop(ctx context.Context, in *pbOop.Delete, result *pbOop.DeleteResult) error {
	mgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db"), utils.GetConfigStr("mongodb.collection.oop"))
	deleteOne := mgoClient.Delete(bson.M{"uid": in.GetUid(), "oid": in.GetOid()})
	if deleteOne == utils.RECODE_DATAINEXISTENCE {
		result.Code = deleteOne
		return nil
	} else {
		result.Code = utils.RECODE_OK
		return nil
	}
}

func (o *Oop) SquareOop(ctx context.Context, query *pbOop.Query, result *pbOop.ManyResult) error {
	privateList := rdb.GetList("private_user", 0, -1)
	if privateList == nil {
		privateList = make([]string, 0)
	}
	mgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db"), utils.GetConfigStr("mongodb.collection.oop"))
	var filter bson.M
	if query.Timestamp == "" {
		filter = bson.M{"is_private": bson.M{"$ne": true}, "uid": bson.M{"$nin": privateList}}
	} else {
		filter = bson.M{"is_private": bson.M{"$ne": true}, "uid": bson.M{"$nin": privateList}, "create_time": bson.M{"$lte": query.Timestamp}}
	}
	documents, errCode := mgoClient.CollectionDocuments(query.ShowNumber*query.PageNumber-query.ShowNumber,
		query.ShowNumber,
		bson.M{"create_time": -1},
		filter)
	if errCode == utils.RECODE_DATAERR {
		result.Code = errCode
		return nil
	} else {
		var oops []*Oop
		var oops2 []*pbOop.Oop
		err := documents.All(context.TODO(), &oops)
		if err != nil {
			logger.Error("binding data failed", err)
			return err
		}
		marshal, err := json.Marshal(&oops)
		if err != nil {
			return err
		}
		err = json.Unmarshal(marshal, &oops2)
		if err != nil {
			return err
		}
		result.Code = utils.RECODE_OK
		assembleLikeData(oops2, query.Uid)
		result.Data = oops2
		result.ShowNumber = query.ShowNumber
		result.CurrentPage = query.PageNumber
		result.Timestamp = query.Timestamp
		return nil
	}
}

func (o *Oop) QueryFollowOop(ctx context.Context, followOop *pbOop.FollowOop, result *pbOop.ManyResult) error {
	mgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db"), utils.GetConfigStr("mongodb.collection.oop"))
	var filter bson.M
	if followOop.Timestamp == "" {
		filter = bson.M{"uid": bson.M{"$in": followOop.Uid}}
	} else {
		filter = bson.M{"uid": bson.M{"$in": followOop.Uid}, "create_time": bson.M{"$lte": followOop.Timestamp}}
	}
	documents, errCode := mgoClient.CollectionDocuments(followOop.ShowNumber*followOop.PageNumber-followOop.ShowNumber,
		followOop.ShowNumber,
		bson.M{"create_time": -1},
		filter)
	if errCode == utils.RECODE_DATAERR {
		result.Code = errCode
		return nil
	} else {
		var oops []*Oop
		var oops2 []*pbOop.Oop
		err := documents.All(context.TODO(), &oops)
		if err != nil {
			logger.Error("binding data failed", err)
			return err
		}
		marshal, err := json.Marshal(&oops)
		if err != nil {
			return err
		}
		err = json.Unmarshal(marshal, &oops2)
		if err != nil {
			return err
		}
		result.Code = utils.RECODE_OK
		assembleLikeData(oops2, followOop.SelfUid)
		result.Data = oops2
		result.ShowNumber = followOop.ShowNumber
		result.CurrentPage = followOop.PageNumber
		result.Timestamp = followOop.Timestamp
		return nil
	}
}

func (o *Oop) LikeOop(ctx context.Context, like *pbOop.Like, result *pbOop.LikeOopResult) error {
	//验证参数正确性
	var starTypeArr = [3]int64{1, 2, 3}
	var starValid = false
	for _, value := range starTypeArr {
		if value == like.Star{
			starValid = true
		}
	}
	if !starValid {
		result.Code = utils.RECODE_LIKE_STAR_INVALID
		return nil
	}

	//查询oops是否存在，如果不存在直接返回错误提示
	mgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db"), utils.GetConfigStr("mongodb.collection.oop"))
	oopResult := mgoClient.FindOne(bson.M{"oid": like.Oid})
	if oopResult.Err() != nil {
		result.Code = utils.RECODE_DATAINEXISTENCE
		return nil
	}
	decodeErr := oopResult.Decode(o)
	if decodeErr != nil {
		logger.Error(decodeErr)
		return nil
	}

	//查询用户是否有点赞该oop，如果则修改点赞信息，如果未点赞则新增
	likeMgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db"), utils.GetConfigStr("mongodb.collection.like"))
	likeRecord := likeMgoClient.FindOne(bson.M{"oid": like.Oid, "uid": like.Uid})
	//组装entity
	likeEntity := &Like{}
	likeEntity.SetValue(like)
	err := utils.RECODE_OK
	if likeRecord.Err() != nil {
		err = likeMgoClient.InsertOne(likeEntity)
		//修改oop点赞数
		mgoClient.FindOneAndUpdate(bson.M{"oid": like.Oid}, bson.M{"$inc": bson.M{"like_times": 1}})
	} else {
		_, updateErr := likeMgoClient.FindOneAndUpdate(bson.M{"oid": like.Oid, "uid": like.Uid}, bson.M{"$set": likeEntity}).DecodeBytes()
		if updateErr != nil {
			err = updateErr.Error()
		}
	}
	result.Code = err
	//如果接口执行成功则发送通知
	if result.Code == utils.RECODE_OK{
		logger.Info("uid=%+v, oid=%+v, liker=%+v, star=%+v" , o.Uid, like.Oid, like.Uid, like.Star)
		_ = utils.SendLikeMsg(o.Uid, like.Oid, like.Uid, like.Star, o.Content)
	}
	return nil
}

func (o *Oop) MyLikeOop(ctx context.Context, like *pbOop.MyLike, result *pbOop.MyLikeOopResult) error {
	//从like表中先查询出所有点赞的oop ID
	likeMgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db"), utils.GetConfigStr("mongodb.collection.like"))
	documents, errCode := likeMgoClient.CollectionDocuments(like.ShowNumber*like.PageNumber-like.ShowNumber,
		like.ShowNumber,
		bson.M{"create_time": -1},
		bson.M{"uid": like.Uid})
	if errCode == utils.RECODE_DATAERR {
		result.Code = errCode
		fmt.Println(result.GetCode())
		return nil
	}
	var entityList []*Like
	err := documents.All(context.TODO(), &entityList)
	if err != nil {
		logger.Error("binding data failed", err)
		return err
	}
	length := len(entityList)
	//从oop表中查询出所有的oop列表信息
	oopIds := make([]string, length)
	likeMap := make(map[string]*pbOop.LikeDetail)
	for i, value := range entityList {
		oopIds[i] = value.Oid
		likeDetail := &pbOop.LikeDetail{}
		likeDetail.Star = value.Star
		likeMap[value.Oid] = likeDetail
	}
	mgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db"), utils.GetConfigStr("mongodb.collection.oop"))
	oopEntityList, oopErr := mgoClient.FindMany(bson.M{"oid": bson.M{"$in": oopIds}})
	if oopErr == utils.RECODE_DATAERR {
		result.Code = errCode
		return nil
	}
	var oops []*Oop
	var oops2 []*pbOop.Oop
	convertErr := oopEntityList.All(context.TODO(), &oops)
	if convertErr != nil {
		logger.Error("binding data failed", err)
		return err
	}
	marshal, err := json.Marshal(&oops)
	if err != nil {
		return err
	}
	err = json.Unmarshal(marshal, &oops2)
	if err != nil {
		return err
	}
	for _, value := range oops2 {
		value.MyLikeInfo = likeMap[value.Oid]
	}

	assembleLikeData(oops2, like.Uid)
	result.Code = utils.RECODE_OK
	result.Data = oops2
	result.ShowNumber = like.ShowNumber
	result.CurrentPage = like.PageNumber
	return nil

}

func (o *Oop) CancelLikeOop(ctx context.Context, like *pbOop.CancelLike, result *pbOop.CancelLikeResult) error {
	//查询用户是否有点赞该oop，如果则修改点赞信息，如果未点赞则新增
	likeMgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db"), utils.GetConfigStr("mongodb.collection.like"))
	likeRecord := likeMgoClient.FindOne(bson.M{"oid": like.Oid, "uid": like.Uid})
	if likeRecord.Err() != nil {
		result.Code = utils.RECODE_DATAINEXISTENCE
		return nil
	}
	//删除点赞记录
	deleteOne := likeMgoClient.Delete(bson.M{"uid": like.GetUid(), "oid": like.GetOid()})
	if deleteOne == utils.RECODE_DATAINEXISTENCE {
		result.Code = deleteOne
		return nil
	} else {
		result.Code = utils.RECODE_OK
		return nil
	}
}

func isGtContentLimit(content string) bool {
	var total int16

	//reg := regexp.MustCompile("/·|，|。|《|》|‘|’|”|“|；|：|【|】|？|（|）|、/")  || reg.Match([]byte(string(r)))
	for _, r := range content {
		if unicode.Is(unicode.Han, r) {
			total = total + 2
		} else {
			total = total + 1
		}
	}
	if total <= 280 && total != 0 {
		return true
	}
	return false
}

func setCreateTime() string {
	timestamp := time.Now().Unix()
	return strconv.FormatInt(timestamp, 10)
}

func (likeEntity *Like) SetValue(likeReq *pbOop.Like) {
	likeEntity.Uid = likeReq.Uid
	likeEntity.Oid = likeReq.Oid
	likeEntity.Star = likeReq.Star
	likeEntity.CreateTime = likeReq.CreateTime
}

func assembleLikeData(oops2 []*pbOop.Oop, uid string) {
	//查询点赞信息及当前用户点赞信息
	length := len(oops2)
	oopIds := make([]string, length)
	for i, v := range oops2 {
		oopIds[i] = v.Oid
	}
	likeMgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db"), utils.GetConfigStr("mongodb.collection.like"))
	findOptions := options.Find().SetSort(bson.M{"oid": -1})
	likes, likeErr := likeMgoClient.FindMany(bson.M{"oid": bson.M{"$in": oopIds}}, findOptions)
	if likeErr == utils.RECODE_OK {
		var likeEntityList []*Like
		err := likes.All(context.TODO(), &likeEntityList)
		if err != nil {
			logger.Error("binding data failed", err)
			return
		}
		//外层key为oopId，内层key为点赞类型value为点赞数量
		likeMap := make(map[string]map[int64]int64)
		for _, v := range oops2 {
			m := likeMap[v.Oid]
			if m == nil {
				likeStarMap := make(map[int64]int64)
				likeMap[v.Oid] = likeStarMap
			}
			assembling := false
			for _, like := range likeEntityList {
				if like.Oid == v.Oid {
					//如果为当前用户点赞则组装当前用户点赞信息
					if like.Uid == uid {
						myLikeInfo := &pbOop.LikeDetail{}
						myLikeInfo.Star = like.Star
						v.MyLikeInfo = myLikeInfo
					}
					likeCount := likeMap[v.Oid][like.Star]
					if likeCount == 0 {
						likeMap[v.Oid][like.Star] = like.Star
						assembling = true
					} else {
						likeMap[v.Oid][like.Star] = like.Star + likeCount
					}
				} else {
					if assembling {
						break
					}
				}
			}
			detailTypeLen := len(likeMap[v.Oid])
			likeDetails := make([]*pbOop.LikeDetail, detailTypeLen)
			index := 0
			for key, value := range likeMap[v.Oid] {
				likeDetail := &pbOop.LikeDetail{}
				likeDetail.Star = key
				likeDetail.Count = value
				likeDetails[index] = likeDetail
				v.LikeDetail = likeDetails
				index++
			}
		}
	}
}
