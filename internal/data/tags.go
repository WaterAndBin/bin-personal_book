package data

import (
	tags "bin-personal-book/api/tags/v1"
	"time"

	"context"

	"bin-personal-book/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	list := make([]*tags.BillTagsInfo, 0)

	err := r.billTagsColl.Find(ctx, bson.M{}).All(&list)

	if err != nil {
		return nil, errors.BadRequest("error", err.Error())
	}

	return &tags.GetBillTagsListResult{
		List:   list,
		Length: int32(len(list)),
	}, nil
}

func (r *tagsData) UpdateBillTags(ctx context.Context, params *tags.UpdateBillTagsParams) (*tags.UpdateBillTagsResult, error) {
	if params.TagName == "" {
		return nil, errors.BadRequest("error", "缺少name")
	}

	if params.TagIcon == "" {
		return nil, errors.BadRequest("error", "缺少icon")
	}

	if params.TagId != "" {
		objectID, _ := primitive.ObjectIDFromHex(params.TagId)
		err := r.billTagsColl.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{
			"$set": bson.M{
				"tag_name":     params.TagName,
				"tag_icon":     params.TagIcon,
				"updated_time": time.Now().UTC(),
			},
		})

		if err != nil {
			return nil, errors.BadRequest("error", "更新标签失败")
		}

		return &tags.UpdateBillTagsResult{}, nil
	} else {
		_, err := r.billTagsColl.InsertOne(ctx, bson.M{
			"tag_id":       uuid.New().String(),
			"tag_name":     params.TagName,
			"tag_icon":     params.TagIcon,
			"updated_time": time.Now().UTC(),
			"created_time": time.Now().UTC(),
		},
		)

		if err != nil {
			return nil, errors.BadRequest("error", "创建标签失败")
		}

		return &tags.UpdateBillTagsResult{}, nil
	}
}
