/*
 *   Copyright `IMDaDaGo` Author(https://houseme.github.io/imdadago/). All Rights Reserved.
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 *
 *   You can obtain one at https://github.com/houseme/imdadago.
 *
 */

package dadago

import (
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol"

	"github.com/houseme/imdadago/domain"
)

// Level is the log level.
type Level hlog.Level

// Logger is the logger.
type Logger hlog.FullLogger

// options is the configuration for the ImDada client.
type options struct {
	AppKey    string
	AppSecret string
	SourceID  string
	Gateway   string
	Callback  string
	ShopNo    string
	LogPath   string // 日志路径
	Level     Level  // 日志级别
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
func WithLevel(level Level) Option {
	return func(o *options) {
		o.Level = level
	}
}

// Client is the ImDada client.
type Client struct {
	request  *domain.Request
	response *protocol.Response
	log      Logger
	op       options
	gateway  string
}
