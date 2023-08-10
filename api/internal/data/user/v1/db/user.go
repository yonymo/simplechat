package db

import (
	"context"
	"errors"
	"fmt"
	user2 "github.com/yonymo/simplechat/api/internal/data/user/v1"
	"github.com/yonymo/simplechat/pkg/common"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func (u *User) List(ctx context.Context, opts common.ListMeta, orderby []string) (*user2.UserDOList, error) {

	limit := opts.PSize
	offset := opts.Pn
	if limit == 0 {
		limit = 10
	}
	if offset > 0 {
		offset = (offset - 1) * limit
	}
	query := u.db
	for _, od := range orderby {
		query = query.Order(od)
	}
	ret := &user2.UserDOList{}
	tx := query.Offset(int(offset)).Limit(int(limit)).Find(&ret.Items).Count(&ret.Total)
	if tx.Error != nil {
		return nil, errors.New(fmt.Sprintf("database error: ", tx.Error.Error()))
	}
	return ret, nil
}

func (u *User) Create(ctx context.Context, user *user2.UserDO) error {
	return u.db.Create(user).Error
}

func (u *User) Update(ctx context.Context, user *user2.UserDO) error {
	return u.db.Save(user).Error
}

func (u *User) GetByMobile(ctx context.Context, mobile string) (*user2.UserDO, error) {
	ud := &user2.UserDO{}
	tx := u.db.Where("mobile = ?", mobile).First(ud)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return ud, nil
}

func (u *User) GetById(ctx context.Context, id uint) (*user2.UserDO, error) {
	ud := &user2.UserDO{}
	tx := u.db.First(ud, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return ud, nil
}

var _ user2.IUserData = &User{}

func NewUserData(db *gorm.DB) user2.IUserData {
	return &User{db: db}
}
