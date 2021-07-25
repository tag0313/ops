package userInfo

import (
	"encoding/json"
	"ops/pkg/model"
	myjwt "ops/pkg/model/jwt"
	utils2 "ops/pkg/utils"
	"ops/proto/userInfo"
	"strconv"
	"time"
)

type StoreReq struct {
	Uid            string `json:"uid" swaggerignore:"true"`
	OpsAccount     string `json:"ops_account,omitempty"`
	NickName       string `json:"nick_name,omitempty" binding:"required"`
	Description    string `json:"description,omitempty"`
	Link           string `json:"link,omitempty"`
	RegisterTime   string `json:"register_time,omitempty" swaggerignore:"true"`
	Profile        string `json:"profile,omitempty"`
	Banner         string `json:"banner,omitempty"`
	OldProfileLink string `json:"old_profile_link,omitempty"`
	OldBannerLink  string `json:"old_banner_link,omitempty"`
	Location       string `json:"location,omitempty"`
	IsOfficialUser bool   `json:"is_official_user"`
	Label          string `json:"label"`
	IsPrivate      bool   `json:"is_private"`
}

type UserInfoResp struct {
	Uid            string  `json:"uid"`
	OpsAccount     string  `json:"ops_account"`
	NickName       string  `json:"nick_name"`
	Description    string  `json:"description"`
	Link           string  `json:"link"`
	RegisterTime   string  `json:"register_time"`
	Profile        string  `json:"profile"`
	Banner         string  `json:"banner"`
	OldProfileLink string  `json:"old_profile_link"`
	OldBannerLink  string  `json:"old_banner_link"`
	Location       string  `json:"location"`
	IsOfficialUser bool    `json:"is_official_user"`
	Label          string  `json:"label"`
	IsPrivate      bool    `json:"is_private"`
	OopNumber      int64   `json:"oop_number"`
	TotalAmount    int64   `json:"total_amount"`
	Price          float64 `json:"price"`
	Following      int64   `json:"following"`
	Followed       int64   `json:"followed"`
}

type UserResp struct {
	model.JSONResult
	Data *UserInfoResp `json:"data"`
}

func (s *StoreReq) SetValue(claims *myjwt.CustomClaims) *pbUserInfo.UserInfo {
	s.RegisterTime = getRegisterTime()
	switch {
	case s.OpsAccount == "":
		s.OpsAccount = utils2.NewLen(12)
	}
	return &pbUserInfo.UserInfo{
		Uid:            claims.Uid,
		OpsAccount:     s.OpsAccount,
		NickName:       s.NickName,
		Description:    s.Description,
		Link:           s.Link,
		RegisterTime:   s.RegisterTime,
		Profile:        s.Profile,
		Banner:         s.Banner,
		OldProfileLink: s.OldProfileLink,
		OldBannerLink:  s.OldBannerLink,
		Location:       s.Location,
		IsOfficialUser: s.IsOfficialUser,
		Label:          s.Label,
		IsPrivate:      s.IsPrivate,
	}
}

func getRegisterTime() string {
	timestamp := time.Now().Unix()
	return strconv.FormatInt(timestamp, 10)
}

func (u *UserResp) NewSuccess(value *pbUserInfo.UserInfo) *UserResp {
	u.JSONResult.NewSuccess()
	marshal, _ := json.Marshal(&value)
	var realData *UserInfoResp
	_ = json.Unmarshal(marshal, &realData)
	u.Data = realData
	return u
}
