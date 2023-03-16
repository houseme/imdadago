/*
 *  Copyright ImDaDa-Go Author(https://houseme.github.io/imdada-go/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the Apache-2.0 License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/imdada-go.
 */

package dada

const (
	// version is the default version of ImDada.
	// 版本号，当前版本：1.0
	// See: http://newopen.imdada.cn/#/development/file/api
	version = "1.0"

	// format is the default format of ImDada.
	// 请求格式，暂时只支持json
	// See: http://newopen.imdada.cn/#/development/file/api
	format = "json"

	// gateway is the default gateway of ImDada.
	// See: http://newopen.imdada.cn/#/development/file/api
	gateway = "https://newopen.imdada.cn"

	// userAgent is the user agent of ImDada.
	// See: http://newopen.imdada.cn/#
	userAgent = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36`
)

// API is the interface of ImDada.
const (
	// recharge 获取充值链接
	recharge = "/api/recharge"

	// queryBalance 查询账户余额
	queryBalance = "/api/balance/query"

	// CityCodeList 获取城市编码列表
	cityCodeList = "/api/cityCode/list"

	// merchantCreate 注册商户账号
	merchantCreate = "/merchantApi/merchant/add"

	// shopCreate 创建门店
	shopCreate = "/api/shop/add"

	// shopUpdate 更新门店
	shopUpdate = "/api/shop/update"

	// shopQuery 查询门店
	shopQuery = "/api/shop/detail"

	// ordersCreate 创建订单
	ordersCreate = "/api/order/addOrder"

	// orderReCreate 重新创建订单
	orderReCreate = "/api/order/reAddOrder"

	// orderDeliverFeeQuery 查询订单配送费
	orderDeliverFeeQuery = "/api/order/queryDeliverFee"

	// orderCreateAfterQuery 查询订单创建后状态
	orderCreateAfterQuery = "/api/order/status/addAfterQuery"

	// orderAddTip 增加小费
	orderAddTip = "/api/order/addTip"

	// orderStatusQuery 查询订单状态
	orderStatusQuery = "/api/order/status/query"

	// orderCancel 取消订单
	// 在订单待接单或待取货情况下，调用此接口可取消订单。
	// 取消费用说明：接单1 分钟以内取消不扣违约金；
	// 接单后1－15分钟内取消订单，运费退回。同时扣除2元作为给配送员的违约金；
	// 配送员接单后15 分钟未到店，商户取消不扣违约金；
	// 系统取消订单说明：超过72小时未接单系统自动取消。每天凌晨2点，取消大于72小时未完成的订单。
	orderCancel = "/api/order/formalCancel"

	// additionalOrders 创建追加订单
	additionalOrders = "/api/order/appoint/exist"

	// cancelAppointOrders 取消追加订单
	cancelAppointOrders = "/api/order/appoint/cancel"

	// transportAppointList 骑手追加订单列表
	transportAppointList = "/api/order/appoint/list/transporter"

	// complaintCreate 创建投诉
	complaintCreate = "/api/complaint/dada"

	// orderConfirmGoods 商户确认物品已返还
	orderConfirmGoods = "/api/order/confirm/goods"

	// messageConfirm 商户确认消息已读
	messageConfirm = "/api/message/confirm"

	// transporterPositionList 骑手列表
	transporterPositionList = "/api/order/transporter/position"

	// transporterTrack 骑手轨迹
	transporterTrack = "/api/order/transporter/track"

	// fetchCodeModify 修改取货码
	fetchCodeModify = "/api/order/fetchCode/update"
)
