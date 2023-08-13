package db

import (
	"context"
	"errors"
	"github.com/yonymo/simplechat/api/internal/data/friend/v1"
	"github.com/yonymo/simplechat/pkg/common"
)

func (f *Friend) Get(ctx context.Context, owner, dst uint) (*friend.FriendDO, error) {
	do := &friend.FriendDO{}
	tx := f.db.Where("owner_id = ? and friend_id = ?", owner, dst).Find(do)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("not found record.")
	}
	return do, nil
}

func (f *Friend) List(ctx context.Context, uid uint, opts common.ListMeta) (*friend.FriendDOList, error) {
	lists := []*friend.FriendDO{}
	query := f.db
	offset := int((opts.Pn - 1) * opts.PSize)
	limit := int(opts.PSize)
	tx := query.Where("owner_id = ?", uid).Offset(offset).Limit(limit).Find(&lists)
	if tx.Error != nil {
		return nil, tx.Error
	}
	doList := &friend.FriendDOList{}
	doList.Total = tx.RowsAffected
	doList.Items = lists
	return doList, nil
}

func (f *Friend) Create(ctx context.Context, friend *friend.FriendDO) error {

	return f.db.Create(friend).Error
}

func (f *Friend) Update(ctx context.Context, friend *friend.FriendDO) error {
	return f.db.Save(friend).Error
}
