package v1

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BillTagsInfo struct {
	TagId       primitive.ObjectID `bson:"_id" json:"tag_id"`
	TagName     string             `bson:"tag_name" json:"tag_name"`
	TagIcon     string             `bson:"tag_icon" json:"tag_icon"`
	CreatedTime time.Time          `bson:"created_time" json:"created_time"`
	UpdatedTime time.Time          `bson:"updated_time" json:"updated_time"`
}

type GetBillTagsListParams struct {
}

type GetBillTagsListResult struct {
	List   []*BillTagsInfo `json:"list"`
	Length int32           `json:"length"`
}
