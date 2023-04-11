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

// Package domain is the domain of ImDaDa.
// See: http://newopen.imdada.cn/#/development/file/merchantIndex
package domain

// CityListQueryRequest is the request of CityListQuery.
type CityListQueryRequest struct {
}

// CityListQueryResponse is the response of CityListQuery.
type CityListQueryResponse struct {
	Status    string      `json:"status"`
	Result    []*CityItem `json:"result"`
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Fail      bool        `json:"fail"`
	Success   bool        `json:"success"`
	ErrorCode int         `json:"errorCode"`
}

// CityItem is the item of CityListQuery.
type CityItem struct {
	CityName string `json:"cityName"`
	CityCode string `json:"cityCode"`
}

// MerchantCreateRequest is the request of MerchantCreate.
type MerchantCreateRequest struct {
	CityName          string `json:"city_name"`
	ContactName       string `json:"contact_name"`
	ContactPhone      string `json:"contact_phone"`
	Email             string `json:"email"`
	EnterpriseAddress string `json:"enterprise_address"`
	EnterpriseName    string `json:"enterprise_name"`
	Mobile            string `json:"mobile"`
}

// MerchantCreateResponse is the response of MerchantCreate.
type MerchantCreateResponse struct {
	Status    string `json:"status"`
	Result    int    `json:"result"` // 商户id
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Fail      bool   `json:"fail"`
	Success   bool   `json:"success"`
	ErrorCode int    `json:"errorCode"`
}

// ShopCreateRequest is the request of ShopCreate.
type ShopCreateRequest []*ShopCreateItem

// ShopCreateItem is the item of ShopCreate.
// See: http://newopen.imdada.cn/#/development/file/shopAdd
type ShopCreateItem struct {
	StationName    string  `json:"station_name"`
	OriginShopID   string  `json:"origin_shop_id,omitempty"`
	StationAddress string  `json:"station_address"`
	ContactName    string  `json:"contact_name"`
	Business       int     `json:"business"`
	Lng            float64 `json:"lng"`
	Phone          string  `json:"phone"`
	Lat            float64 `json:"lat"`
	IDCard         string  `json:"id_card,omitempty"`
	Password       string  `json:"password,omitempty"`
	Username       string  `json:"username,omitempty"`
	SettlementType int     `json:"settlement_type,omitempty"`
}

// ShopCreateResponse is the response of ShopCreate.
type ShopCreateResponse struct {
	Status    string            `json:"status"`
	Result    *ShopCreateResult `json:"result"`
	Code      int               `json:"code"`
	Msg       string            `json:"msg"`
	ErrorCode int               `json:"errorCode"`
}

// ShopCreateResult is the result of ShopCreate.
type ShopCreateResult struct {
	Success     int                      `json:"success"`
	SuccessList []*ShopCreateSuccessItem `json:"successList"`
	FailedList  []*ShopCreateFailedItem  `json:"failedList"`
}

// ShopCreateSuccessItem is the item of ShopCreateSuccess.
type ShopCreateSuccessItem struct {
	Phone          string  `json:"phone"`
	Business       int     `json:"business"`
	Lng            float64 `json:"lng"`
	Lat            float64 `json:"lat"`
	StationName    string  `json:"stationName"`
	OriginShopID   string  `json:"originShopId"`
	ContactName    string  `json:"contactName"`
	StationAddress string  `json:"stationAddress"`
	CityName       string  `json:"cityName"`
	AreaName       string  `json:"areaName"`
}

// ShopCreateFailedItem is the item of ShopCreateFailed.
type ShopCreateFailedItem struct {
	ShopNo   string `json:"shopNo"`
	Msg      string `json:"msg"`
	ShopName string `json:"shopName"`
}

// ShopUpdateRequest is the request of ShopUpdate.
type ShopUpdateRequest struct {
	Business       int     `json:"business,omitempty"`
	ContactName    string  `json:"contact_name,omitempty"`
	Lat            float64 `json:"lat,omitempty"`
	Lng            float64 `json:"lng,omitempty"`
	OriginShopID   string  `json:"origin_shop_id"`
	Phone          string  `json:"phone,omitempty"`
	StationAddress string  `json:"station_address,omitempty"`
	StationName    string  `json:"station_name,omitempty"`
}

// ShopUpdateResponse is the response of ShopUpdate.
type ShopUpdateResponse struct {
	Status    string `json:"status"`
	Result    int    `json:"result"` // 商户id
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Fail      bool   `json:"fail"`
	Success   bool   `json:"success"`
	ErrorCode int    `json:"errorCode"`
}

// ShopQueryRequest is the request of ShopQuery.
type ShopQueryRequest struct {
	OriginShopID string `json:"origin_shop_id"`
}

// ShopQueryResponse is the response of ShopQuery.
type ShopQueryResponse struct {
	Status    string         `json:"status"`
	Result    *ShopQueryItem `json:"result"`
	Code      int            `json:"code"`
	Msg       string         `json:"msg"`
	Fail      bool           `json:"fail"`
	Success   bool           `json:"success"`
	ErrorCode int            `json:"errorCode"`
}

// ShopQueryItem is the item of ShopQuery.
type ShopQueryItem struct {
	StationName    string  `json:"station_name"`
	AreaName       string  `json:"area_name"`
	StationAddress string  `json:"station_address"`
	CityName       string  `json:"city_name"`
	ContactName    string  `json:"contact_name"`
	OriginShopID   string  `json:"origin_shop_id"`
	Business       int     `json:"business"`
	Lng            float64 `json:"lng"`
	Phone          string  `json:"phone"`
	IDCard         string  `json:"id_card"`
	Lat            float64 `json:"lat"`
	Status         int     `json:"status"`
	ApproveStatus  int     `json:"approveStatus"`
}
