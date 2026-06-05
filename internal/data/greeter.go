package data

import (
	"bin-personal-book/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
