package handler

type UserInfoOnMongo struct {
	Uid            string  `json:"uid,omitempty" bson:"uid,omitempty"`
	OpsAccount     string  `json:"ops_account,omitempty" bson:"ops_account,omitempty"`
	NickName       string  `json:"nick_name,omitempty" bson:"nick_name,omitempty"`
	Description    string  `json:"description,omitempty" bson:"description,omitempty"`
	Link           string  `json:"link,omitempty" bson:"link,omitempty"`
	RegisterTime   string  `json:"register_time,omitempty" bson:"register_time,omitempty"`
	Profile        string  `json:"profile,omitempty" bson:"profile,omitempty"`
	Banner         string  `json:"banner,omitempty" bson:"banner,omitempty"`
	OldProfileLink string  `json:"old_profile_link,omitempty" bson:"old_profile_link,omitempty"`
	OldBannerLink  string  `json:"old_banner_link,omitempty" bson:"old_banner_link,omitempty"`
	Location       string  `json:"location,omitempty" bson:"location,omitempty"`
	IsOfficialUser bool    `json:"is_official_user,omitempty" bson:"is_official_user,omitempty"`
	Label          string  `json:"label,omitempty" bson:"label,omitempty"`
	IsPrivate      bool    `json:"is_private,omitempty" bson:"is_private,omitempty"`
	OopNumber      int64   `json:"oop_number,omitempty" bson:"oop_number,omitempty"`
	TotalAmount    int64   `json:"total_amount,omitempty" bson:"total_amount,omitempty"`
	Price          float64 `json:"price,omitempty" bson:"price,omitempty"`
	Following      int64   `json:"following,omitempty" bson:"following,omitempty"`
	Followed       int64   `json:"followed,omitempty" bson:"followed,omitempty"`
}
