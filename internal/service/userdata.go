package service

import (
	"binance_dai_admin/internal/biz"
	"context"

	pb "binance_dai_admin/api/userdata/v1"
)

type UserdataService struct {
	pb.UnimplementedUserdataServer
	uc *biz.UserdataUsecase
}

func NewUserdataService(uc *biz.UserdataUsecase) *UserdataService {
	return &UserdataService{
		uc: uc,
	}
}

func (s *UserdataService) GetUserOrders(ctx context.Context, req *pb.GetUserOrdersRequest) (*pb.GetUserOrdersReply, error) {
	return s.uc.GetUserOrders(ctx, req)
}

func (s *UserdataService) PullUserIncome(ctx context.Context, req *pb.PullUserIncomeRequest) (*pb.PullUserIncomeReply, error) {
	return s.uc.PullUserIncome(ctx, req)
}

func (s *UserdataService) GetUsersIncome(ctx context.Context, req *pb.GetUsersIncomeRequest) (*pb.GetUsersIncomeReply, error) {
	return s.uc.GetUsersIncome(ctx, req)
}
