package core

type CoreBillTag struct {
	TagId   string `bson:"tag_id" json:"tag_id"`
	TagName string `bson:"tag_name" json:"tag_name"`
	TagIcon string `bson:"tag_icon" json:"tag_icon"`
}
