/*
 *  Copyright `IMDaDa-Go` Author(https://houseme.github.io/imdada-go/). All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *  You can obtain one at https://github.com/houseme/imdada-go.
 */

// Package dada is the ImDaDa-go client.
package dada

import (
	"context"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"errors"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/houseme/imdada-go/domain"
	"github.com/houseme/imdada-go/internal/log"
)

// options is the configuration for the ImDada client.
type options struct {
	AppKey    string
	AppSecret string
	SourceID  string
	Gateway   string
	Callback  string
	ShopNo    string
	LogPath   string     // 日志路径
	Level     hlog.Level // 日志级别
	TimeOut   time.Duration
	UserAgent []byte
	Debug     bool
}

// Option the option is an ImDada option.
type Option func(o *options)

// WithAppKey sets the app key.
func WithAppKey(appKey string) Option {
	return func(o *options) {
		o.AppKey = appKey
	}
}

// WithAppSecret sets the app secret.
func WithAppSecret(appSecret string) Option {
	return func(o *options) {
		o.AppSecret = appSecret
	}
}

// WithSourceID sets the source id.
func WithSourceID(sourceID string) Option {
	return func(o *options) {
		o.SourceID = sourceID
	}
}

// WithGateway sets the gateway.
func WithGateway(gateway string) Option {
	return func(o *options) {
		o.Gateway = gateway
	}
}

// WithTimeOut sets the timeout.
func WithTimeOut(timeout time.Duration) Option {
	return func(o *options) {
		o.TimeOut = timeout
	}
}

// WithUserAgent sets the user agent.
func WithUserAgent(userAgent []byte) Option {
	return func(o *options) {
		o.UserAgent = userAgent
	}
}

// WithDebug sets the debug.
func WithDebug(debug bool) Option {
	return func(o *options) {
		o.Debug = debug
	}
}

// WithCallback sets the callback.
func WithCallback(callback string) Option {
	return func(o *options) {
		o.Callback = callback
	}
}

// WithShopNo sets the shop no.
func WithShopNo(shopNo string) Option {
	return func(o *options) {
		o.ShopNo = shopNo
	}
}

// WithLogPath sets the log path.
func WithLogPath(logPath string) Option {
	return func(o *options) {
		o.LogPath = logPath
	}
}

// WithLevel sets the log level.
func WithLevel(level hlog.Level) Option {
	return func(o *options) {
		o.Level = level
	}
}

// Client is the ImDada client.
type Client struct {
	request  *domain.Request
	response *protocol.Response
	log      hlog.FullLogger
	op       options
	gateway  string
}

// New creates a new ImDada client.
func New(ctx context.Context, opts ...Option) *Client {
	op := options{
		TimeOut:   5 * time.Second,
		UserAgent: []byte(userAgent),
		Gateway:   gateway,
		Level:     hlog.LevelDebug,
		LogPath:   os.TempDir(),
	}

	for _, option := range opts {
		option(&op)
	}

	c := &Client{
		op:       op,
		log:      log.InitLog(ctx, op.LogPath, op.Level),
		response: &protocol.Response{},
		request: &domain.Request{
			AppKey:   op.AppKey,
			V:        version,
			Format:   format,
			SourceID: op.SourceID,
		},
	}
	c.log.SetLevel(op.Level)
	c.log.CtxInfof(ctx, "im dada init client start level:%s", op.Level)
	return c
}

// generateTimestamp Generate current time
func (c *Client) generateTimestamp() {
	c.request.Timestamp = time.Now().Unix()
}

// md5Sign
func (c *Client) md5Sign() {
	var builder strings.Builder
	builder.WriteString(c.op.AppSecret)
	builder.WriteString("app_key" + c.request.AppKey)
	builder.WriteString("body" + c.request.Body)
	builder.WriteString("format" + c.request.Format)
	builder.WriteString("source_id" + c.request.SourceID)
	builder.WriteString("timestamp" + strconv.FormatInt(c.request.Timestamp, 10))
	builder.WriteString("v" + c.request.V)
	builder.WriteString(c.op.AppSecret)
	h := md5.New()
	h.Write([]byte(builder.String()))
	c.request.Signature = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

// initRequest
func (c *Client) initRequest(method string) {
	c.generateTimestamp()
	c.md5Sign()
	c.gateway = c.op.Gateway + method
}

// doRequest does the request.
func (c *Client) doRequest(ctx context.Context, method string) error {
	c.initRequest(method)
	c.log.CtxDebugf(ctx, "request data: %+v", c.request)
	jsonBytes, err := sonic.Marshal(c.request)
	if err != nil {
		return err
	}
	c.log.CtxDebugf(ctx, "jsonBytes: %s", string(jsonBytes))
	request := &protocol.Request{}
	request.SetBody(jsonBytes)
	request.Header.SetContentTypeBytes([]byte("application/json"))
	request.Header.Set("accept", "application/json")
	c.log.CtxDebugf(ctx, "request url: %s", c.gateway)
	request.SetRequestURI(c.gateway)
	request.Header.SetMethod(consts.MethodPost)
	request.Header.SetUserAgentBytes(c.op.UserAgent)
	c.log.CtxDebugf(ctx, "request create end")

	hertz, err := client.NewClient(client.WithTLSConfig(&tls.Config{
		InsecureSkipVerify: true,
	}), client.WithDialTimeout(c.op.TimeOut))
	if err != nil {
		return err
	}

	c.log.CtxDebugf(ctx, "do request start")
	if err = hertz.Do(ctx, request, c.response); err != nil {
		return err
	}
	return nil
}

// QueryBalance query balance.
// 查询账户余额 url: http://newopen.imdada.cn/#/development/file/balanceQuery
func (c *Client) QueryBalance(ctx context.Context, req *domain.QueryBalanceRequest) (resp *domain.QueryBalanceResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "QueryBalance request data: %s", c.request.Body)
	if err = c.doRequest(ctx, queryBalance); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "QueryBalance response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}
	return
}

// Recharge account recharge.
// 获取充值链接 url: http://newopen.imdada.cn/#/development/file/recharge
func (c *Client) Recharge(ctx context.Context, req *domain.RechargeRequest) (resp *domain.RechargeResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "Recharge request data: %s", c.request.Body)
	if err = c.doRequest(ctx, recharge); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "Recharge response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}
	return
}

// CreateMerchant create merchant.
// 添加商户 url: http://newopen.imdada.cn/#/development/file/merchantAdd
func (c *Client) CreateMerchant(ctx context.Context, req *domain.MerchantCreateRequest) (resp *domain.MerchantCreateResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.request.SourceID = ""
	c.log.CtxDebugf(ctx, "CreateMerchant request data: %s", c.request.Body)
	if err = c.doRequest(ctx, merchantCreate); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "CreateMerchant response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}
	return
}

// CreateShop create shop.
// 添加门店 url: http://newopen.imdada.cn/#/development/file/shopAdd
func (c *Client) CreateShop(ctx context.Context, req *domain.ShopCreateRequest) (resp *domain.ShopCreateResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "CreateShop request data: %s", c.request.Body)
	if err = c.doRequest(ctx, shopCreate); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "CreateShop response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}
	return
}

// ModifyShop modify shop.
// 编辑门店 url: http://newopen.imdada.cn/#/development/file/shopUpdate
func (c *Client) ModifyShop(ctx context.Context, req *domain.ShopUpdateRequest) (resp *domain.ShopUpdateResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "ModifyShop request data: %s", c.request.Body)
	if err = c.doRequest(ctx, shopUpdate); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "ModifyShop response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}
	return
}

// QueryShop query shop.
// 门店详情 url: http://newopen.imdada.cn/#/development/file/shopDetail
func (c *Client) QueryShop(ctx context.Context, req *domain.ShopQueryRequest) (resp *domain.ShopQueryResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "QueryShop request data: %s", c.request.Body)
	if err = c.doRequest(ctx, shopQuery); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "QueryShop response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}
	return
}

// QueryCity query city list
// 获取城市信息列表 http://newopen.imdada.cn/#/development/file/cityList
func (c *Client) QueryCity(ctx context.Context, req *domain.CityListQueryRequest) (resp *domain.CityListQueryResponse, err error) {
	if req == nil {
		c.request.Body = ""
	} else {
		if c.request.Body, err = sonic.MarshalString(req); err != nil {
			return nil, err
		}
	}
	c.log.CtxDebugf(ctx, "QueryCity request data: %s ", c.request.Body)
	if err = c.doRequest(ctx, cityCodeList); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "QueryCity response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}
	return
}

// CreateOrder create order.
// 添加订单 url: http://newopen.imdada.cn/#/development/file/add
func (c *Client) CreateOrder(ctx context.Context, req *domain.OrdersCreateRequest) (resp *domain.OrdersCreateResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "CreateOrder request data: %s", c.request.Body)
	if err = c.doRequest(ctx, ordersCreate); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "CreateOrder response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}

	return
}

// ReCreateOrder recreate order.
// 重新发布订单 url: http://newopen.imdada.cn/#/development/file/reAdd
func (c *Client) ReCreateOrder(ctx context.Context, req *domain.OrdersCreateRequest) (resp *domain.OrdersCreateResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "ReCreateOrder request data: %s", c.request.Body)
	if err = c.doRequest(ctx, orderReCreate); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "ReCreateOrder response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}

	return
}

// QueryDeliverFee query deliver fee.
// 订单运费查询 url: http://newopen.imdada.cn/#/development/file/readyAdd
func (c *Client) QueryDeliverFee(ctx context.Context, req *domain.DeliverFeeQueryRequest) (resp *domain.DeliverFeeQueryResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "QueryDeliverFee request data: %s", c.request.Body)
	if err = c.doRequest(ctx, orderDeliverFeeQuery); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "QueryDeliverFee response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}

	return
}

// OrdersCreateByDeliverFeeQuery create order by deliver the fee query.
// 通过运费接口创建订单 url: http://newopen.imdada.cn/#/development/file/addAfterQuery
func (c *Client) OrdersCreateByDeliverFeeQuery(ctx context.Context, req *domain.OrdersCreateByDeliverFeeQueryRequest) (resp *domain.OrdersCreateByDeliverFeeQueryResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "OrdersCreateByDeliverFeeQuery request data: %s", c.request.Body)
	if err = c.doRequest(ctx, orderCreateAfterQuery); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "OrdersCreateByDeliverFeeQuery response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}

	return
}

// OrdersAddTip add tip.
// 添加小费 url: http://newopen.imdada.cn/#/development/file/addTip
func (c *Client) OrdersAddTip(ctx context.Context, req *domain.OrdersAddTipRequest) (resp *domain.OrdersAddTipResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "OrdersAddTip request data: %s", c.request.Body)
	if err = c.doRequest(ctx, orderAddTip); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "OrdersAddTip response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}

	return
}

// QueryOrderStatus query order status.
// 订单详情查询 url: http://newopen.imdada.cn/#/development/file/statusQuery
func (c *Client) QueryOrderStatus(ctx context.Context, req *domain.OrdersQueryRequest) (resp *domain.OrdersQueryResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "QueryOrderStatus request data: %s", c.request.Body)
	if err = c.doRequest(ctx, orderStatusQuery); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "QueryOrderStatus response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}

	return
}

// CancelOrder cancel order.
// 取消订单 url: http://newopen.imdada.cn/#/development/file/formalCancel
func (c *Client) CancelOrder(ctx context.Context, req *domain.OrdersCancelRequest) (resp *domain.OrdersCancelResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "CancelOrder request data: %s", c.request.Body)
	if err = c.doRequest(ctx, orderCancel); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "CancelOrder response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}

	return
}

// AdditionalOrders additional order.
// 增加订单 url: http://newopen.imdada.cn/#/development/file/appointOrder
func (c *Client) AdditionalOrders(ctx context.Context, req *domain.OrdersAddAppointRequest) (resp *domain.OrdersAddAppointResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "AdditionalOrders request data: %s", c.request.Body)
	if err = c.doRequest(ctx, additionalOrders); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "AdditionalOrders response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}

	return
}

// CancelTheAddOnOrder cancel appoint order CancelTheAddOnOrder
// 取消预约单 url: http://newopen.imdada.cn/#/development/file/appointOrderCancel
func (c *Client) CancelTheAddOnOrder(ctx context.Context, req *domain.OrdersCancelAppointRequest) (resp *domain.OrdersCancelAppointResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "CancelTheAddOnOrder request data: %s", c.request.Body)
	if err = c.doRequest(ctx, cancelAppointOrders); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "CancelTheAddOnOrder response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}

	return
}

// QueriesCanAppendKnights query can append knights.
// 查询可追加骑士 url: http://newopen.imdada.cn/#/development/file/listTransportersToAppoint
func (c *Client) QueriesCanAppendKnights(ctx context.Context, req *domain.OrdersAppointTransporterRequest) (resp *domain.OrdersAppointTransporterResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "QueriesCanAppendKnights request data: %s", c.request.Body)
	if err = c.doRequest(ctx, transportAppointList); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "QueriesCanAppendKnights response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}

	return
}

// CreateAComplaint create a complaint.
// 创建投诉 url: http://newopen.imdada.cn/#/development/file/complaintDada
func (c *Client) CreateAComplaint(ctx context.Context, req *domain.ComplaintRequest) (resp *domain.ComplaintResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "CreateAComplaint request data: %s", c.request.Body)
	if err = c.doRequest(ctx, complaintCreate); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "CreateAComplaint response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}

	return
}

// QueryComplaint query complaint.
// 查询投诉 url: http://newopen.imdada.cn/#/development/file/queryComplaintDada
func (c *Client) QueryComplaint(ctx context.Context, req *domain.ComplaintReasonRequest) (resp *domain.ComplaintReasonResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "QueryComplaint request data: %s", c.request.Body)
	if err = c.doRequest(ctx, complaintReasons); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "QueryComplaint response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}

	return
}

// OrderConfirmGoods order confirm goods.
// 商户确认物品已返还
func (c *Client) OrderConfirmGoods(ctx context.Context, req *domain.OrdersConfirmGoodsRequest) (resp *domain.OrdersConfirmGoodsResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "OrderConfirmGoods request data: %s", c.request.Body)
	if err = c.doRequest(ctx, orderConfirmGoods); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "OrderConfirmGoods response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}
	return
}

// OrderConfirmCancel order confirm cancel.
// 商户审核骑士取消订单 url: http://newopen.imdada.cn/#/development/file/applicationCancel
func (c *Client) OrderConfirmCancel(ctx context.Context, req *domain.OrdersTransporterCancelAsyncConfirmRequest) (resp *domain.OrdersTransporterCancelAsyncConfirmResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "OrderConfirmCancel request data: %s", c.request.Body)
	if err = c.doRequest(ctx, messageConfirm); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "OrderConfirmCancel response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}
	return
}

// QueryTransporterPosition query transporter position.
// 查询骑士位置 url: http://newopen.imdada.cn/#/development/file/queryLocation
func (c *Client) QueryTransporterPosition(ctx context.Context, req *domain.OrdersTransporterPositionRequest) (resp *domain.OrdersTransporterPositionResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "QueryTransporterPosition request data: %s", c.request.Body)
	if err = c.doRequest(ctx, transporterPosition); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "QueryTransporterPosition response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}
	return
}

// QueryTransporterTrack query transporter track.
// 查询骑士轨迹 url: http://newopen.imdada.cn/#/development/file/queryDeliverTrack
func (c *Client) QueryTransporterTrack(ctx context.Context, req *domain.OrdersTransporterTrackRequest) (resp *domain.OrdersTransporterTrackResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "QueryTransporterTrack request data: %s", c.request.Body)
	if err = c.doRequest(ctx, transporterTrack); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "QueryTransporterTrack response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}
	return
}

// ModifyFetchCode modify fetch code.
// 修改取货码 url: http://newopen.imdada.cn/#/development/file/modifyFetchCode
func (c *Client) ModifyFetchCode(ctx context.Context, req *domain.OrdersFetchCodeModifyRequest) (resp *domain.OrdersFetchCodeModifyResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "ModifyFetchCode request data: %s", c.request.Body)
	if err = c.doRequest(ctx, fetchCodeModify); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "ModifyFetchCode response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}
	return
}

//

// VerifySignature verify signature.
func (c *Client) VerifySignature(_ context.Context, updateTime int64, clientID, orderID, signature string) (err error) {
	var list []string
	list = append(list, strconv.FormatInt(updateTime, 10))
	list = append(list, clientID)
	list = append(list, orderID)
	sort.Strings(list)
	sign := strings.Join(list, "")
	h := md5.New()
	h.Write([]byte(sign))
	if hex.EncodeToString(h.Sum(nil)) != signature {
		return errors.New("signature error")
	}
	return
}
