package biz

import (
	pb "bin-personal-book/api/user/v1"
	"bin-personal-book/internal/conf"
	"context"
	"errors"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

// UserRepo 方法
type UserRepo interface {
	GetUserAccount(ctx context.Context, params *pb.LoginParams) (*pb.LoginParams, error)
}

type UserUsecase struct {
	confData *conf.Data
	repo     UserRepo
	log      *log.Helper
}

func NewUserUseBiz(confData *conf.Data, repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{confData: confData, repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Login(ctx context.Context, params *pb.LoginParams) (*jwt.Token, error) {
	// 查找用户是否存在
	user, err := uc.repo.GetUserAccount(ctx, params)
	if err != nil {
		return nil, err
	}

	// 对比密码是否相同
	if user.Password != params.Password {
		return nil, errors.New("密码错误！")
	}

	// 传入指定的签名方法和payload信息,创建Token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"account":   user.Account,
		"ExpiresAt": uc.confData.Jwt.Expire,
	})

	tokenString, err := token.SignedString([]byte(uc.confData.Jwt.Secret))

	fmt.Println(tokenString)

	return token, nil
}
