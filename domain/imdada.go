/*
 *  Copyright ImDaDa-Go Author(https://houseme.github.io/imdada-go/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the Apache-2.0 License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/imdada-go.
 */

// Package domain is the domain of ImDaDa.
package domain

// Request is the request of ImDaDa.
type Request struct {
	AppKey    string `json:"app_key"`
	Signature string `json:"signature"`
	V         string `json:"v"`
	Format    string `json:"format"`
	SourceID  string `json:"source_id"`
	Body      string `json:"body"`
	Timestamp int64  `json:"timestamp"`
}

// Response is the response of ImDaDa.
type Response struct {
	Status    string      `json:"status"`
	ErrorCode int         `json:"errorCode"`
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Result    interface{} `json:"result"`
	Success   bool        `json:"success"`
	Fail      bool        `json:"fail"`
}
