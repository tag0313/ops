package userInfo

import (
	"encoding/json"
	"ops/pkg/model"
	myjwt "ops/pkg/model/jwt"
	utils2 "ops/pkg/utils"
	"ops/proto/userInfo"
)

type UpdateReq struct {
	Uid            string `json:"uid" swaggerignore:"true"`
	OpsAccount     string `json:"ops_account,omitempty"`
	NickName       string `json:"nick_name,omitempty"`
	Description    string `json:"description,omitempty"`
	Link           string `json:"link,omitempty"`
	Profile        string `json:"profile,omitempty"`
	Banner         string `json:"banner,omitempty"`
	OldProfileLink string `json:"old_profile_link,omitempty"`
	OldBannerLink  string `json:"old_banner_link,omitempty"`
	Location       string `json:"location,omitempty"`
	IsOfficialUser bool   `json:"is_official_user"`
	Label          string `json:"label"`
}

func (s *UpdateReq) SetValue(claims *myjwt.CustomClaims) *pbUserInfo.UserInfo {
	return &pbUserInfo.UserInfo{
		Uid:            claims.Uid,
		OpsAccount:     s.OpsAccount,
		NickName:       s.NickName,
		Description:    s.Description,
		Link:           s.Link,
		Profile:        s.Profile,
		Banner:         s.Banner,
		OldProfileLink: s.OldProfileLink,
		OldBannerLink:  s.OldBannerLink,
		Location:       s.Location,
		IsOfficialUser: s.IsOfficialUser,
		Label:          s.Label,
	}
}

type QuesryReq struct {
	Uid string `json:"uid"`
}

type RealPrice struct {
	Uid   string  `json:"uid"`
	Price float64 `json:"price"`
}

type PriceResp struct {
	model.JSONResult
	Data *RealPrice `json:"data"`
}

func (p *PriceResp) NewError(recode string) *PriceResp {
	p.Code = recode
	p.Message = utils2.RecodeTest(recode)
	p.Success = false
	return p
}

func (p *PriceResp) NewSuccess(data *pbUserInfo.Price) *PriceResp {
	p.Code = utils2.RECODE_OK
	p.Message = utils2.RecodeTest(utils2.RECODE_OK)
	p.Success = true
	marshal, _ := json.Marshal(&data)
	var realData *RealPrice
	_ = json.Unmarshal(marshal, &realData)
	p.Data = realData
	return p
}

type GetOpsAccountByUidResp struct {
	model.JSONResult
	Data []GetOpsAccountByUidData `json:"data"`
}

type GetOpsAccountByUidData struct {
	Uid        string `json:"uid"`
	OpsAccount string `json:"ops_account"`
}
