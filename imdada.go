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
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/houseme/imdada-go/domain"
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
		log:      nil,
		response: &protocol.Response{},
		request: &domain.Request{
			AppKey:   op.AppKey,
			V:        version,
			Format:   format,
			SourceID: op.SourceID,
		},
	}
	c.initLog(ctx, op)
	return c
}

// generateTimestamp Generate current time
func (c *Client) generateTimestamp() {
	c.request.Timestamp = time.Now().Unix()
}

// md5Sign
func (c *Client) md5Sign() {
	var (
		h       = md5.New()
		builder strings.Builder
	)
	builder.WriteString(c.op.AppSecret)
	builder.WriteString("app_key" + c.request.AppKey)
	builder.WriteString("body" + c.request.Body)
	builder.WriteString("format" + c.request.Format)
	builder.WriteString("source_id" + c.request.SourceID)
	builder.WriteString("timestamp" + strconv.FormatInt(c.request.Timestamp, 10))
	builder.WriteString("v" + c.request.V)
	builder.WriteString(c.op.AppSecret)
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
func (c *Client) QueryBalance(ctx context.Context, req *domain.QueryBalanceRequest) (resp *domain.QueryBalanceResponse, err error) {
	if c.request.Body, err = sonic.MarshalString(req); err != nil {
		return nil, err
	}
	if err = c.doRequest(ctx, queryBalance); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}
	return
}
