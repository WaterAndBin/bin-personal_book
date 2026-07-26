package data

import (
	tags "bin-personal-book/api/tags/v1"

	"context"

	"bin-personal-book/internal/biz"
	"bin-personal-book/internal/core"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type tagsData struct {
	data         *Data
	log          *log.Helper
	billTagsColl *qmgo.Collection
}

func NewTagsData(data *Data, logger log.Logger) biz.TagsZip {
	return &tagsData{
		data:         data,
		log:          log.NewHelper(logger),
		billTagsColl: data.db.Collection("bill_tags"),
	}
}

func (r *tagsData) GetBillTagsList(ctx context.Context, params *tags.GetBillTagsListParams) (*tags.GetBillTagsListResult, error) {
	rows := make([]core.CoreBillTag, 0)

	err := r.billTagsColl.Find(ctx, bson.M{}).All(&rows)

	if err != nil {
		return nil, errors.BadRequest("error", err.Error())
	}

	list := make([]*tags.BillTagsInfo, 0, len(rows))
	for _, item := range rows {
		list = append(list, &tags.BillTagsInfo{
			TagId:   item.TagId,
			TagName: item.TagName,
			TagIcon: item.TagIcon,
		})
	}

	return &tags.GetBillTagsListResult{
		List:   list,
		Length: int32(len(list)),
	}, nil
}

func (r *tagsData) UpdateBillTags(ctx context.Context, params *tags.BillTagsInfo) (*tags.UpdateBillTagsResult, error) {
	if params.TagName == "" {
		return nil, errors.BadRequest("error", "缺少name")
	}

	if params.TagIcon == "" {
		return nil, errors.BadRequest("error", "缺少icon")
	}

	_, err := r.billTagsColl.InsertOne(ctx, core.CoreBillTag{
		TagId:   uuid.New().String(),
		TagName: params.TagName,
		TagIcon: params.TagIcon,
	})

	if err != nil {
		return nil, errors.BadRequest("error", "创建标签失败")
	}

	return &tags.UpdateBillTagsResult{}, nil
}
