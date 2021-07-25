package message

import (
	"ops/pkg/model"
	pbMessage "ops/proto/message"
)

type GetNumMessageReq struct {
	Num int64  `json:"num" binding:"required"`
	Uid string `swaggerignore:"true"`
}

func (g GetNumMessageReq) SetValue() *pbMessage.Num {
	return &pbMessage.Num{
		Num: g.Num,
		Uid: g.Uid,
	}
}

type MessageResp struct {
	model.JSONResult
	Data []interface{} `json:"data"`
}

func (m *MessageResp) NewSuccess(value *pbMessage.Msgs) *MessageResp {
	m.JSONResult.NewSuccess()
	for _, data := range value.Data {
		switch data.MsgType {
		case 1:
			m.Data = append(m.Data, data.Txn)
		case 2:
			m.Data = append(m.Data, data.Follow)
		case 3:
			m.Data = append(m.Data, data.Like)
		case 4:
			m.Data = append(m.Data, data.Foword)
		case 5:
			m.Data = append(m.Data, data.Comment)
		case 6:
			m.Data = append(m.Data, data.PrivateMessage)
		case 7:
			m.Data = append(m.Data, data.OcardNotEnough)
		case 8:
			m.Data = append(m.Data, data.CreateCard)
		}
	}
	return m
}

type PushMessageResult struct {
	model.JSONResult
}
