package data

import (
	pb "bin-personal-book/api/user/v1"
	"context"
	"errors"

	"bin-personal-book/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type userRepo struct {
	data     *Data
	log      *log.Helper
	userColl *mongo.Collection
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data:     data,
		log:      log.NewHelper(logger),
		userColl: data.db.Collection("user"),
	}
}

func (r *userRepo) GetUserAccount(ctx context.Context, g *pb.LoginParams) (*pb.LoginParams, error) {
	user := &pb.LoginParams{}

	err := r.userColl.FindOne(ctx, bson.M{
		"account": g.Account,
	}).Decode(user)

	if err != nil {
		return nil, errors.New("查找不到用户信息")
	}

	return user, nil
}
