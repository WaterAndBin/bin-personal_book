package data

import (
	"bin-personal-book/internal/conf"
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/qiniu/qmgo"
)

type Data struct {
	db *qmgo.Database
}

// NewData .
func NewMonodb(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	// 创建一个10秒的超时控制
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := qmgo.NewClient(
		ctx,
		&qmgo.Config{Uri: "mongodb://localhost:27017"},
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

	fmt.Println(c.Mongodb.Database)

	db := client.Database(c.Mongodb.Database)

	return &Data{db: db}, cleanup, nil
}
