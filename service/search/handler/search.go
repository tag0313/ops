package handler

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ops/pkg/logger"
	"ops/pkg/model/mgodb"
	"ops/pkg/model/rdb"
	utils2 "ops/pkg/utils"
	"ops/proto/search"
	"strconv"
	"time"
)

type Search struct {
	Oid          string `json:"oid,omitempty" bson:"oid,omitempty"`
	Uid          string `json:"uid,omitempty" bson:"uid,omitempty"`
	Content      string `json:"content,omitempty" bson:"content,omitempty"`
	CreateTime   string `json:"create_time,omitempty" bson:"create_time,omitempty"`
	LikeTimes    int64  `json:"like_times,omitempty" bson:"like_times,omitempty"`
	CommentTimes int64  `json:"comment_times,omitempty" bson:"comment_times,omitempty"`
	ShardTimes   int64  `json:"shard_times,omitempty" bson:"shard_times,omitempty"`
}

type UserInfo struct {
	Uid          string `json:"uid,omitempty" bson:"uid,omitempty"`
	OpsAccount   string `json:"ops_account,omitempty" bson:"ops_account,omitempty"`
	NickName     string `json:"nick_name,omitempty" bson:"nick_name,omitempty"`
	RegisterTime string `json:"register_time,omitempty" bson:"register_time,omitempty"`
	Profile      string `json:"profile,omitempty" bson:"profile,omitempty"`
	Banner       string `json:"banner,omitempty" bson:"banner,omitempty"`
	Description  string `json:"description,omitempty" bson:"description,omitempty"`
	Link         string `json:"link,omitempty" bson:"link,omitempty"`
	Location     string `json:"location,omitempty" bson:"location,omitempty"`
}

func (s Search) SearchID(ctx context.Context, userID *pbSearch.UserID, SearchContentResults *pbSearch.SearchContentResults) error {
	logger.Debug("Received recieve.Service.SearchID request")
	privateList := rdb.GetList("private_user", 0, -1)
	if privateList == nil {
		privateList = make([]string, 0)
	}
	// 如果uid在redis列表里面
	for _, v := range privateList {
		if v == userID.Uid {
			return nil
		}
	}
	mgoClientOop := mgodb.NewMgo(utils2.GetConfigStr("mongodb.db.ops"), utils2.GetConfigStr("mongodb.collection.oop"))
	filter := bson.M{"uid": userID.Uid, "is_private": bson.M{"$ne": true}}
	if userID.Timestamp != "" {
		filter["create_time"] = bson.M{"$lte": userID.Timestamp}
	}
	oops, errString := mgoClientOop.CollectionDocuments(userID.ShowNumber*userID.PageNumber-userID.ShowNumber,
		userID.ShowNumber,
		bson.M{"create_time": -1},
		filter)
	if errString == utils2.RECODE_DATAINEXISTENCE {
		logger.Error("RECODE_DATAINEXISTENCE")
		return nil
	}
	var oopArray []*Search
	var returnOops []*pbSearch.SearchContentResult
	err := oops.All(context.TODO(), &oopArray)
	if err != nil {
		logger.Error("binding data failed", err)
		return err
	}
	marshal, err := json.Marshal(&oopArray)
	if err != nil {
		logger.Error("Marshall data failed", err)
		return err
	}
	err = json.Unmarshal(marshal, &returnOops)
	if err != nil {
		logger.Error("Unmarshall data failed", err)
		return err
	}
	SearchContentResults.Data = returnOops
	SearchContentResults.Timestamp = userID.Timestamp
	SearchContentResults.CurrentPage = userID.PageNumber
	SearchContentResults.ShowNumber = userID.ShowNumber
	return nil
}

func (s Search) SearchContent(ctx context.Context, content *pbSearch.Content, SearchContentResults *pbSearch.SearchContentResults) error {
	logger.Debug("Received recieve.Service.SearchContent request")
	privateList := rdb.GetList("private_user", 0, -1)
	if privateList == nil {
		privateList = make([]string, 0)
	}
	mgoClient := mgodb.NewMgo(utils2.GetConfigStr("mongodb.db.ops"), utils2.GetConfigStr("mongodb.collection.oop"))
	filter := bson.M{"content": primitive.Regex{Pattern: content.Content, Options: "i"}, "is_private": bson.M{"$ne": true}, "uid": bson.M{"$nin": privateList}}
	if content.Timestamp != "" {
		filter["create_time"] = bson.M{"$lte": content.Timestamp}
	}
	oops, errString := mgoClient.CollectionDocuments(content.ShowNumber*content.PageNumber-content.ShowNumber,
		content.ShowNumber,
		bson.M{"create_time": -1},
		filter)
	if errString == utils2.RECODE_DATAINEXISTENCE {
		logger.Error("RECODE_DATAINEXISTENCE")
		return nil
	}
	var oopArray []*Search
	var returnOops []*pbSearch.SearchContentResult
	err := oops.All(context.TODO(), &oopArray)
	if err != nil {
		logger.Error("binding data failed", err)
		return err
	}
	marshal, err := json.Marshal(&oopArray)
	if err != nil {
		logger.Error("Marshall data failed", err)
		return err
	}
	err = json.Unmarshal(marshal, &returnOops)
	if err != nil {
		logger.Error("Unmarshall data failed", err)
		return err
	}
	SearchContentResults.Data = returnOops
	SearchContentResults.Timestamp = content.Timestamp
	SearchContentResults.CurrentPage = content.PageNumber
	SearchContentResults.ShowNumber = content.ShowNumber
	return nil
}

func (s Search) SearchUser(ctx context.Context, content *pbSearch.Content, searchUserResults *pbSearch.SearchUserResults) error {
	mgoClient := mgodb.NewMgo(utils2.GetConfigStr("mongodb.db.user_info"), utils2.GetConfigStr("mongodb.collection.user_detail"))
	filterBoth := bson.M{"ops_account": primitive.Regex{Pattern: content.Content, Options: "i"}, "nick_name": primitive.Regex{Pattern: content.Content, Options: "i"}}
	userDetails, errString := mgoClient.CollectionDocuments(content.ShowNumber*content.PageNumber-content.ShowNumber,
		content.ShowNumber,
		nil,
		filterBoth)
	var userInfoArray []*UserInfo
	err := userDetails.All(context.TODO(), &userInfoArray)
	if err != nil {
		logger.Error("第一次绑定数据失败", err)
		return err
	}
	if errString == utils2.RECODE_DATAINEXISTENCE || len(userInfoArray) < int(content.ShowNumber) {
		var pageNumber int64
		if result, _ := getBothPageNumKey(content.Content, content.ShowNumber); result == "" {
			setBothPageNumKey(content.Content, content.PageNumber, content.ShowNumber)
			pageNumber = 1
		} else {
			PageUpNumber, _ := strconv.ParseInt(result, 10, 64)
			if content.PageNumber > 1 {
				pageNumber = content.PageNumber - PageUpNumber + 1
			} else {
				pageNumber = 1
			}
		}
		filterOpsAccount := bson.M{"ops_account": primitive.Regex{Pattern: content.Content, Options: "i"}, "nick_name": bson.M{"$not": primitive.Regex{Pattern: content.Content, Options: "i"}}}
		showNumber := content.ShowNumber - int64(len(userInfoArray))
		userDetails, errString = mgoClient.CollectionDocuments(showNumber*pageNumber-showNumber,
			showNumber,
			nil,
			filterOpsAccount)
		var userInfoArrayOpsAccount []*UserInfo
		err := userDetails.All(context.TODO(), &userInfoArrayOpsAccount)
		if err != nil {
			logger.Error("第二次绑定数据失败", err)
			return err
		}
		userInfoArray = append(userInfoArray, userInfoArrayOpsAccount...)
		if errString == utils2.RECODE_DATAINEXISTENCE || len(userInfoArray) < int(content.ShowNumber) {
			if result, _ := getOpsAccountPageNumKey(content.Content, content.ShowNumber); result == "" {
				setOpsAccountPageNumKey(content.Content, content.PageNumber, content.ShowNumber)
				pageNumber = 1
			} else {
				PageUpNumber, _ := strconv.ParseInt(result, 10, 64)
				if content.PageNumber > 1 {
					pageNumber = content.PageNumber - PageUpNumber + 1
				} else {
					pageNumber = 1
				}
			}
			filterNickName := bson.M{"nick_name": primitive.Regex{Pattern: content.Content, Options: "i"}, "ops_account": bson.M{"$not": primitive.Regex{Pattern: content.Content, Options: "i"}}}
			showNumber = content.ShowNumber - int64(len(userInfoArray))
			userDetails, _ = mgoClient.CollectionDocuments(showNumber*pageNumber-showNumber,
				showNumber,
				nil,
				filterNickName)
			var userInfoArrayNickName []*UserInfo
			err = userDetails.All(context.TODO(), &userInfoArrayNickName)
			if err != nil {
				logger.Error("第三次绑定数据失败", err)
				return err
			}
			userInfoArray = append(userInfoArray, userInfoArrayNickName...)
		}
	}

	// 满足ShowNumber条记录 或者记录不够了
	str, err := json.Marshal(userInfoArray)
	if err != nil {
		logger.Error("marshal失败", err)
		return err
	}
	var returnUserDetails []*pbSearch.SearchUserResult
	err = json.Unmarshal(str, &returnUserDetails)
	if err != nil {
		logger.Error("Unmarshal失败", err)
		return err
	}
	searchUserResults.Data = returnUserDetails
	searchUserResults.CurrentPage = content.PageNumber
	searchUserResults.ShowNumber = content.ShowNumber
	return nil
}

func setBothPageNumKey(keyWord string, page, showNumber int64) string {
	err := rdb.SetS("pageUpBoth_"+keyWord+"_"+strconv.FormatInt(showNumber, 10), page, 5*time.Minute)
	return err
}

func getBothPageNumKey(keyWord string, showNumber int64) (string, string) {
	result, err := rdb.Get("pageUpBoth_" + keyWord + "_" + strconv.FormatInt(showNumber, 10))
	return result, err
}

func setOpsAccountPageNumKey(keyWord string, page, showNumber int64) string {
	err := rdb.SetS("pageUpOpsAccount_"+keyWord+"_"+strconv.FormatInt(showNumber, 10), page, 5*time.Minute)
	return err
}

func getOpsAccountPageNumKey(keyWord string, showNumber int64) (string, string) {
	result, err := rdb.Get("pageUpOpsAccount_" + keyWord + "_" + strconv.FormatInt(showNumber, 10))
	return result, err
}
