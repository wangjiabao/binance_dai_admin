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

func (s *UserdataService) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersReply, error) {
	return s.uc.GetUsers(ctx, req)
}

func (s *UserdataService) GetUsersIncome(ctx context.Context, req *pb.GetUsersIncomeRequest) (*pb.GetUsersIncomeReply, error) {
	return s.uc.GetUsersIncome(ctx, req)
}

func (s *UserdataService) GetNum(ctx context.Context, req *pb.GetNumRequest) (*pb.GetNumReply, error) {
	return s.uc.GetNum(ctx, req)
}

func (s *UserdataService) UpdateUserNum(ctx context.Context, req *pb.UpdateUserNumRequest) (*pb.UpdateUserNumReply, error) {
	return s.uc.UpdateUserNum(ctx, req)
}

func (s *UserdataService) UpdateNum(ctx context.Context, req *pb.UpdateNumRequest) (*pb.UpdateNumReply, error) {
	return s.uc.UpdateNum(ctx, req)
}
