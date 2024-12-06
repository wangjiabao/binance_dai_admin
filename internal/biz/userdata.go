package biz

import (
	pb "binance_dai_admin/api/userdata/v1"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Userdata is a Userdata model.
type Userdata struct {
	Hello string
}

// UserdataRepo is a Greater repo.
type UserdataRepo interface {
	Save(context.Context, *Userdata) (*Userdata, error)
}

// UserdataUsecase is a Userdata usecase.
type UserdataUsecase struct {
	repo UserdataRepo
	tx   Transaction
	log  *log.Helper
}

// NewUserdataUsecase new a Userdata usecase.
func NewUserdataUsecase(repo UserdataRepo, tx Transaction, logger log.Logger) *UserdataUsecase {
	return &UserdataUsecase{repo: repo, tx: tx, log: log.NewHelper(logger)}
}

func (uc *UserdataUsecase) GetUserOrders(ctx context.Context, req *pb.GetUserOrdersRequest) (*pb.GetUserOrdersReply, error) {
	return &pb.GetUserOrdersReply{OrderId: 1}, nil
}
