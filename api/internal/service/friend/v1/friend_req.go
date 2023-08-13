package friend

import (
	"context"
	df "github.com/yonymo/simplechat/api/internal/data/friend/v1"
	"github.com/yonymo/simplechat/pkg/code"
	"github.com/yonymo/simplechat/pkg/errors"
	"github.com/yonymo/simplechat/pkg/log"
)

type FriendReqDTO struct {
	OwnerID   uint   `json:"owner_id" `
	FriendID  uint   `json:"friend_id" `
	Remark    string `json:"remark"`
	AddSource string `json:"add_source" `
	ReqText   string `json:"req_text,omitempty"`
	Extra     string `json:"extra" `
}

type FriendReqDTOList struct {
	Total int64           `json:"total"`
	Items []*FriendReqDTO `json:"items"`
}

func (f *friendService) AddFriendReq(ctx context.Context, dto *FriendReqDTO) error {
	reqDo, err := f.friendData.GetFriendReq(ctx, dto.OwnerID, dto.FriendID)
	if reqDo == nil {
		reqDo = &df.FriendReqDO{
			OwnerID:   dto.OwnerID,
			FriendID:  dto.FriendID,
			Remark:    dto.Remark,
			ReqText:   dto.ReqText,
			AddSource: dto.AddSource,
			Extra:     dto.Extra,
		}
		err = f.friendData.CreateFriendReq(ctx, reqDo)
		if err != nil {
			log.Debugf("create friend failed: %v", err)
			return errors.WithCode(code.ErrServerInternal, "创建好友申请失败")
		}
	} else {
		// 如果两人同时申请好友
		if reqDo.Remark == "" {
			// 这是被申请方
			reqDo = &df.FriendReqDO{
				OwnerID:   dto.OwnerID,
				FriendID:  dto.FriendID,
				Remark:    dto.Remark,
				ReqText:   dto.ReqText,
				AddSource: dto.AddSource,
				Extra:     dto.Extra,
			}
			err = f.friendData.UpdateFriendReq(ctx, reqDo)
			if err != nil {
				log.Debugf("update friend failed: %v, info: %v\n", err, reqDo)
				return errors.WithCode(code.ErrServerInternal, "创建好友申请失败")
			}
			return nil
		} else {
			return errors.WithCode(code.ErrFriendReqAlreadyCommit, "好友申请已提交")
		}
	}

	reqDo, err = f.friendData.GetFriendReq(ctx, dto.FriendID, dto.OwnerID)
	if reqDo == nil {
		reqDo = &df.FriendReqDO{
			OwnerID:   dto.FriendID,
			FriendID:  dto.OwnerID,
			AddSource: dto.AddSource,
		}
		err = f.friendData.CreateFriendReq(ctx, reqDo)
		if err != nil {
			log.Debugf("Second record create friend failed: %v", err)
			return errors.WithCode(code.ErrServerInternal, "创建好友申请失败")
		}
	}

	return nil
}
