package oopResp

import (
	"ops/pkg/model"
	pbOop "ops/proto/oop"
	"encoding/json"
)

type OopResp struct {
	Oid          string        `json:"oid"`
	Uid          string        `json:"uid"`
	Content      string        `json:"content"`
	CreateTime   string        `json:"create_time"`
	ShardTimes   int32         `json:"shard_times"`
	CommentTimes int32         `json:"comment_times"`
	LikeTimes    int32         `json:"like_times"`
	IsPrivate    bool          `json:"is_private"`
	LikeDetail   []*pbOop.LikeDetail `json:"like_detail"`
	MyLikeInfo   *pbOop.LikeDetail   `json:"my_like_info"`
	AtUsers      string        `json:"at_users"`
}

type Resp struct {
	model.JSONResult
	Data *OopResp `json:"data"`
}

func (r *Resp) NewSuccess(data *pbOop.Oop) *Resp {
	r.JSONResult.NewSuccess()
	marshal, _ := json.Marshal(&data)
	realData := &OopResp{}
	_ = json.Unmarshal(marshal, &realData)
	r.Data = realData
	return r
}

type ManyResultResp struct {
	Code        string `json:"code"`
	ShowNumber  int64  `json:"show_number"`
	CurrentPage int64  `json:"current_page"`
	Timestamp   string `json:"timestamp"`
	Data        []*OopResp `json:"data"`
}

type QueryResp struct {
	model.JSONResult
	Data *ManyResultResp `json:"data"`
}

func (q *QueryResp) NewSuccess(data *pbOop.ManyResult) *QueryResp {
	q.JSONResult.NewSuccess()
	marshal, _ := json.Marshal(&data)
	realData := &ManyResultResp{}
	_ = json.Unmarshal(marshal, &realData)
	q.Data = realData
	return q
}

type MyLikeResp struct {
	model.JSONResult
	Data *ManyResultResp `json:"data"`
}

func (q *MyLikeResp) NewSuccess(data *pbOop.MyLikeOopResult) *MyLikeResp {
	q.JSONResult.NewSuccess()
	marshal, _ := json.Marshal(&data)
	realData := &ManyResultResp{}
	_ = json.Unmarshal(marshal, &realData)
	q.Data = realData
	return q
}
