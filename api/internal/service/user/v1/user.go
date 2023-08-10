package user

import (
	"context"
	du "github.com/yonymo/simplechat/api/internal/data/user/v1"
	"github.com/yonymo/simplechat/pkg/common"
)

type UserDTO struct {
	*du.UserDO
}

type UserDTOList struct {
	Total int64      `json:"total"`
	Items []*UserDTO `json:"items"`
}

type IUserSrv interface {
	MobileLogin(ctx context.Context, mobile, password string) (*UserDTO, error)
	Register(ctx context.Context, mobile, password, codes string) (*UserDTO, error)
	Update(ctx context.Context, dto *UserDTO) error
	Get(ctx context.Context, userId uint) (*UserDTO, error)
	GetByMobile(ctx context.Context, mobile string) (*UserDTO, error)
	CheckPassword(ctx context.Context, password, EncryptedPassword string) (bool, error)
}

type userService struct {
	userData du.IUserData
}

func (u *userService) MobileLogin(ctx context.Context, mobile, password string) (*UserDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userService) Register(ctx context.Context, mobile, password, codes string) (*UserDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userService) Get(ctx context.Context, userId uint) (*UserDTO, error) {
	do, err := u.userData.GetById(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &UserDTO{do}, nil
}

func (u *userService) CheckPassword(ctx context.Context, password, EncryptedPassword string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userService) List(ctx context.Context, opts common.ListMeta, orderby []string) (*UserDTOList, error) {
	udl, err := u.userData.List(ctx, opts, orderby)
	if err != nil {
		return nil, err
	}
	dtoList := &UserDTOList{Total: udl.Total}
	for _, val := range udl.Items {
		dtoList.Items = append(dtoList.Items, &UserDTO{val})
	}
	return dtoList, nil
}

func (u *userService) Update(ctx context.Context, user *UserDTO) error {
	return u.userData.Update(ctx, user.UserDO)
}

func (u *userService) GetByMobile(ctx context.Context, mobile string) (*UserDTO, error) {
	do, err := u.userData.GetByMobile(ctx, mobile)
	if err != nil {
		return nil, err
	}

	return &UserDTO{do}, nil
}

var _ IUserSrv = &userService{}

func NewUserSrv(iu du.IUserData) IUserSrv {
	return &userService{iu}
}
