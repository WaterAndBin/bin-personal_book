package data

import (
	"bin-personal-book/internal/conf"
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var ProviderSet = wire.NewSet(NewMonodb)

type Data struct {
	db *mongo.Database
}

// NewData .
func NewMonodb(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	// 创建一个10秒的超时控制
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(
		options.Client().
			ApplyURI("mongodb://127.0.0.1:27017"),
	)

	if err != nil {
		log.NewHelper(logger).Fatal("数据库连接失败")
		return nil, nil, err
	}

	// 检查连接是否成功
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.NewHelper(logger).Fatal("数据库连接失败")
		return nil, nil, err
	} else {
		fmt.Println("===数据库连接成功===")
	}

	cleanup := func() {
		log.NewHelper(logger).Info("断开数据库连接")
		err := client.Disconnect(ctx)
		if err != nil {
			return
		}
	}

	fmt.Println(c.Mongodb.Database)

	db := client.Database(c.Mongodb.Database)

	return &Data{db: db}, cleanup, nil
}
