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

	// ComplaintReasons 投诉原因列表
	complaintReasons = "/api/complaint/reasons"

	// orderConfirmGoods 商户确认物品已返还
	orderConfirmGoods = "/api/order/confirm/goods"

	// messageConfirm 商户确认消息已读
	messageConfirm = "/api/message/confirm"

	// transporterPosition 查询骑士位置
	transporterPosition = "/api/order/transporter/position"

	// transporterTrack 骑手轨迹
	transporterTrack = "/api/order/transporter/track"

	// fetchCodeModify 修改取货码
	fetchCodeModify = "/api/order/fetchCode/update"
)

const (
	// rechargeCateH5 H5充值
	rechargeCateH5 = "H5"
	// rechargeCatePC PC充值
	rechargeCatePC = "PC"
)

const (
	foodSnacks            = 1  // 食品小吃
	drink                 = 2  // 饮料
	flowersAndGreenery    = 3  // 鲜花绿植
	other                 = 5  // 其他
	printingTicketing     = 8  // 文印票务
	convenienceStores     = 9  // 便利店
	freshFruit            = 13 // 水果生鲜
	intraCityECommerce    = 19 // 同城电商
	medicine              = 20 // 医药
	cake                  = 21 // 蛋糕
	wine                  = 24 // 酒品
	smallCommodityMarkets = 25 // 小商品市场
	clothing              = 26 // 服装
	autoRepairParts       = 27 // 汽修零配
	digitalAppliances     = 28 // 数码家电
	crayfishBBQ           = 29 // 小龙虾/烧烤
	supermarket           = 31 // 超市
	chafingDish           = 51 // 火锅
	personalCareMakeup    = 53 // 个护美妆
	mother                = 55 // 母婴
	homeTextiles          = 57 // 家居家纺
	cellPhone             = 59 // 手机
	home                  = 61 // 家装
	adultProducts         = 63 // 成人用品
	campus                = 65 // 校园
	highEndMarket         = 66 // 高端市场
)

// RechargeCateH5 H5充值
func RechargeCateH5() string {
	return rechargeCateH5
}

// RechargeCatePC PC充值
func RechargeCatePC() string {
	return rechargeCatePC
}

// FoodSnacks 食品小吃
func FoodSnacks() int {
	return foodSnacks
}

// Drink 饮料
func Drink() int {
	return drink
}

// FlowersAndGreenery 鲜花绿植
func FlowersAndGreenery() int {
	return flowersAndGreenery
}

// Other 其他
func Other() int {
	return other
}

// PrintingTicketing 文印票务
func PrintingTicketing() int {
	return printingTicketing
}

// ConvenienceStores 便利店
func ConvenienceStores() int {
	return convenienceStores
}

// FreshFruit 水果生鲜
func FreshFruit() int {
	return freshFruit
}

// IntraCityECommerce 同城电商
func IntraCityECommerce() int {
	return intraCityECommerce
}

// Medicine 医药
func Medicine() int {
	return medicine
}

// Cake 蛋糕
func Cake() int {
	return cake
}

// Wine 酒品
func Wine() int {
	return wine
}

// SmallCommodityMarkets 小商品市场
func SmallCommodityMarkets() int {
	return smallCommodityMarkets
}

// Clothing 服装
func Clothing() int {
	return clothing
}

// AutoRepairParts 汽修零配
func AutoRepairParts() int {
	return autoRepairParts
}

// DigitalAppliances 数码家电
func DigitalAppliances() int {
	return digitalAppliances
}

// CrayfishBBQ 小龙虾/烧烤
func CrayfishBBQ() int {
	return crayfishBBQ
}

// Supermarket 超市
func Supermarket() int {
	return supermarket
}

// ChafingDish 火锅
func ChafingDish() int {
	return chafingDish
}

// PersonalCareMakeup 个护美妆
func PersonalCareMakeup() int {
	return personalCareMakeup
}

// Mother 母婴
func Mother() int {
	return mother
}

// HomeTextiles 家居家纺
func HomeTextiles() int {
	return homeTextiles
}

// CellPhone 手机
func CellPhone() int {
	return cellPhone
}

// Home 家装
func Home() int {
	return home
}

// AdultProducts 成人用品
func AdultProducts() int {
	return adultProducts
}

// Campus 校园
func Campus() int {
	return campus
}

// HighEndMarket 高端市场
func HighEndMarket() int {
	return highEndMarket
}
