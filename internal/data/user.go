package data

import (
	pb "bin-personal-book/api/user/v1"
	"context"

	"bin-personal-book/internal/biz"
	"bin-personal-book/internal/core"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type userData struct {
	data     *Data
	log      *log.Helper
	userColl *mongo.Collection
}

func NewUserData(data *Data, logger log.Logger) biz.UserZip {
	return &userData{
		data:     data,
		log:      log.NewHelper(logger),
		userColl: data.db.Collection("user"),
	}
}

func (r *userData) GetUserAccount(ctx context.Context, g *core.GetUserAccountParams) *pb.LoginParams {
	user := &pb.LoginParams{}

	err := r.userColl.FindOne(ctx, bson.M{
		"account": g.Account,
	}).Decode(user)

	if err != nil {
		return nil
	}

	return user
}

func (r *userData) InsertUserAccount(ctx context.Context, g *pb.RegisterParams) (*struct{}, error) {
	_, err := r.userColl.InsertOne(ctx, bson.M{
		"account":  g.Account,
		"password": g.Password,
	},
	)

	if err != nil {
		return nil, errors.BadRequest("error", "暂无该用户")
	}

	return nil, nil
}
