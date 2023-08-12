package db

import (
	"context"
	"github.com/yonymo/simplechat/api/internal/data/friend/v1"
	"github.com/yonymo/simplechat/pkg/common"
	"gorm.io/gorm"
)

type Friend struct {
	db *gorm.DB
}

func (f *Friend) Get(ctx context.Context, owner, dst uint) (*friend.FriendDO, error) {
	do := &friend.FriendDO{}
	err := f.db.Where("owner_id = ? and friend_id = ?", owner, dst).Find(do).Error
	if err != nil {
		return nil, err
	}
	return do, nil
}

func (f *Friend) List(ctx context.Context, uid uint, opts common.ListMeta, orderby []string) (*friend.FriendDOList, error) {
	//TODO implement me
	panic("implement me")
}

func (f *Friend) Create(ctx context.Context, friend *friend.FriendDO) error {

	return f.db.Create(friend).Error
}

func (f *Friend) Update(ctx context.Context, friend *friend.FriendDO) error {
	return f.db.Save(friend).Error
}

var _ friend.IFriendData = &Friend{}

func NewFriendData(db *gorm.DB) friend.IFriendData {
	return &Friend{db: db}
}
