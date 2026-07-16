package biz

import (
	pb "bin-personal-book/api/user/v1"
	"bin-personal-book/internal/conf"
	"bin-personal-book/internal/core"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

// UserRepo 方法
type UserRepo interface {
	GetUserAccount(ctx context.Context, params *core.GetUserAccountParams) *pb.LoginParams
	InsertUserAccount(ctx context.Context, params *pb.RegisterParams) (*struct{}, error)
}

type UserUsecase struct {
	confData *conf.Data
	repo     UserRepo
	log      *log.Helper
}

func NewUserUseBiz(confData *conf.Data, repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{confData: confData, repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Login(ctx context.Context, params *pb.LoginParams) (*pb.LoginResult, error) {
	// 查找用户是否存在
	user := uc.repo.GetUserAccount(ctx, &core.GetUserAccountParams{
		Account: params.Account,
	})

	if user == nil {
		return nil, core.NewError("暂无该用户")
	}

	// 对比密码是否相同
	if user.Password != params.Password {
		return nil, core.NewError("密码错误")
	}

	// 传入指定的签名方法和payload信息,创建Token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"account":   user.Account,
		"ExpiresAt": uc.confData.Jwt.Expire,
	})

	// 转成string类型
	tokenString, err := token.SignedString([]byte(uc.confData.Jwt.Secret))

	if err != nil {
		return nil, err
	}

	return &pb.LoginResult{
		Token: tokenString,
	}, nil
}

func (uc *UserUsecase) Register(ctx context.Context, params *pb.RegisterParams) (*pb.RegisterResult, error) {
	// 查找用户是否存在
	user := uc.repo.GetUserAccount(ctx, &core.GetUserAccountParams{
		Account: params.Account,
	})

	if user != nil {
		return nil, core.NewError("该用户已注册")
	}

	_, InsertErr := uc.repo.InsertUserAccount(ctx, params)

	return &pb.RegisterResult{}, InsertErr
}
