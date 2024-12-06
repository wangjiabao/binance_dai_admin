package data

import (
	"context"

	"binance_dai_admin/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UserdataRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserdataRepo .
func NewUserdataRepo(data *Data, logger log.Logger) biz.UserdataRepo {
	return &UserdataRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u *UserdataRepo) Save(ctx context.Context, ud *biz.Userdata) (*biz.Userdata, error) {
	return ud, nil
}
