package data

import (
	tags "bin-personal-book/api/tags/v1"

	"context"

	"bin-personal-book/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type tagsData struct {
	data         *Data
	log          *log.Helper
	billTagsColl *mongo.Collection
}

func NewTagsData(data *Data, logger log.Logger) biz.TagsZip {
	return &tagsData{
		data:         data,
		log:          log.NewHelper(logger),
		billTagsColl: data.db.Collection("bill_tags"),
	}
}

func (r *tagsData) GetBillTagsList(ctx context.Context, params *tags.GetBillTagsListParams) (*tags.GetBillTagsListResult, error) {
	cursor, err := r.billTagsColl.Find(ctx, bson.M{})

	if err != nil {
		return nil, errors.BadRequest("error", err.Error())
	}

	list := make([]*tags.GetBillTagsInfo, 0)

	// 把 Cursor（查询结果游标）里面的所有数据，一次性读取出来，并转换成 Go 的切片。
	err = cursor.All(ctx, &list)
	if err != nil {
		return nil, errors.BadRequest("error", err.Error())
	}

	return &tags.GetBillTagsListResult{
		List:   list,
		Length: int32(len(list)),
	}, nil
}
