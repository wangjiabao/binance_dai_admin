package biz

import (
	pb "binance_dai_admin/api/userdata/v1"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/antihax/optional"
	"github.com/gateio/gateapi-go/v6"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Userdata is a Userdata model.
type Userdata struct {
	Hello string
}

// UserdataRepo is a Greater repo.
type UserdataRepo interface {
	Save(context.Context, *Userdata) (*Userdata, error)
	GetUsers(ctx context.Context) ([]*User, error)
	GetUsersStatusOk(ctx context.Context) ([]*User, error)
	GetUsersByIds(ctx context.Context, userIds []uint64) ([]*User, error)
	GetUserById(ctx context.Context, userIds uint64) (*User, error)
	UpdateUserNum(ctx context.Context, apiKey string, num float64) error
	GetUserIncomesBinance(ctx context.Context, userId uint64) ([]*UserIncomeBinance, error)
	GetUserIncomesBinanceBySymbolAndTraderId(ctx context.Context, userId uint64, symbol string, traderId string) (*UserIncomeBinance, error)
	GetUserIncomesBinanceOrderIdDesc(ctx context.Context, userId uint64) (*UserIncomeBinance, error)
	GetUserIncomesBinanceByUserIds(ctx context.Context, userIds []uint64, start int64, end int64) ([]*UserIncomeBinance, error)
	InsertUserIncomeBinance(ctx context.Context, userIncome *UserIncomeBinance) error
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

type User struct {
	ID                  uint64
	Address             string
	ApiStatus           uint64
	ApiKey              string
	ApiSecret           string
	BindTraderStatus    uint64
	BindTraderStatusTfi uint64
	IsDai               uint64
	UseNewSystem        uint64
	CreatedAt           time.Time
	UpdatedAt           time.Time
	BinanceId           uint64
	OkxId               string
	NeedInit            uint64
	Dai                 uint64
	Plat                string
	OkxPass             string
	Num                 float64
	Ip                  string
}

func (uc *UserdataUsecase) GetUserOrders(ctx context.Context, req *pb.GetUserOrdersRequest) (*pb.GetUserOrdersReply, error) {

	var (
		err     error
		users   []*User
		symbols []string
	)
	users, err = uc.repo.GetUsers(ctx)
	if nil != err {
		return nil, err
	}

	symbols = make([]string, 0)
	symbols = append(symbols, "ETHUSDT")
	for _, vUser := range users {
		if 1 != vUser.ApiStatus { // 结束了
			continue
		}

		if "binance" == vUser.Plat {
			for _, vSymbol := range symbols {
				pullUserAllOrdersBinance(vSymbol, vUser.ApiKey, vUser.ApiSecret)
			}
		}
	}

	return &pb.GetUserOrdersReply{OrderId: 1}, nil
}

func pullUserAllOrdersBinance(symbol, apiKey, apiSecret string) {
	var (
		err           error
		binanceOrders []*binanceOrder
	)
	binanceOrders, err = getAllOrders(&AllOrdersParams{
		Symbol:    symbol,
		StartTime: 0,
		EndTime:   0,
		Limit:     10,
	}, apiKey, apiSecret)

	if nil != err {
		fmt.Println("binance，订单查询错误：", err)
	}

	for _, vOrder := range binanceOrders {
		fmt.Println(vOrder)
	}
}

// API Key 和 Secret
const (
	baseBinanceURL = "https://fapi.binance.com"
)

// AllOrdersParams 查询全部订单请求参数
type AllOrdersParams struct {
	Symbol     string `url:"symbol"`
	OrderID    int64  `url:"orderId,omitempty"`
	StartTime  int64  `url:"startTime,omitempty"`
	EndTime    int64  `url:"endTime,omitempty"`
	Limit      int    `url:"limit,omitempty"`
	RecvWindow int64  `url:"recvWindow,omitempty"`
	Timestamp  int64  `url:"timestamp"`
}

// Order 查询全部订单响应数据结构
type binanceOrder struct {
	OrderID       int64  `json:"orderId"`
	Symbol        string `json:"symbol"`
	Status        string `json:"status"`
	ClientOrderID string `json:"clientOrderId"`
	Price         string `json:"price"`
	AvgPrice      string `json:"avgPrice"`
	OrigQty       string `json:"origQty"`
	ExecutedQty   string `json:"executedQty"`
	CumQuote      string `json:"cumQuote"`
	TimeInForce   string `json:"timeInForce"`
	Type          string `json:"type"`
	Side          string `json:"side"`
	StopPrice     string `json:"stopPrice"`
	Time          int64  `json:"time"`
	UpdateTime    int64  `json:"updateTime"`
	WorkingType   string `json:"workingType"`
	ReduceOnly    bool   `json:"reduceOnly"`
	ClosePosition bool   `json:"closePosition"`
	ActivatePrice string `json:"activatePrice"`
	PriceRate     string `json:"priceRate"`
}

// 签名方法
func binanceSign(data string, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}

// 查询全部订单方法
func getAllOrders(params *AllOrdersParams, apiKey, apiSecret string) ([]*binanceOrder, error) {
	// 设置时间戳
	params.Timestamp = time.Now().UnixMilli()

	// 构造查询字符串
	query := url.Values{}
	query.Set("symbol", params.Symbol)
	if params.OrderID != 0 {
		query.Set("orderId", strconv.FormatInt(params.OrderID, 10))
	}
	if params.StartTime != 0 {
		query.Set("startTime", strconv.FormatInt(params.StartTime, 10))
	}
	if params.EndTime != 0 {
		query.Set("endTime", strconv.FormatInt(params.EndTime, 10))
	}
	if params.Limit != 0 {
		query.Set("limit", strconv.Itoa(params.Limit))
	}
	if params.RecvWindow != 0 {
		query.Set("recvWindow", strconv.FormatInt(params.RecvWindow, 10))
	}
	query.Set("timestamp", strconv.FormatInt(params.Timestamp, 10))

	// 签名
	query.Set("signature", binanceSign(query.Encode(), apiSecret))

	// 构造请求
	reqURL := fmt.Sprintf("%s/fapi/v1/allOrders?%s", baseBinanceURL, query.Encode())
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-MBX-APIKEY", apiKey)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	// 解析响应
	var orders []*binanceOrder
	err = json.Unmarshal(body, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

type UserIncomeBinance struct {
	ID        uint64
	UserId    uint64
	Symbol    string
	Income    string
	Info      string
	TradeId   string
	Time      uint64
	TranId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (uc *UserdataUsecase) PullUserIncome(ctx context.Context, req *pb.PullUserIncomeRequest) (*pb.PullUserIncomeReply, error) {

	var (
		err     error
		symbols []string
		users   []*User
	)
	users, err = uc.repo.GetUsersStatusOk(ctx)
	if nil != err {
		return nil, err
	}

	symbols = make([]string, 0)
	symbols = append(symbols, "ETHUSDT")

	for _, vUser := range users {
		if "binance" == vUser.Plat {
			for _, vSymbol := range symbols {
				uc.pullUserIncomeBinance(ctx, vUser.ID, vUser.CreatedAt, vSymbol, vUser.ApiKey, vUser.ApiSecret)
			}
		} else if "gate" == vUser.Plat {
			uc.pullUserIncomeGate(ctx, vUser.ID, vUser.CreatedAt, "ETH_USDT", vUser.ApiKey, vUser.ApiSecret)
		}
	}

	return &pb.PullUserIncomeReply{}, nil
}

func (uc *UserdataUsecase) pullUserIncomeBinance(ctx context.Context, userId uint64, timeC time.Time, symbol, apiKey, apiSecret string) {
	var (
		err     error
		incomes []*binanceOrderIncomeResponse
	)
	params := &binanceOrderIncomeParams{
		Symbol: symbol, // 可选，查询特定交易对
	}

	// 查询资金流水
	incomes, err = getIncomeHistory(apiKey, apiSecret, params)
	if nil != err {
		fmt.Println("binance，资金流水查询错误：", err)
	}
	if 0 >= len(incomes) {
		return
	}

	var (
		binanceIncome *UserIncomeBinance
		insert        []*UserIncomeBinance
	)
	binanceIncome, err = uc.repo.GetUserIncomesBinanceOrderIdDesc(ctx, userId)
	if nil != err {
		fmt.Println("binance，查询数据库错误：", err)
		return
	}

	insert = make([]*UserIncomeBinance, 0)
	lastKey := len(incomes) - 1
	for i := lastKey; i >= 0; i-- {
		tmp := incomes[i]

		if tmp.Time <= uint64(timeC.UnixMilli()) {
			continue
		}

		tranIdStr := strconv.FormatUint(tmp.TranID, 10)
		// 按插入倒序，最新一条在最后，默认系统和远程能返回正确的顺序，且远程每次返回的结果都一样，
		if nil != binanceIncome {
			if tmp.TradeID == binanceIncome.TradeId && tranIdStr == binanceIncome.TranId && tmp.Symbol == binanceIncome.Symbol {
				break
			}
		}

		insert = append(insert, &UserIncomeBinance{
			UserId:  userId,
			Symbol:  tmp.Symbol,
			Income:  tmp.Income,
			Info:    tmp.Info,
			TradeId: tmp.TradeID,
			Time:    tmp.Time,
			TranId:  tranIdStr,
		})
	}

	if 0 < len(insert) {
		for i := len(insert) - 1; i >= 0; i-- {
			err = uc.repo.InsertUserIncomeBinance(ctx, insert[i])
			if nil != err {
				fmt.Println("binance，插入数据库失败：", err, insert[i])
			}
		}
	}
}

// binanceOrderIncomeParams 查询资金流水请求参数结构体
type binanceOrderIncomeParams struct {
	Symbol     string // 可选，交易对
	IncomeType string // 可选，资金类型
	StartTime  int64  // 可选，起始时间
	EndTime    int64  // 可选，结束时间
	Limit      int    // 可选，返回的最大结果数量
	Timestamp  int64  // 必填，时间戳
}

// IncomeResponse 资金流水响应结构体
type binanceOrderIncomeResponse struct {
	Symbol     string `json:"symbol"`
	IncomeType string `json:"incomeType"`
	Income     string `json:"income"`
	Asset      string `json:"asset"`
	Info       string `json:"info"`
	Time       uint64 `json:"time"`
	TranID     uint64 `json:"tranId"`
	TradeID    string `json:"tradeId"`
}

// 查询资金流水方法
func getIncomeHistory(apiKey, apiSecret string, params *binanceOrderIncomeParams) ([]*binanceOrderIncomeResponse, error) {
	// 设置时间戳
	params.Timestamp = time.Now().UnixMilli()

	// 获取最近10分钟的
	endTime := strconv.FormatInt(params.Timestamp, 10)
	startTime := strconv.FormatInt(time.Now().Add(-7*24*time.Hour).UnixMilli(), 10)

	// 构造查询字符串
	query := url.Values{}
	query.Set("timestamp", strconv.FormatInt(params.Timestamp, 10))
	query.Set("incomeType", "REALIZED_PNL")
	query.Set("startTime", startTime)
	query.Set("endTime", endTime)
	query.Set("limit", "1000")
	if params.Symbol != "" {
		query.Set("symbol", params.Symbol)
	}

	//fmt.Println(startTime, endTime)

	// 签名
	query.Set("signature", binanceSign(query.Encode(), apiSecret))

	// 构造请求
	reqURL := fmt.Sprintf("%s/fapi/v1/income?%s", baseBinanceURL, query.Encode())
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-MBX-APIKEY", apiKey)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	// 解析响应
	var incomeResponse []*binanceOrderIncomeResponse
	err = json.Unmarshal(body, &incomeResponse)
	if err != nil {
		return nil, err
	}

	return incomeResponse, nil
}

func (uc *UserdataUsecase) pullUserIncomeGate(ctx context.Context, userId uint64, timeC time.Time, symbol, apiKey, apiSecret string) {

	var (
		res []gateapi.PositionClose
		err error
	)

	res, err = getGateContract(apiKey, apiSecret, symbol)
	if nil != err {
		fmt.Println(err, "获取gate信息失败", userId)
		return
	}

	if 0 >= len(res) {
		return
	}
	var (
		income *UserIncomeBinance
		insert []*UserIncomeBinance
	)
	income, err = uc.repo.GetUserIncomesBinanceOrderIdDesc(ctx, userId)
	if nil != err {
		fmt.Println("binance，查询数据库错误：", err)
		return
	}

	insert = make([]*UserIncomeBinance, 0)

	for _, v := range res {
		if v.FirstOpenTime <= timeC.Unix() {
			continue
		}

		if "api" != v.Text {
			continue
		}

		// 相同的一条
		if nil != income {
			if v.Pnl == income.Income && income.TradeId == strconv.FormatInt(v.FirstOpenTime, 10) {
				break
			}
		}

		insert = append(insert, &UserIncomeBinance{
			UserId:  userId,
			Symbol:  v.Contract,
			Income:  v.Pnl,
			Info:    "gate",
			TradeId: strconv.FormatInt(v.FirstOpenTime, 10),
			Time:    uint64(v.Time),
			TranId:  "",
		})
	}

	if 0 < len(insert) {
		for i := len(insert) - 1; i >= 0; i-- {
			err = uc.repo.InsertUserIncomeBinance(ctx, insert[i])
			if nil != err {
				fmt.Println("binance，插入数据库失败：", err, insert[i])
			}
		}
	}

	return
}

// getGateContract
func getGateContract(apiK, apiS string, symbol string) ([]gateapi.PositionClose, error) {
	client := gateapi.NewAPIClient(gateapi.NewConfiguration())
	// uncomment the next line if your are testing against testnet
	// client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
	ctx := context.WithValue(context.Background(),
		gateapi.ContextGateAPIV4,
		gateapi.GateAPIV4{
			Key:    apiK,
			Secret: apiS,
		},
	)

	result, _, err := client.FuturesApi.ListPositionClose(ctx, "usdt", &gateapi.ListPositionCloseOpts{
		Contract: optional.NewString(symbol),
	})
	if err != nil {
		if e, ok := err.(gateapi.GateAPIError); ok {
			fmt.Println("gate api error: ", e.Error())
		} else {
			fmt.Println("generic error: ", err.Error())
		}
	}

	return result, nil
}

func (uc *UserdataUsecase) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersReply, error) {
	var (
		users []*User
		err   error
	)

	users, err = uc.repo.GetUsers(ctx)
	if nil != err {
		return nil, err
	}

	res := make([]*pb.GetUsersReply_DataList, 0)
	if 0 >= len(users) {
		return &pb.GetUsersReply{List: res}, nil
	}

	ips := make(map[string]string, 0)
	for _, vUser := range users {
		if 10 <= len(vUser.Ip) {
			ips[vUser.Ip] = vUser.Ip
		}
	}

	// 获取其他服务器上的信息
	nums := make(map[string]*numList, 0)
	if 0 < len(ips) {
		for _, vIps := range ips {
			var (
				tmpNumList []*numList
			)
			tmpNumList, err = getNums(vIps)
			if nil != err {
				continue
			}

			if nil != tmpNumList && 0 < len(tmpNumList) {
				for _, vTmpNumList := range tmpNumList {
					nums[vTmpNumList.ApiKey] = vTmpNumList
				}
			}
		}
	}

	for _, v := range users {
		var tmpNum float64
		if 10 > len(v.Ip) { // 屏蔽
			continue
		}

		if _, ok := nums[v.ApiKey]; ok {
			tmpNum = nums[v.ApiKey].Num
		}

		res = append(res, &pb.GetUsersReply_DataList{
			Name:   v.Address,
			UserId: v.ID,
			Ip:     v.Ip,
			Num:    tmpNum,
		})
	}

	return &pb.GetUsersReply{List: res}, nil
}

func (uc *UserdataUsecase) GetUsersIncome(ctx context.Context, req *pb.GetUsersIncomeRequest) (*pb.GetUsersIncomeReply, error) {
	var (
		users   []*User
		userIds []uint64
		err     error
	)

	userIds = make([]uint64, 0)
	if 0 < req.UserId {
		userIds = append(userIds, req.UserId)
	}

	if 0 >= len(userIds) {
		return nil, nil
	}

	users, err = uc.repo.GetUsersByIds(ctx, userIds)
	if nil != err {
		return nil, err
	}

	userIds = make([]uint64, 0)
	usersMap := make(map[uint64]*User, 0)
	for _, vUser := range users {
		userIds = append(userIds, vUser.ID)
		usersMap[vUser.ID] = vUser
	}

	if _, ok := usersMap[req.UserId]; !ok {
		return nil, nil
	}

	// 时间参数，格式修正，binance毫秒级
	start := req.Start
	end := req.End
	if "binance" == usersMap[req.UserId].Plat {
		start *= 1000
		end *= 1000
	}

	var (
		userIncomes []*UserIncomeBinance
	)
	userIncomes, err = uc.repo.GetUserIncomesBinanceByUserIds(ctx, userIds, start, end)
	if nil != err {
		return nil, err
	}

	res := make([]*pb.GetUsersIncomeReply_DataList, 0)

	for _, v := range userIncomes {
		tmpTime := ""
		if "gate" == v.Info {
			tmpTime = time.Unix(int64(v.Time), 0).Add(8 * time.Hour).Format("2006-01-02 15:04:05")
		} else {
			tmpTime = time.UnixMilli(int64(v.Time)).Add(8 * time.Hour).Format("2006-01-02 15:04:05")
		}
		res = append(res, &pb.GetUsersIncomeReply_DataList{
			Symbol: v.Symbol,
			Income: v.Income,
			Time:   tmpTime,
		})
	}

	return &pb.GetUsersIncomeReply{List: res}, nil
}

type numListResponse struct {
	List []*numList `json:"list"`
}

type numList struct {
	Num    float64 `json:"num"`
	ApiKey string  `json:"apiKey"`
}

// 查询资金流水方法
func getNums(ip string) ([]*numList, error) {
	// 构造请求
	req, err := http.NewRequest("GET", "http://"+ip+"/api/binance_dai_admin/get_num", nil)
	if err != nil {
		return nil, err
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// 解析响应
	var numResponse *numListResponse
	err = json.Unmarshal(body, &numResponse)
	if err != nil {
		return nil, err
	}

	return numResponse.List, nil
}

func (uc *UserdataUsecase) GetNum(ctx context.Context, req *pb.GetNumRequest) (*pb.GetNumReply, error) {
	var (
		users []*User
		err   error
	)

	users, err = uc.repo.GetUsers(ctx)
	if nil != err {
		return nil, err
	}

	res := make([]*pb.GetNumReply_DataList, 0)
	if 0 >= len(users) {
		return &pb.GetNumReply{List: res}, nil
	}

	for _, v := range users {
		res = append(res, &pb.GetNumReply_DataList{
			ApiKey: v.ApiKey,
			Num:    v.Num,
		})
	}

	return &pb.GetNumReply{List: res}, nil
}

func (uc *UserdataUsecase) UpdateUserNum(ctx context.Context, req *pb.UpdateUserNumRequest) (*pb.UpdateUserNumReply, error) {
	var (
		user *User
		err  error
	)

	user, err = uc.repo.GetUserById(ctx, req.UserId)
	if nil != err {
		return nil, err
	}

	if nil == user {
		return nil, nil
	}

	err = setNum(user.Ip, user.ApiKey, req.Num)

	return nil, err
}

func (uc *UserdataUsecase) UpdateNum(ctx context.Context, req *pb.UpdateNumRequest) (*pb.UpdateNumReply, error) {
	return nil, uc.repo.UpdateUserNum(ctx, req.ApiKey, req.Num)
}

// 查询资金流水方法
func setNum(ip string, apiKey string, num float64) error {
	// 构造请求
	req, err := http.NewRequest("GET", "http://"+ip+"/api/binance_dai_admin/update_num?apikey="+apiKey+"&num="+strconv.FormatFloat(num, 'f', -1, 64), nil)
	if err != nil {
		return err
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {

		}
	}(resp.Body)

	return nil
}
