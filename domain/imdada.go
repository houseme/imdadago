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
