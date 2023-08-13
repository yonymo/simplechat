package db

import (
	"context"
	"github.com/yonymo/simplechat/api/internal/data/friend/v1"
)

func (f *Friend) CreateFriendReq(ctx context.Context, friendReq *friend.FriendReqDO) error {
	return f.db.Create(friendReq).Error
}

func (f *Friend) GetFriendReq(ctx context.Context, owner, dst uint) (*friend.FriendReqDO, error) {
	do := &friend.FriendReqDO{}
	err := f.db.Where("owner_id = ? and friend_id = ?", owner, dst).Find(do).Error
	if err != nil {
		return nil, err
	}
	return do, nil
}

func (f *Friend) UpdateFriendReq(ctx context.Context, friendReq *friend.FriendReqDO) error {
	return f.db.Save(friendReq).Error
}
