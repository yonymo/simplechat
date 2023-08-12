package friend

import (
	"context"
	df "github.com/yonymo/simplechat/api/internal/data/friend/v1"
	"github.com/yonymo/simplechat/pkg/code"
	"github.com/yonymo/simplechat/pkg/errors"
	"github.com/yonymo/simplechat/pkg/log"
)

type FriendDTO struct {
	OwnerID   uint   `json:"owner_id" binding:"required"`
	FriendID  uint   `json:"friend_id" binding:"required"`
	Remark    string `json:"remark" binding:"required"`
	AddSource string `json:"add_source" binding:"required"`
	Extra     string `json:"extra" `
}

type FriendDTOList struct {
	Total int64        `json:"total"`
	Items []*FriendDTO `json:"items"`
}

type IFriendSrv interface {
	AddFriend(ctx context.Context, dto *FriendDTO) error
	GetFriendList(ctx context.Context, owner uint) (*FriendDTOList, error)
	GetFriend(ctx context.Context, owner, dst uint) (*FriendDTO, error)
}

type friendService struct {
	friendData df.IFriendData
}

func (f *friendService) GetFriend(ctx context.Context, owner, dst uint) (*FriendDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (f *friendService) AddFriend(ctx context.Context, dto *FriendDTO) error {
	friendDo, err := f.friendData.Get(ctx, dto.OwnerID, dto.FriendID)
	if friendDo == nil {
		friendDo = &df.FriendDO{
			OwnerID:   dto.OwnerID,
			FriendID:  dto.FriendID,
			Remark:    dto.Remark,
			Relation:  df.Relation_Normal,
			Black:     df.Not_Black_List,
			AddSource: dto.AddSource,
			Extra:     dto.Extra,
		}
		err = f.friendData.Create(ctx, friendDo)
		if err != nil {
			log.Debugf("create friend failed: %v", err)
			return errors.WithCode(code.ErrServerInternal, "创建失败")
		}

	} else {
		if friendDo.Relation == df.Relation_Normal {
			return errors.WithCode(code.ErrFriendAlreadyExist, "好友已存在")
		} else {
			friendDo.AddSource = dto.AddSource
			friendDo.Extra = dto.Extra
			friendDo.Remark = dto.Remark
			friendDo.Relation = df.Relation_Normal
			err = f.friendData.Update(ctx, friendDo)
			if err != nil {
				log.Debugf("update friend failed: %v", err)
				return errors.WithCode(code.ErrServerInternal, "修改好友状态失败")
			}

		}
	}

	friendDo, err = f.friendData.Get(ctx, dto.FriendID, dto.OwnerID)
	if friendDo == nil {
		friendDo = &df.FriendDO{
			OwnerID:   dto.FriendID,
			FriendID:  dto.OwnerID,
			Relation:  df.Relation_Normal,
			Black:     df.Not_Black_List,
			AddSource: dto.AddSource,
			Extra:     dto.Extra,
		}
		err = f.friendData.Create(ctx, friendDo)
		if err != nil {
			log.Debugf("create friend failed: %v info: %v", err, *friendDo)
			return errors.WithCode(code.ErrServerInternal, "创建失败")
		}
	}
	return nil
}

func (f *friendService) GetFriendList(ctx context.Context, owner uint) (*FriendDTOList, error) {
	//TODO implement me
	panic("implement me")
}

var _ IFriendSrv = &friendService{}

func NewFriendSrv(ifd df.IFriendData) IFriendSrv {
	return &friendService{friendData: ifd}
}
