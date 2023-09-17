package constant

const (
	UnSent      = 0 //发送中 pending
	Send        = 1 //发送成功
	SendErr     = 2 //发送失败
	UnCheck     = 3 //未审核
	CheckPass   = 4 //审核通过
	CheckRejrct = 5 //审核驳回
	TxUnVerify  = 6 //交易未确认

	OrderSet                       = 1  //1存
	OrderDraw                      = 2  //2取
	OrderSetServiceFee             = 3  //存服务费
	OrderDrawServiceFee            = 4  //取手续费
	OrderOutRedTransfer            = 5  //发红包
	OrderInRedTransfer             = 6  //收红包
	OrderSetTransferFee            = 7  //存fcc钱包服务费
	OrderDrawTransferFee           = 8  //取fcc钱包服务费
	OrderOutTransfer               = 9  //转出
	OrderInTransfer                = 10 //转入
	BlindBoxBuyOrder               = 11 //群应用-盲盒用户购买订单记录
	BlindBoxPrizesOrder            = 12 //群应用-盲盒奖励订单记录
	OrderBlindBoxSplit             = 13 //群应用-盲盒分成记录
	OrderGroupDraw                 = 14 //群钱包提现
	OrderChangeFrom                = 15 //兑换出记录
	OrderChangeTo                  = 16 //兑换入记录
	OrderGroupAuthFrozen           = 17 // 群认证保证金
	OrderGroupAuthUnfreeze         = 18 // 群认证保证金返回
	OrderActSignReward             = 19 // 签到小程序FCC奖励
	OrderSpeedCastBuyVideo         = 20 // 购买视频记录
	OrderSpeedCastPlatformEarnings = 21 // 平台分成记录
	OrderSpeedCastDraw             = 22 // 速播钱包提现
)

// 群认证状态
const (
	GroupNotAuth           = 0 // 未认证
	GroupAuthPass          = 1 // 审核通过，已认证
	GroupAuthINAudit       = 2 // 审核中
	GroupAuthReject        = 3 // 审核被驳回
	GroupAuthNotEnough     = 4 // 认证条件不满足
	GroupAuthCancelByAdmin = 5 // 认证被群管理员取消
)

// 群认证方式
const (
	GroupAuthWayApply = 1 // 申请
	GroupAuthWayAdmin = 2 // 后台管理员手动认证
)

const (
	FccCoinId  = 1
	FccChainId = 100
	ETHChainId = 1
)

// 平台流转数据类型
const (
	Deposit  = 1 //存入
	WithDraw = 2 // 取
	Pay      = 3 //支付
	Transfer = 4 //转账
	RedPack  = 5 //红包
)

// 签到小程序
const (
	SignTypeReward = 1 // 签到领FCC
	SignTypeHelp   = 2 // 助力签到

	AwardStatusUnshipped = 1 // 未发货
	AwardStatusShipped   = 2 //
)

// 活动代码
const (
	ActivitySign = 1 // 签到小程序
)
