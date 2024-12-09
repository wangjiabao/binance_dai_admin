package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
	"time"

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

type User struct {
	ID                  uint64    `gorm:"primarykey;type:int"`
	Address             string    `gorm:"type:varchar(100);not null"`
	ApiStatus           uint64    `gorm:"type:int;not null"`
	ApiKey              string    `gorm:"type:varchar(200);not null"`
	ApiSecret           string    `gorm:"type:varchar(200);not null"`
	BindTraderStatus    uint64    `gorm:"type:int;not null"`
	BindTraderStatusTfi uint64    `gorm:"type:int;not null"`
	IsDai               uint64    `gorm:"type:int;not null"`
	UseNewSystem        uint64    `gorm:"type:int;not null"`
	CreatedAt           time.Time `gorm:"type:datetime;not null"`
	UpdatedAt           time.Time `gorm:"type:datetime;not null"`
	BinanceId           uint64    `gorm:"type:int;not null"`
	OkxId               string    `gorm:"type:varchar(200);not null"`
	NeedInit            uint64    `gorm:"type:int;not null"`
	Dai                 uint64    `gorm:"type:int;not null"`
	Plat                string    `gorm:"type:varchar(200);not null"`
	OkxPass             string    `gorm:"type:varchar(200);not null"`
	Num                 float64   `gorm:"type:decimal(65,20);not null"`
}

// GetUsers .
func (u *UserdataRepo) GetUsers(ctx context.Context) ([]*biz.User, error) {
	var users []*User
	if err := u.data.DB(ctx).Table("new_user").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ERROR", err.Error())
	}

	res := make([]*biz.User, 0)
	for _, v := range users {
		res = append(res, &biz.User{
			ID:                  v.ID,
			Address:             v.Address,
			ApiStatus:           v.ApiStatus,
			ApiKey:              v.ApiKey,
			ApiSecret:           v.ApiSecret,
			BindTraderStatus:    v.BindTraderStatus,
			BindTraderStatusTfi: v.BindTraderStatusTfi,
			IsDai:               v.Dai,
			UseNewSystem:        v.UseNewSystem,
			CreatedAt:           v.CreatedAt,
			UpdatedAt:           v.UpdatedAt,
			BinanceId:           v.BinanceId,
			OkxId:               v.OkxId,
			NeedInit:            v.NeedInit,
			Dai:                 v.Dai,
			Plat:                v.Plat,
			OkxPass:             v.OkxPass,
			Num:                 v.Num,
		})
	}

	return res, nil
}

// GetUsersStatusOk .
func (u *UserdataRepo) GetUsersStatusOk(ctx context.Context) ([]*biz.User, error) {
	var users []*User
	if err := u.data.DB(ctx).Table("new_user").Where("api_status=?", 1).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ERROR", err.Error())
	}

	res := make([]*biz.User, 0)
	for _, v := range users {
		res = append(res, &biz.User{
			ID:                  v.ID,
			Address:             v.Address,
			ApiStatus:           v.ApiStatus,
			ApiKey:              v.ApiKey,
			ApiSecret:           v.ApiSecret,
			BindTraderStatus:    v.BindTraderStatus,
			BindTraderStatusTfi: v.BindTraderStatusTfi,
			IsDai:               v.Dai,
			UseNewSystem:        v.UseNewSystem,
			CreatedAt:           v.CreatedAt,
			UpdatedAt:           v.UpdatedAt,
			BinanceId:           v.BinanceId,
			OkxId:               v.OkxId,
			NeedInit:            v.NeedInit,
			Dai:                 v.Dai,
			Plat:                v.Plat,
			OkxPass:             v.OkxPass,
			Num:                 v.Num,
		})
	}

	return res, nil
}

type UserIncomeBinance struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null"`
	Symbol    string    `gorm:"type:varchar(45);not null"`
	Income    string    `gorm:"type:varchar(45);not null"`
	Info      string    `gorm:"type:varchar(45);not null"`
	TradeId   string    `gorm:"type:varchar(45);not null"`
	Time      uint64    `gorm:"type:bigint(20);not null"`
	TranId    string    `gorm:"type:varchar(45);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

// GetUserIncomesBinance .
func (u *UserdataRepo) GetUserIncomesBinance(ctx context.Context, userId uint64) ([]*biz.UserIncomeBinance, error) {
	var usersIncomes []*UserIncomeBinance
	if err := u.data.DB(ctx).Table("user_income_binance").Where("user_id=?", userId).Find(&usersIncomes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_INCOME_ERROR", err.Error())
	}

	res := make([]*biz.UserIncomeBinance, 0)
	for _, v := range usersIncomes {
		res = append(res, &biz.UserIncomeBinance{
			ID:        v.ID,
			UserId:    v.UserId,
			Symbol:    v.Symbol,
			Income:    v.Income,
			Info:      v.Info,
			TradeId:   v.TradeId,
			Time:      v.Time,
			TranId:    v.TranId,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserIncomesBinanceBySymbolAndTraderId .
func (u *UserdataRepo) GetUserIncomesBinanceBySymbolAndTraderId(ctx context.Context, userId uint64, symbol string, traderId string) (*biz.UserIncomeBinance, error) {
	var usersIncome *UserIncomeBinance
	if err := u.data.DB(ctx).Table("user_income_binance").Where("user_id=? and trader_id=? and symbol=?", userId, symbol, traderId).First(&usersIncome).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_INCOME_ERROR", err.Error())
	}

	return &biz.UserIncomeBinance{
		ID:        usersIncome.ID,
		UserId:    usersIncome.UserId,
		Symbol:    usersIncome.Symbol,
		Income:    usersIncome.Income,
		Info:      usersIncome.Info,
		TradeId:   usersIncome.TradeId,
		Time:      usersIncome.Time,
		TranId:    usersIncome.TranId,
		CreatedAt: usersIncome.CreatedAt,
		UpdatedAt: usersIncome.UpdatedAt,
	}, nil
}

// GetUserIncomesBinanceOrderIdDesc .
func (u *UserdataRepo) GetUserIncomesBinanceOrderIdDesc(ctx context.Context, userId uint64) (*biz.UserIncomeBinance, error) {
	var usersIncome *UserIncomeBinance
	if err := u.data.DB(ctx).Table("user_income_binance").Where("user_id=?", userId).Order("id desc").First(&usersIncome).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_INCOME_ERROR", err.Error())
	}

	return &biz.UserIncomeBinance{
		ID:        usersIncome.ID,
		UserId:    usersIncome.UserId,
		Symbol:    usersIncome.Symbol,
		Income:    usersIncome.Income,
		Info:      usersIncome.Info,
		TradeId:   usersIncome.TradeId,
		Time:      usersIncome.Time,
		TranId:    usersIncome.TranId,
		CreatedAt: usersIncome.CreatedAt,
		UpdatedAt: usersIncome.UpdatedAt,
	}, nil
}

// InsertUserIncomeBinance .
func (u *UserdataRepo) InsertUserIncomeBinance(ctx context.Context, userIncome *biz.UserIncomeBinance) error {
	insert := &UserIncomeBinance{
		UserId:  userIncome.UserId,
		Symbol:  userIncome.Symbol,
		Income:  userIncome.Income,
		Info:    userIncome.Info,
		TradeId: userIncome.TradeId,
		Time:    userIncome.Time,
		TranId:  userIncome.TranId,
	}

	res := u.data.DB(ctx).Table("user_income_binance").Create(&insert)
	if res.Error != nil {
		return errors.New(500, "CREATE_USER_INCOME_ERROR", "创建数据失败")
	}

	return nil
}
