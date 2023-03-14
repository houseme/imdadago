/*
 *  Copyright ImDaDa-Go Author(https://houseme.github.io/imdada-go/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the Apache-2.0 License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/imdada-go.
 */

// Package dada is the ImDaDa-go client.
package dada

import (
	"context"
	"crypto/tls"
	"os"
	"time"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/houseme/imdada-go/log"
)

// options is the configuration for the ImDada client.
type options struct {
	AppKey    string
	AppSecret string
	SourceID  string
	Gateway   string
	Callback  string
	ShopNo    string
	LogPath   string    // 日志路径
	Level     log.Level // 日志级别
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
func WithLevel(level log.Level) Option {
	return func(o *options) {
		o.Level = level
	}
}

// Client is the ImDada client.
type Client struct {
	request  *protocol.Request
	response *protocol.Response
	log      log.ILogger
	op       options
}

// New creates a new ImDada client.
func New(ctx context.Context, opts ...Option) *Client {
	op := options{
		TimeOut:   5 * time.Second,
		UserAgent: userAgent,
		Gateway:   gateway,
		Level:     log.DebugLevel,
		LogPath:   os.TempDir(),
	}

	for _, option := range opts {
		option(&op)
	}

	return &Client{
		op:       op,
		log:      log.New(ctx, log.WithLevel(op.Level), log.WithLogPath(op.LogPath)),
		request:  &protocol.Request{},
		response: &protocol.Response{},
	}
}

// doRequest does the request.
func (c *Client) doRequest(ctx context.Context, formData map[string]string) error {
	c.log.Debug(ctx, "formData:", formData)
	c.request.SetMultipartFormData(formData)
	c.request.SetRequestURI(c.op.Gateway)
	c.request.Header.SetMethod(consts.MethodPost)
	c.request.Header.SetUserAgentBytes(c.op.UserAgent)
	c.log.Debug(ctx, "request content: ", c.request)

	hertz, err := client.NewClient(client.WithTLSConfig(&tls.Config{
		InsecureSkipVerify: true,
	}), client.WithDialTimeout(c.op.TimeOut))
	if err != nil {
		return err
	}

	c.log.Debug(ctx, "do request start")
	err = hertz.Do(ctx, c.request, c.response)
	if err != nil {
		return err
	}
	c.log.Debug(ctx, "do request end")
	var resp map[string]interface{}
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return err
	}
	return nil
}
