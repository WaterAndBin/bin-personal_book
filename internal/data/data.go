package data

import (
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewMonodbClient, NewRedisClient, NewRedisData, NewUserData, NewTagsData, NewUploadData)
