package friend

const (
	Friend_Req_Result_Approve uint8 = 1
	Friend_Req_Already_Read   uint8 = 1
)

type FriendReqDO struct {
	BaseModel
	OwnerID    uint   `json:"owner_id" gorm:"index:idx_owner;type:uint"`
	FriendID   uint   `json:"friend_id" gorm:"index:idx_friend;type:uint"`
	ReadStatus uint8  `json:"read_status" gorm:"type:tinyint;default:0;comment:'1已读'"`
	Result     uint8  `json:"result" gorm:"type:tinyint;default:0;comment:'1同意'"`
	Remark     string `json:"remark" gorm:"type:varchar(100); comment:'备注'"`
	ReqText    string `json:"req_text" gorm:"type:varchar(255); comment:'申请内容'"`
	AddSource  string `json:"add_source" gorm:"type:varchar(20); comment:'来源'"`
	Extra      string `json:"extra" gorm:"type:varchar(1000)"`
}

func (u *FriendReqDO) TableName() string {
	return "friend_req"
}

type FriendReqDOList struct {
	Total int64          `json:"total,omitempty"`
	Items []*FriendReqDO `json:"items"`
}
