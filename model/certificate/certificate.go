package certificate

import "time"

type ThinkMerchantsInfo struct {
	ID                  uint   `gorm:"primary_key"`
	Company             string `gorm:"type:varchar(255);not null;default:'';comment:'商户名称'"`
	CompanyType         uint   `gorm:"company_type;comment:'企业类型'"`
	Portrait            string `gorm:"type:varchar(255);not null;default:'';comment:'商户头像'"`
	Contact             string `gorm:"type:varchar(255);not null;default:'';comment:'联系人'"`
	ContactPhone        string `gorm:"type:char(20);not null;default:'';comment:'联系人电话'"`
	CompanyInt          string `gorm:"type:varchar(255);not null;default:'';comment:'商户简介'"`
	MerID               int    `gorm:"not null"`
	UserName            string `gorm:"type:varchar(255);not null;default:'';comment:'商户名称'"`
	CreateTime          uint   `gorm:"not null;comment:'添加时间'"`
	EndTime             uint   `gorm:"not null;comment:'服务到期时间'"`
	BankUsername        string `gorm:"type:varchar(30);not null;default:'';comment:'银行卡持有人姓名'"`
	BankDeposit         string `gorm:"type:varchar(255);not null;default:'';comment:'开户行'"`
	BankNumbers         string `gorm:"type:varchar(30);not null;default:'';comment:'银行卡号'"`
	Address             string `gorm:"type:varchar(255);not null;default:'';comment:'钱包地址'"`
	AccountID           string `gorm:"type:varchar(20);comment:'商户子账户'"`
	IDCard              string `gorm:"type:varchar(20);comment:'身份证号'"`
	InterlineNumber     string `gorm:"type:varchar(100);comment:'联行号'"`
	AccountCodePhoto    string `gorm:"type:text;comment:'商户银行卡或者许可证照片信息'"`
	AccountCodeURLPhoto string `gorm:"type:text;comment:'商户银行卡或者许可证照片URL信息'"`
	Mobile              string `gorm:"type:varchar(20);comment:'手机号'"`
	BondMoney           int    `gorm:"default:0;comment:'商户保证金'"`
	AccountCodeDetail   string `gorm:"type:text;comment:'商户银行卡信息详情'"`
	ShopApplyID         int    `gorm:"default:0;comment:'商户申请id'"`
	MainBodyInfo        string `gorm:"type:text;comment:'主体信息资料'"`
	CorporationInfo     string `gorm:"type:text;comment:'法人信息资料'"`
	OperatorInfo        string `gorm:"type:text;comment:'经办人信息资料'"`
	AuditStatus         int    `gorm:"column:audit_status;comment:'审核状态'" json:"audit_status"`
}

func (ThinkMerchantsInfo) TableName() string {
	return "think_merchants_info"
}

type ThinkMember struct {
	ID             uint      `gorm:"primary_key;column:id;comment:'用户主键ID'"`
	ThirdMoney     float64   `gorm:"column:third_money;type:decimal(20,4);default:0.0000;comment:'三方账户金额'"`
	Money          float64   `gorm:"column:money;type:decimal(20,4);default:0.0000;comment:'金额（佣金）'"`
	Integral       float64   `gorm:"column:integral;type:decimal(12,2);not null;comment:'积分'"`
	Mobile         string    `gorm:"column:mobile;type:char(20);comment:'手机号'"`
	ParentID       uint      `gorm:"column:parent_id;default:0;comment:'推荐人ID'"`
	ParentIDs      string    `gorm:"column:parent_ids;type:text;comment:'上级用户id集合,用逗号分开'"`
	Password       string    `gorm:"column:password;type:char(64);default:'';comment:'用户密码'"`
	AccountAddress string    `gorm:"column:account_address;type:char(42);default:'';comment:'地址ID'"`
	MemberGroupID  uint8     `gorm:"column:member_group_id;default:0;comment:'会员分组ID'"`
	AppInitType    uint8     `gorm:"column:appinit_type;not null;default:1;comment:'应用初始化类型'"`
	SmsCode        string    `gorm:"column:sms_code;type:varchar(20);comment:'国家区号'"`
	Code           string    `gorm:"column:code;type:varchar(10);comment:'推荐码'"`
	InviteCount    int       `gorm:"column:invite_count;default:0;comment:'直推人数'"`
	TeamCount      int       `gorm:"column:team_count;not null;default:0;comment:'团队人数'"`
	ReceiptImg     string    `gorm:"column:receipt_img;type:varchar(200);comment:'收款二维码'"`
	Number         string    `gorm:"column:number;type:char(20);default:'';comment:'客户编号'"`
	BuyPasswd      string    `gorm:"column:buy_passwd;type:char(100);comment:'交易密码'"`
	SafetyCode     string    `gorm:"column:safety_code;type:char(64);default:'';comment:'安全密码'"`
	Nonce          string    `gorm:"column:nonce;type:varchar(10);default:'';comment:'盐'"`
	Nickname       string    `gorm:"column:nickname;type:varchar(30);default:'';comment:'用户昵称'"`
	Portrait       string    `gorm:"column:portrait;type:varchar(255);default:'';comment:'头像'"`
	VirtualMoney   float64   `gorm:"column:virtual_money;type:decimal(16,8);default:0.00000000;comment:'虚拟币'"`
	FreezeMoney    float64   `gorm:"column:freeze_money;type:decimal(16,8);default:0.00000000;comment:'冻结金额'"`
	Exp            uint      `gorm:"column:exp;not null;default:0;comment:'经验'"`
	Points         uint      `gorm:"column:points;not null;default:0;comment:'贡献值'"`
	RegisterIP     string    `gorm:"column:register_ip;type:char(15);default:'';comment:'注册IP'"`
	LoginIP        string    `gorm:"column:login_ip;type:char(15);default:'';comment:'最后登录IP'"`
	LoginTime      uint      `gorm:"column:login_time;not null;default:0;comment:'最后登录时间'"`
	LoginNum       uint      `gorm:"column:login_num;not null;default:1;comment:'登录次数'"`
	Status         uint8     `gorm:"column:status;not null;default:1;comment:'用户状态'"`
	Remark         string    `gorm:"column:remark;type:varchar(30);default:'';comment:'备注'"`
	ActiveTime     time.Time `gorm:"column:active_time;comment:'激活时间'"`
	ExpireTime     time.Time `gorm:"column:expire_time;comment:'超级VIP到期时间'"`
	CreateTime     time.Time `gorm:"column:create_time;comment:'创建时间'"`
	UpdateTime     time.Time `gorm:"column:update_time;comment:'更新时间'"`
	DeleteTime     string    `gorm:"column:delete_time;type:varchar(6);default:'0';comment:'删除时间'"`
	Gender         string    `gorm:"column:gender;type:varchar(10);default:'';comment:'性别'"`
	Ages           int       `gorm:"column:ages;default:0;comment:'年龄'"`
	OpenID         string    `gorm:"column:openid;type:varchar(255);default:'';comment:'第三方登录OpenID'"`
	SpecialID      uint      `gorm:"column:special_id;default:0;comment:'特殊ID'"`
	CID            int       `gorm:"column:cid;default:0;comment:'所属客户编号'"`
	IsUserWallet   int       `gorm:"column:is_user_wallet;default:0;comment:'是否用户钱包'"`
	QRCodeURL      string    `gorm:"column:qrcode_url;type:varchar(200);comment:'二维码地址'"`
	ShipLevel      int       `gorm:"column:ship_level;default:0;comment:'船级社等级'"`
	IdCardType     string    `gorm:"type:varchar(255);not null;default:'';comment:'身份证类型'"`
	IdCard         string    `gorm:"type:varchar(255);not null;default:'';comment:'身份证号码'"`
	DueDate        string    `gorm:"type:varchar(255);not null;default:'';comment:'身份证过期时间'"`
	Address        string    `gorm:"type:varchar(255);not null;default:'';comment:'详细地址'"`
	Name           string    `gorm:"type:varchar(255);not null;default:'';comment:'姓名'"`
	IdCardImg      string    `gorm:"type:text;comment:'身份证图片地址'"`
	AuditStatus    int       `gorm:"column:audit_status;comment:'审核状态'" json:"audit_status"`
}

func (ThinkMember) TableName() string {
	return "think_member"
}

type ThinkGoods struct {
	ID                         uint      `gorm:"primaryKey"`
	SpecialID                  int       `gorm:"not null;default:0;comment:'专题id'"`
	SpecialBlindBoxID          string    `gorm:"type:varchar(255);comment:'盲盒藏品表ID（盲盒）'"`
	Stand                      uint8     `gorm:"not null;default:1;comment:'上下架状态：1上架 ；0下架'"`
	Title                      string    `gorm:"type:varchar(255);comment:'商品名称'"`
	Description                string    `gorm:"type:varchar(255);comment:'商品描述'"`
	Images                     string    `gorm:"type:text;comment:'商品图片'"`
	Thumb                      string    `gorm:"type:varchar(255);comment:'缩略图'"`
	Code                       int       `gorm:"comment:'编号'"`
	Type                       int       `gorm:"default:1;comment:'类型 1藏品 2盲盒未拆 3盲盒已拆'"`
	MemberID                   int       `gorm:"not null;default:0;comment:'当前持有用户，0为平台持有，-1合成消耗'"`
	MerID                      int       `gorm:"default:null"`
	Status                     uint8     `gorm:"not null;default:0;comment:'状态: 0持有 1正在出售 2被买家锁定  6合成消耗,7藏品未拆'"`
	SellStatus                 uint8     `gorm:"default:0;comment:'转售状态 1是0否'"`
	SellStatusType             uint8     `gorm:"default:1;comment:'1.普通寄售 2竞拍寄售'"`
	BidsPrice                  float64   `gorm:"type:decimal(10,2);default:0.00;comment:'竞拍价格'"`
	Price                      float64   `gorm:"not null;type:decimal(10,2);default:0.00;comment:'当前价格'"`
	SaleCount                  int       `gorm:"default:0;comment:'销量/转售次数'"`
	TrilateralPlatformFee      float64   `gorm:"type:decimal(12,2);default:null;comment:'三方平台结算服务费'"`
	AisleFee                   float64   `gorm:"type:decimal(12,2);default:null;comment:'支付通道费'"`
	ChainFee                   float64   `gorm:"type:decimal(12,2);default:null;comment:'链上手续费'"`
	PlatformFee                float64   `gorm:"type:decimal(12,2);default:null;comment:'平台手续费'"`
	ChuangRateFee              float64   `gorm:"type:decimal(12,2);default:null;comment:'创作者版权分成'"`
	TrilateralPlatformTransfer float64   `gorm:"type:decimal(12,2);default:null;comment:'三方平台转账手续费'"`
	ChainMark                  int       `gorm:"default:0;comment:'链上标示'"`
	OwnerCode                  string    `gorm:"type:varchar(20);default:null;comment:'所有人ID'"`
	ContractAddress            string    `gorm:"type:varchar(255);default:null;comment:'合约地址'"`
	Hash                       string    `gorm:"type:varchar(255);default:null;comment:'哈希'"`
	StandTime                  time.Time `gorm:"default:null;comment:'转售时间'"`
	UpdateTime                 time.Time `gorm:"default:null;comment:'修改时间'"`
	CreateTime                 time.Time `gorm:"not null;comment:'创建时间'"`
	DdcURI                     string    `gorm:"type:varchar(255);default:null;comment:'ddcuri'"`
	TxHashStatus               uint8     `gorm:"default:0;comment:'0初始值  1铸造广播中 2铸造成功'"`
	IsReclaim                  uint8     `gorm:"default:0;comment:'是否回收（0否，1是）'"`
}

func (ThinkGoods) TableName() string {
	return "think_goods"
}
