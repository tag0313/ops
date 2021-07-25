package follower

import (
	"ops/pkg/model"
	"ops/pkg/model/oopResp"
	"ops/pkg/model/userInfo"
	"ops/proto/follower"
	pbOop "ops/proto/oop"
	"ops/proto/userInfo"
	"encoding/json"
)

type FollowNumResp struct {
	Uid       string `json:"uid"`
	Following string `json:"following"`
	Followed  string `json:"followed"`
}

type Resp struct {
	model.JSONResult
	Data *FollowNumResp `json:"data"`
}

func (r *Resp) NewSuccess(data *pbFollower.FollowNum) *Resp {
	r.JSONResult.NewSuccess()
	marshal, _ := json.Marshal(&data)
	realData := &FollowNumResp{}
	_ = json.Unmarshal(marshal, &realData)
	r.Data = realData
	return r
}

type OopReq struct {
	Uid        string `json:"uid"`
	Timestamp  string `json:"timestamp"`
	PageNumber int64  `json:"page_number"`
	ShowNumber int64  `json:"show_number"`
}

type OopResp struct {
	model.JSONResult
	Data *oopResp.ManyResultResp `json:"data"`
}

func (o *OopResp) NewSuccess(data *pbOop.ManyResult) *OopResp {
	o.JSONResult.NewSuccess()
	marshal, _ := json.Marshal(&data)
	realData := &oopResp.ManyResultResp{}
	_ = json.Unmarshal(marshal, &realData)
	o.Data = realData
	return o
}



type FollowListResp struct {
	model.JSONResult
	Data []*userInfo.UserInfoResp `json:"data"`
}

func (ing *FollowListResp) NewSuccess(data []*pbUserInfo.UserInfo) *FollowListResp {
	ing.JSONResult.NewSuccess()
	marshal, _ := json.Marshal(&data)
	var realData []*userInfo.UserInfoResp
	_ = json.Unmarshal(marshal, &realData)
	ing.Data = realData
	return ing
}

type RelationResp struct {
	model.JSONResult
	Data []string `json:"data"`
}

func (ing *RelationResp) NewSuccess(data []string) *RelationResp {
	ing.JSONResult.NewSuccess()
	ing.Data = data
	return ing
}
