package user

import (
	"context"
	"github.com/golang-jwt/jwt/v4"

	//"github.com/dgrijalva/jwt-go"

	"github.com/yonymo/simplechat/pkg/code"
	"github.com/yonymo/simplechat/pkg/common/pwd"
	"github.com/yonymo/simplechat/pkg/errors"
	"github.com/yonymo/simplechat/pkg/log"
	"time"

	du "github.com/yonymo/simplechat/api/internal/data/user/v1"
	"github.com/yonymo/simplechat/pkg/common"
	mjwt "github.com/yonymo/simplechat/pkg/middleware/jwt"
	"github.com/yonymo/simplechat/pkg/options"
)

type UserDTO struct {
	*du.UserDO
	Expire int64
}

type UserDTOList struct {
	Total int64      `json:"total"`
	Items []*UserDTO `json:"items"`
}

type IUserSrv interface {
	MobileLogin(ctx context.Context, mobile, password string) (*UserDTO, error)
	Register(ctx context.Context, mobile, password string) (*UserDTO, error)
	Update(ctx context.Context, dto *UserDTO) error
	Get(ctx context.Context, userId uint) (*UserDTO, error)
	GetByMobile(ctx context.Context, mobile string) (*UserDTO, error)
	CheckPassword(ctx context.Context, password, EncryptedPassword string) (bool, error)
}

type userService struct {
	userData du.IUserData
	jwtOps   *options.JwtOptions
}

func (u *userService) MobileLogin(ctx context.Context, mobile, password string) (*UserDTO, error) {
	userDTO, err := u.GetByMobile(ctx, mobile)
	if err != nil {
		return nil, err
	}
	if ok := pwd.VerifyPwd(password, userDTO.Passwd); !ok {
		log.Debugf("user passwd incorrect: %s == %s\n", password, userDTO.Passwd)
		return nil, errors.WithCode(code.ErrUserPasswordIncorrect, "密码错误")
	}
	jwtObj := mjwt.NewJWT(u.jwtOps.Key)
	tNow := time.Now()
	expire := time.Now().Add(u.jwtOps.Timeout).Local()
	claims := mjwt.CustomClaims{
		ID:          userDTO.ID,
		NickName:    userDTO.Nickname,
		AuthorityId: userDTO.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    u.jwtOps.Realm,
			Subject:   "chat",
			ID:        "1",
			Audience:  []string{"chat"},
			NotBefore: jwt.NewNumericDate(tNow),
			ExpiresAt: jwt.NewNumericDate(expire),
			IssuedAt:  jwt.NewNumericDate(tNow),
		},
	}

	userDTO.Token, err = jwtObj.CreateToken(claims)
	if err != nil {
		log.Debugf("user token create failed: %v claims: %v\n", err, claims)
		return nil, errors.WithCode(code.ErrTokenCreate, "token create failed")
	}
	userDTO.Expire = expire.Unix()
	return userDTO, nil
}

func (u *userService) Register(ctx context.Context, mobile, passwd string) (*UserDTO, error) {
	udo := &du.UserDO{
		Mobile:   mobile,
		Passwd:   passwd,
		Nickname: mobile,
	}

	//密码加密
	udo.Passwd = pwd.EncodePwd(passwd)

	err := u.userData.Create(ctx, udo)
	if err != nil {
		return nil, err
	}

	jwtObj := mjwt.NewJWT(u.jwtOps.Key)
	tNow := time.Now()
	expire := time.Now().Add(u.jwtOps.Timeout).Local()
	claims := mjwt.CustomClaims{
		ID:          udo.ID,
		NickName:    mobile,
		AuthorityId: udo.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    u.jwtOps.Realm,
			Subject:   "chat",
			ID:        "1",
			Audience:  []string{"chat"},
			NotBefore: jwt.NewNumericDate(tNow),
			ExpiresAt: jwt.NewNumericDate(expire),
			IssuedAt:  jwt.NewNumericDate(tNow),
		},
	}

	udo.Token, err = jwtObj.CreateToken(claims)
	if err != nil {
		return nil, err
	}
	return &UserDTO{UserDO: udo, Expire: expire.Unix()}, nil
}

func (u *userService) Get(ctx context.Context, userId uint) (*UserDTO, error) {
	do, err := u.userData.GetById(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &UserDTO{UserDO: do}, nil
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
		dtoList.Items = append(dtoList.Items, &UserDTO{UserDO: val})
	}
	return dtoList, nil
}

func (u *userService) Update(ctx context.Context, user *UserDTO) error {
	return u.userData.Update(ctx, user.UserDO)
}

func (u *userService) GetByMobile(ctx context.Context, mobile string) (*UserDTO, error) {
	do, err := u.userData.GetByMobile(ctx, mobile)
	if err != nil {
		log.Debugf("user not found: %v\n", err)
		return nil, errors.WithCode(code.ErrUserNotFound, "用户未注册")
	}

	return &UserDTO{UserDO: do}, nil
}

var _ IUserSrv = &userService{}

func NewUserSrv(iu du.IUserData, jwtOps *options.JwtOptions) IUserSrv {
	return &userService{userData: iu, jwtOps: jwtOps}
}
