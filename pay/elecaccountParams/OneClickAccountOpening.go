package elecaccountParams

const (
	C2BConsumption          = "04010001" //用户消费（C2B）
	C2CConsumption          = "04010003" //用户转账（C2C）
	C2CGuaranteeConsumption = "04010004" //担保消费  (C2C)
	MemberAccountOpening    = "00000001" //会员开户-协议签约
)

// FundOperationConfirmationParams 收款资金确认
type FundOperationConfirmationParams struct {
	CustomerOrderNo    string `json:"customerOrderNo"`    //商户号下每次请求的唯一流水号
	BizUserNo          string `json:"bizUserNo"`          //用户在商户系统中的唯一编号
	OriCustomerOrderNo string `json:"oriCustomerOrderNo"` //原交易订单号
	OriOrderAmt        string `json:"oriOrderAmt"`        //原订单金额
	SmsCode            string `json:"smsCode"`            //验证码
	ApiHost            string `json:"apiHost"`            //请求api地址
}

// PayExtendQuickPay 充值扩展域 快捷充值
type PayExtendQuickPay struct {
	RelatedCardNo string `json:"relatedCardNo"` //关联卡号
}

// BackendRechargeOrderPlacementParams 后台充值下单
// WX_JSBRIDGE 微信公众号支付
// WX_JSAPI 微信小程序支付
// QUICKPAY快捷充值
// UNION_PAY 银联SDK
// ALI_JSAPI 支付宝小程序支付
// ALI_JSBRIDGE 支付宝生活号支付
// UNION_PAY_H5 银联H5快捷
type BackendRechargeOrderPlacementParams struct {
	//商户号下每次请求的唯一流水号
	CustomerOrderNo string `json:"customerOrderNo"`
	//用户在商户系统中的唯一编号
	BizUserNo string `json:"bizUserNo"`
	//请求api地址
	ApiHost   string `json:"apiHost"`
	NotifyUrl string `json:"notifyUrl"`
	FrontUrl  string `json:"frontUrl"` //前台通知地址

	OrderTimeOut string      `json:"orderTimeOut"` //订单超时时间 格式：yyyymmddHHmmss默认2小时
	PayTool      string      `json:"payTool"`      //充值方式
	PayExtend    interface{} `json:"payExtend"`    //充值扩展域
	WalletAmt    string      `json:"walletAmt"`    //充值金额
	Extend       string      `json:"extend"`       //附加数据
}

// BalanceQueryParams 查询用户余额
type BalanceQueryParams struct {
	//商户号下每次请求的唯一流水号
	CustomerOrderNo string `json:"customerOrderNo"`
	//用户在商户系统中的唯一编号
	BizUserNo string `json:"bizUserNo"`
	//请求api地址
	ApiHost     string `json:"apiHost"`
	AccountType string `json:"accountType"` //01：支付电子户 02：权益账户 03：奖励金户
}

// PasswordManagementParams 密码管理
type PasswordManagementParams struct {
	//商户号下每次请求的唯一流水号
	CustomerOrderNo string `json:"customerOrderNo"`
	//用户在商户系统中的唯一编号
	BizUserNo string `json:"bizUserNo"`
	//请求api地址
	ApiHost   string `json:"apiHost"`
	NotifyUrl string `json:"notifyUrl"`
	FrontUrl  string `json:"frontUrl"` //前台通知地址

	PageType       string `json:"pageType"`       //标准页面 01
	ManagementType string `json:"managementType"` // 01：设置/重置支付密码 02：修改支付密码 03:   重置会员手机号
	Extend         string `json:"extend"`         //扩展域
}

// UnbindAssociatedCardsParams 解绑银行卡
type UnbindAssociatedCardsParams struct {
	//商户号下每次请求的唯一流水号
	CustomerOrderNo string `json:"customerOrderNo"`
	//用户在商户系统中的唯一编号
	BizUserNo string `json:"bizUserNo"`
	//请求api地址
	ApiHost       string `json:"apiHost"`
	RelatedCardNo string `json:"relatedCardNo"` //需要查询具体某张卡时上传
	NotifyUrl     string `json:"notifyUrl"`
}

// SetAssociatedBankCardConfirmParams 绑定银行卡
type SetAssociatedBankCardConfirmParams struct {
	//商户号下每次请求的唯一流水号
	CustomerOrderNo string `json:"customerOrderNo"`
	//用户在商户系统中的唯一编号
	BizUserNo string `json:"bizUserNo"`
	//请求api地址
	ApiHost            string `json:"apiHost"`
	OriCustomerOrderNo string `json:"oriCustomerOrderNo"` //原单号
	SmsCode            string `json:"smsCode"`            //验证码
	NotifyUrl          string `json:"notifyUrl"`
}

// SetAssociatedBankCardParams 设置关联银行卡
type SetAssociatedBankCardParams struct {
	//商户号下每次请求的唯一流水号
	CustomerOrderNo string `json:"customerOrderNo"`
	//用户在商户系统中的唯一编号
	BizUserNo string `json:"bizUserNo"`
	//请求api地址
	ApiHost         string `json:"apiHost"`
	CardNo          string `json:"cardNo"`          //卡号
	BankMobile      string `json:"bankMobile"`      //银行预留手机号
	RelatedCardType string `json:"relatedCardType"` // 01：提现卡（默认） 02:快捷充值+提现卡
	NotifyUrl       string `json:"notifyUrl"`
	FrontUrl        string `json:"frontUrl"`
}

// AssociatedCardQueryParams 关联卡查询
type AssociatedCardQueryParams struct {
	//商户号下每次请求的唯一流水号
	CustomerOrderNo string `json:"customerOrderNo"`
	//用户在商户系统中的唯一编号
	BizUserNo string `json:"bizUserNo"`
	//请求api地址
	ApiHost       string `json:"apiHost"`
	RelatedCardNo string `json:"relatedCardNo"` //需要查询具体某张卡时上传
	NotifyUrl     string `json:"notifyUrl"`
	FrontUrl      string `json:"frontUrl"`
}

// IsSetPayPasswordParams 是否设置支付密码请求体
type IsSetPayPasswordParams struct {
	//商户号下每次请求的唯一流水号
	CustomerOrderNo string `json:"customerOrderNo"`
	//用户在商户系统中的唯一编号
	BizUserNo string `json:"bizUserNo"`
	//请求api地址
	ApiHost string `json:"apiHost"`

	NotifyUrl string `json:"notifyUrl"`
	FrontUrl  string `json:"frontUrl"`
}

// CloudAccountCancellationConfirmParams 销户确认接口
type CloudAccountCancellationConfirmParams struct {
	//商户号下每次请求的唯一流水号
	CustomerOrderNo string `json:"customerOrderNo"`
	//用户在商户系统中的唯一编号
	BizUserNo string `json:"bizUserNo"`
	//请求api地址
	ApiHost string `json:"apiHost"`
	//原交易单号 请求会员状态管理所用的单号
	OriCustomerOrderNo string `json:"oriCustomerOrderNo"`
	//短信验证码 请求会员状态管理销户会下发给用户
	SmsCode   string `json:"smsCode"`
	NotifyUrl string `json:"notifyUrl"`
	FrontUrl  string `json:"frontUrl"`
}
type CloudAccountCancellationParams struct {
	//商户号下每次请求的唯一流水号
	CustomerOrderNo string `json:"customerOrderNo"`
	//用户在商户系统中的唯一编号
	BizUserNo string `json:"bizUserNo"`
	//请求api地址
	ApiHost string `json:"apiHost"`
	//操作类型 CLOSE－销户
	BizType string `json:"bizType"`
	//回调地址
	NotifyUrl string `json:"notifyUrl"`
	FrontUrl  string `json:"frontUrl"`
	Remark    string `json:"remark"`
}

type CloudAccountUserInfoParams struct {
	//商户号下每次请求的唯一流水号
	CustomerOrderNo string `json:"customerOrderNo"`
	//用户在商户系统中的唯一编号
	BizUserNo string `json:"bizUserNo"`
	//请求api地址
	ApiHost string `json:"apiHost"`
}

//CloudAccountTransferParams 云账户转账（企业转个人）
type CloudAccountTransferParams struct {
	//商户号下每次请求的唯一流水号
	CustomerOrderNo string `json:"customerOrderNo"`
	//账户类型01：支付电子户  02：宝易付权益电子户
	AccountType string `json:"accountType"`
	//转账金额 元
	OrderAmt string `json:"orderAmt"`
	//收款方 用户在商户系统中的唯一编号
	BizUserNo string `json:"bizUserNo"`
	//收款方姓名
	Name string `json:"string"`
	//附言
	Postscript string `json:"postscript"`
	//备注
	Remark string `json:"remark"`
	//请求api地址
	ApiHost string `json:"apiHost"`
	//异步通知地址
	NotifyUrl string `json:"notifyUrl"`
}

// AgreementSigningParam 协议签约
type AgreementSigningParam struct {
	//商户号下每次请求的唯一流水号
	CustomerOrderNo string `json:"customerOrderNo"`
	//用户在商户系统中的唯一编号
	BizUserNo string `json:"bizUserNo"`
	//异步通知地址
	NotifyUrl string `json:"notifyUrl"`
	//前台通知地址
	FrontUrl string `json:"frontUrl"`
	//请求api地址
	ApiHost string `json:"apiHost"`
}

// BindCardToOpenAnAccountParam 云账户开户&&绑卡
type BindCardToOpenAnAccountParam struct {
	//商户号下每次请求的唯一流水号
	CustomerOrderNo string `json:"customerOrderNo"`
	//用户在商户系统中的唯一编号
	BizUserNo string `json:"bizUserNo"`
	//会员昵称
	NickName string `json:"nickName"`
	//会员姓名
	Name string `json:"name"`
	//01：身份证
	IdType string `json:"idType"`
	//身份证号
	IdNo string `json:"idNo"`
	//卡号
	CardNo string `json:"cardNo"`
	//会员手机号
	Mobile string `json:"mobile"`
	//异步通知地址
	NotifyUrl string `json:"notifyUrl"`
	//前台通知地址
	FrontUrl string `json:"frontUrl"`
	//请求api地址
	ApiHost string `json:"apiHost"`
}

//OneClickAccountOpening 云账户支付参数定义
type OneClickAccountOpening struct {
	//商户号下每次请求的唯一流水号
	CustomerOrderNo string `json:"customerOrderNo"`
	//用户在商户系统中的唯一编号
	BizUserNo string `json:"bizUserNo"`
	//会员昵称
	NickName string `json:"nickName"`
	//会员姓名
	Name string `json:"name"`
	//01：身份证
	IdType string `json:"idType"`
	//身份证号
	IdNo string `json:"idNo"`
	//会员手机号
	Mobile string `json:"mobile"`
	//异步通知地址
	NotifyUrl string `json:"notifyUrl"`
	//前台通知地址
	FrontUrl string `json:"frontUrl"`
	//请求api地址
	ApiHost string `json:"apiHost"`
}

// CloudAccountPackage 云账户封装版请求参数
type CloudAccountPackage struct {
	//商户号下每次请求的唯一流水号
	OrderId string `json:"orderId"`
	//yyyyMMddHHmmss 例20180813142415，建议设置0.5～1小时
	ExpireTime string `json:"expireTime"`
	//异步通知地址
	NotifyUrl string `json:"notifyUrl"`
	//请求api地址
	ApiHost string `json:"apiHost"`
	//前台通知地址
	FrontUrl string `json:"frontUrl"`
	//ip
	CreateIp string `json:"createIp"`
	//订单金额
	OrderAmt string `json:"order_amt"`
	//商品名称
	GoodsName string `json:"goods_name"`
	//产品编码 https://www.yuque.com/sd_cw/xfq1vq/ut7292#jm1IE
	ProductCode string `json:"product_code"`
	//支付扩展域
	PayExtra string `json:"pay_extra"`
	//扩展域
	Extends string `json:"extends"`
}

//WithdrawalApplicationParam 云账户提现申请参数定义
type WithdrawalApplicationParam struct {
	CustomerOrderNo string `json:"customerOrderNo"` //商户号下每次请求的唯一流水号
	BizUserNo       string `json:"bizUserNo"`       //用户在商户系统中的唯一编号
	AccountType     string `json:"accountType"`     //01：支付电子户 02：宝易付权益电子户
	OrderAmt        string `json:"idType"`          //提现金额
	RelatedCardNo   string `json:"relatedCardNo"`   //关联卡号id
	Remark          string `json:"remark"`          //备注
	NotifyUrl       string `json:"notifyUrl"`       //异步通知地址
	FrontUrl        string `json:"frontUrl"`        //前台通知地址
	ApiHost         string `json:"apiHost"`         //请求api地址
}
