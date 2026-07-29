package data

import (
	"bin-personal-book/internal/conf"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/qiniu/qmgo"
	qmgoOptions "github.com/qiniu/qmgo/options"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
)

type MongoData struct {
	db *qmgo.Database
}

// NewData .
func NewMonodbClient(c *conf.Data, logger log.Logger) (*MongoData, func(), error) {
	// 创建一个10秒的超时控制
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := qmgo.NewClient(
		ctx,
		&qmgo.Config{Uri: "mongodb://localhost:27017"},
		qmgoOptions.ClientOptions{
			ClientOptions: mongoOptions.Client().SetBSONOptions(&mongoOptions.BSONOptions{
				UseJSONStructTags: true,
			}),
		},
	)

	if err != nil {
		log.NewHelper(logger).Fatal("数据库连接失败")
		return nil, nil, err
	}

	cleanup := func() {
		log.NewHelper(logger).Info("断开数据库连接")
		err := client.Close(ctx)
		if err != nil {
			return
		}
	}

	db := client.Database(c.Mongodb.Database)

	log.NewHelper(logger).Info("数据库连接成功")

	return &MongoData{db: db}, cleanup, nil
}
