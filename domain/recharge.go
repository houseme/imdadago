/*
 *  Copyright `IMDaDaGo` Author(https://houseme.github.io/imdadago/). All Rights Reserved.
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
 *  You can obtain one at https://github.com/houseme/imdadago.
 */

// Package domain is the domain of ImDaDa.
// See: http://newopen.imdada.cn/#/development/file/rechargeIndex
package domain

// RechargeRequest is the request of recharge.
// See: http://newopen.imdada.cn/#/development/file/recharge
type RechargeRequest struct {
	Amount    float64 `json:"amount,string"`        // 充值金额（单位元，可以精确到分）
	Category  string  `json:"category"`             // 生成链接适应场景（category有二种类型值：PC、H5）
	NotifyURL string  `json:"notify_url,omitempty"` // 支付成功后跳转的页面（支付宝在支付成功后可以跳转到某个指定的页面，微信支付不支持）
	ShopNo    string  `json:"shop_no,omitempty"`    // 门店编号。如需要为商户账号下独立结算子门店充值，则需要传入(充值到大客户账户则不传)，如门店非独立结算则返回错误
}

// RechargeResponse is the response of recharge.
// See: http://newopen.imdada.cn/#/development/file/recharge
type RechargeResponse struct {
	Status  string `json:"status"`
	Result  string `json:"result"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
	Fail    bool   `json:"fail"`
}

// QueryBalanceRequest is the request of QueryBalance.
// See: http://newopen.imdada.cn/#/development/file/balanceQuery
type QueryBalanceRequest struct {
	Category int    `json:"category"`          // 查询运费账户类型（1：运费账户；2：红包账户，3：所有），默认查询运费账户余额
	ShopNo   string `json:"shop_no,omitempty"` // 门店编号。如需要查询大客户下独立结算子门店余额，则需要传入(查询大客户账户则不传)，如门店非独立结算则返回0
}

// QueryBalanceResponse is the response of QueryBalance.
// See: http://newopen.imdada.cn/#/development/file/balanceQuery
// 运费账户或红包账户的余额。如未传入门店编号字段，则返回大客户账户余额，
// 如传入门店编号且为独立结算则返回子门店账户余额，如门店非独立结算则返回0
type QueryBalanceResponse struct {
	Result  *BalanceResult `json:"result"`
	Status  string         `json:"status"`
	Msg     string         `json:"msg"`
	Code    int            `json:"code"`
	Success bool           `json:"success"`
	Fail    bool           `json:"fail"`
}

// BalanceResult is the result of QueryBalance.
type BalanceResult struct {
	RedPacketBalance float64 `json:"redPacketBalance"`
	DeliverBalance   float64 `json:"deliverBalance"`
}
