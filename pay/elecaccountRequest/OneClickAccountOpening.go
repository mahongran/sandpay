package elecaccountRequest

// CloudAccountCommon 云账户公共请求参数
type CloudAccountCommon struct {
	Mid             string `json:"mid"`             //商户号
	Sign            string `json:"sign"`            //签名
	Timestamp       string `json:"timestamp"`       //时间戳
	Version         string `json:"version"`         //版本
	CustomerOrderNo string `json:"customerOrderNo"` //商户号下每次请求的唯一流水号
	SignType        string `json:"signType"`        //签名方式
	EncryptType     string `json:"encryptType"`     //加密方式
	EncryptKey      string `json:"encryptKey"`      //加密key
	Data            string `json:"data"`            //报文体
}

// AccountChangeDetailsQueryRequest 交易记录
type AccountChangeDetailsQueryRequest struct {
	CloudAccountCommon
	BizUserNo   string `json:"bizUserNo"`   //用户在商户系统中的唯一编号
	AccountType string `json:"accountType"` // 01：支付电子户
	BeginDate   string `json:"beginDate"`   // 起始日期至结束日期最大不超过7日
	EndDate     string `json:"endDate"`     // 起始日期至结束日期最大不超过7日
	IoFlag      string `json:"IoFlag"`      // 00：全部
	PageNo      string `json:"pageNo"`      // 必须从1开始
	PageSize    string `json:"pageSize"`    // 每页条数
}

// AccountOpeningInformationQueryRequest 开户信息查询
type AccountOpeningInformationQueryRequest struct {
	CloudAccountCommon
	BizUserNo string `json:"bizUserNo"` //用户在商户系统中的唯一编号
}

// WithdrawApplyRequest 提现申请
type WithdrawApplyRequest struct {
	CloudAccountCommon
	BizUserNo     string `json:"bizUserNo"`     //用户在商户系统中的唯一编号
	ApiHost       string `json:"apiHost"`       //请求api地址
	AccountType   string `json:"accountType"`   //账户类型
	OrderAmt      string `json:"orderAmt"`      //订单金额
	RelatedCardNo string `json:"relatedCardNo"` //关联卡编号
	UserFeeAmt    string `json:"userFeeAmt"`    //用户服务费
	NotifyUrl     string `json:"notifyUrl"`     //异步通知地址
	FrontUrl      string `json:"frontUrl"`      //前台通知地址
}

// FundOperationConfirmationRequest 收款资金确认
type FundOperationConfirmationRequest struct {
	CloudAccountCommon
	BizUserNo          string `json:"bizUserNo"`          //用户在商户系统中的唯一编号
	OriCustomerOrderNo string `json:"oriCustomerOrderNo"` //原交易订单号
	OriOrderAmt        string `json:"oriOrderAmt"`        //原订单金额
	SmsCode            string `json:"smsCode"`
}

// BackendRechargeOrderPlacementRequest 后台充值
type BackendRechargeOrderPlacementRequest struct {
	CloudAccountCommon
	BizUserNo string `json:"bizUserNo"` //用户在商户系统中的唯一编号
	NotifyUrl string `json:"notifyUrl"` //异步通知地址
	FrontUrl  string `json:"frontUrl"`  //前台通知地址

	OrderTimeOut string      `json:"orderTimeOut"` //订单超时时间 格式：yyyymmddHHmmss默认2小时
	PayTool      string      `json:"payTool"`      //充值方式
	PayExtend    interface{} `json:"payExtend"`    //充值扩展域
	WalletAmt    string      `json:"walletAmt"`    //充值金额
	Extend       string      `json:"extend"`       //附加数据
}

// BalanceQueryRequest 查询用户余额
type BalanceQueryRequest struct {
	CloudAccountCommon
	BizUserNo   string `json:"bizUserNo"`   //用户在商户系统中的唯一编号
	AccountType string `json:"accountType"` //01：支付电子户 02：权益账户 03：奖励金户
}

// PasswordManagementRequest 密码管理
type PasswordManagementRequest struct {
	CloudAccountCommon
	BizUserNo string `json:"bizUserNo"` //用户在商户系统中的唯一编号
	NotifyUrl string `json:"notifyUrl"` //异步通知地址
	FrontUrl  string `json:"frontUrl"`  //前台通知地址

	PageType       string `json:"pageType"`       //标准页面 01
	ManagementType string `json:"managementType"` // 01：设置/重置支付密码 02：修改支付密码 03:   重置会员手机号
	Extend         string `json:"extend"`         //扩展域
}

// UnbindAssociatedCardsRequest 解绑银行卡
type UnbindAssociatedCardsRequest struct {
	CloudAccountCommon
	BizUserNo     string `json:"bizUserNo"`     //用户在商户系统中的唯一编号
	RelatedCardNo string `json:"relatedCardNo"` //需要查询具体某张卡时上传
	NotifyUrl     string `json:"notifyUrl"`     //异步通知地址
}

// SetAssociatedBankCardConfirmRequest 绑定银行卡
type SetAssociatedBankCardConfirmRequest struct {
	CloudAccountCommon
	BizUserNo          string `json:"bizUserNo"`          //用户在商户系统中的唯一编号
	OriCustomerOrderNo string `json:"oriCustomerOrderNo"` //原单号
	SmsCode            string `json:"smsCode"`            //验证码
	NotifyUrl          string `json:"notifyUrl"`          //异步通知地址
}

// SetAssociatedBankCardRequest 关联卡查询
type SetAssociatedBankCardRequest struct {
	CloudAccountCommon
	BizUserNo       string `json:"bizUserNo"`       //用户在商户系统中的唯一编号
	CardNo          string `json:"cardNo"`          //卡号
	BankMobile      string `json:"bankMobile"`      //银行预留手机号
	RelatedCardType string `json:"relatedCardType"` // 01：提现卡（默认） 02:快捷充值+提现卡
	NotifyUrl       string `json:"notifyUrl"`       //异步通知地址
	FrontUrl        string `json:"frontUrl"`        //前台通知地址
}

// AssociatedCardQueryRequest 关联卡查询
type AssociatedCardQueryRequest struct {
	CloudAccountCommon
	BizUserNo     string `json:"bizUserNo"`     //用户在商户系统中的唯一编号
	RelatedCardNo string `json:"relatedCardNo"` //需要查询具体某张卡时上传
	NotifyUrl     string `json:"notifyUrl"`     //异步通知地址
	FrontUrl      string `json:"frontUrl"`      //前台通知地址
}

// IsSetPayPasswordRequest 是否设置支付密码
type IsSetPayPasswordRequest struct {
	CloudAccountCommon
	BizUserNo string `json:"bizUserNo"` //用户在商户系统中的唯一编号
	NotifyUrl string `json:"notifyUrl"` //异步通知地址
	FrontUrl  string `json:"frontUrl"`  //前台通知地址
}

// WithdrawalApplicationRequest 提现申请
type WithdrawalApplicationRequest struct {
	CloudAccountCommon
	BizUserNo     string `json:"bizUserNo"`     //用户在商户系统中的唯一编号
	AccountType   string `json:"accountType"`   //01：支付电子户 02：宝易付权益电子户
	OrderAmt      string `json:"idType"`        //提现金额
	RelatedCardNo string `json:"relatedCardNo"` //关联卡号id
	Remark        string `json:"remark"`        //备注
	NotifyUrl     string `json:"notifyUrl"`     //异步通知地址
	FrontUrl      string `json:"frontUrl"`      //前台通知地址
}
type SignPageRequiredParam struct {
	BindCardRequired bool `json:"bindCardRequired"` //绑卡设置
	IdFileRequired   bool `json:"idFileRequired"`   //上传身份证设置
	PwdRequired      bool `json:"pwdRequired"`      //密码设置
}
type BankInfoParam struct {
	CardNo string `json:"cardNo"`
}
type SignProtocolParam struct {
	ProtocolNo string `json:"protocolNo"` //XY001：开户协议 XY002：商户代扣协议
}

// AgreementSigningRequest 协议签约
type AgreementSigningRequest struct {
	CloudAccountCommon
	BizUserNo    string            `json:"bizUserNo"`    //用户在商户系统中的唯一编号
	NotifyUrl    string            `json:"notifyUrl"`    //异步通知地址
	FrontUrl     string            `json:"frontUrl"`     //前台通知地址
	SignProtocol SignProtocolParam `json:"signProtocol"` //银行信息域
}

// OneClickAccountOpeningRequest 个人会员绑卡并开户
type OneClickAccountOpeningRequest struct {
	CloudAccountCommon
	BizUserNo string        `json:"bizUserNo"` //用户在商户系统中的唯一编号
	NickName  string        `json:"nickName"`  //会员昵称
	Name      string        `json:"name"`      //会员姓名
	IdType    string        `json:"idType"`    //01：身份证
	IdNo      string        `json:"idNo"`      //身份证号
	Mobile    string        `json:"mobile"`    //会员手机号
	NotifyUrl string        `json:"notifyUrl"` //异步通知地址
	FrontUrl  string        `json:"frontUrl"`  //前台通知地址
	BankInfo  BankInfoParam `json:"bankInfo"`  //银行信息域
}

// OneClickAccountOpening 一键开户
type OneClickAccountOpening struct {
	CloudAccountCommon
	BizUserNo        string                `json:"bizUserNo"`        //用户在商户系统中的唯一编号
	NickName         string                `json:"nickName"`         //会员昵称
	Name             string                `json:"name"`             //会员姓名
	IdType           string                `json:"idType"`           //01：身份证
	IdNo             string                `json:"idNo"`             //身份证号
	Mobile           string                `json:"mobile"`           //会员手机号
	NotifyUrl        string                `json:"notifyUrl"`        //异步通知地址
	FrontUrl         string                `json:"frontUrl"`         //前台通知地址
	SignPageRequired SignPageRequiredParam `json:"signPageRequired"` //签约设置域
}

// CloudAccountCancellationConfirmRequest 销户确认参数参数
type CloudAccountCancellationConfirmRequest struct {
	CloudAccountCommon
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

// CloudAccountCancellationRequest 销户参数参数
type CloudAccountCancellationRequest struct {
	CloudAccountCommon
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

// CloudAccountUserInfoRequest 云账户详情参数
type CloudAccountUserInfoRequest struct {
	CloudAccountCommon
	BizUserNo string `json:"bizUserNo"`
}
type CloudAccountTransferRequest struct {
	CloudAccountCommon
	//账户类型01：支付电子户  02：宝易付权益电子户
	AccountType string `json:"accountType"`
	//转账金额 元
	OrderAmt string `json:"orderAmt"`
	//收款方
	Payee PayeeJSONObject `json:"payee"`
	//附言
	Postscript string `json:"postscript"`
	//备注
	Remark string `json:"remark"`
	//异步通知地址
	NotifyUrl string `json:"notifyUrl"`
}

// PayeeJSONObject 收款方信息
type PayeeJSONObject struct {
	BizUserNo string `json:"bizUserNo"`
	Name      string `json:"name"`
}

// CloudAccountPackage 云账户封装版
type CloudAccountPackage struct {
	Version      string `json:"version"`       //版本
	MerNo        string `json:"mer_no"`        //商户号
	CreateTime   string `json:"create_time"`   //时间戳
	MerOrderNo   string `json:"mer_order_no"`  //商户号下每次请求的唯一流水号
	OrderAmt     string `json:"order_amt"`     //订单金额
	NotifyUrl    string `json:"notify_url"`    //异步通知地址
	FrontUrl     string `json:"return_url"`    //前台通知地址
	CreateIp     string `json:"create_ip"`     //ip
	PayExtra     string `json:"pay_extra"`     //支付扩展域
	AccsplitFlag string `json:"accsplit_flag"` //分账标识 例：NO 无分账：NO；有分账：YES
	SignType     string `json:"sign_type"`     //签名类型，默认RSA
	StoreId      string `json:"store_id"`      //门店号 没有就填默认值 000000
	ExpireTime   string `json:"expire_time"`   //yyyyMMddHHmmss 例20180813142415，建议设置0.5～1小时
	GoodsName    string `json:"goods_name"`    //商品名称
	ProductCode  string `json:"product_code"`  //产品编码 https://www.yuque.com/sd_cw/xfq1vq/ut7292#jm1IE
	ClearCycle   string `json:"clear_cycle"`   //清算模式 3-D1;0-T1;1-T0;2-D0
	Sign         string `json:"sign"`          //MD5签名结果
	JumpScheme   string `json:"jump_scheme"`   //没有就填默认值  `sandcash://scpay`    此参数是安卓支付宝SDK跳转所需参数，如自定义，需要和客户端工程配置保持一致，例： android:scheme = "aaa"，android:host = "bbb"，jump_scheme 需填“aaa://bbb”。
	MetaOption   string `json:"meta_option"`   //[{"s":"Android","n":"","id":"","sc":""},{"s":"IOS","n":"","id":"","sc":""}] //固定值
	Extends      string `json:"extend"`
}

// PayExtraMemberAccountOpening 支付扩展域用户开户 || C2B
type PayExtraMemberAccountOpening struct {
	UserId   string `json:"userId"`
	NickName string `json:"nickName"`
}

// PayExtraOpeningC2C 支付扩展域  C2C
type PayExtraOpeningC2C struct {
	OperationType  string `json:"operationType"`  //1:转账申请
	RecvUserId     string `json:"recvUserId"`     //收款方会员编号
	Remark         string `json:"remark"`         //备注
	BizType        string `json:"bizType"`        //转账类型, //必填 1：转账确认模式 2：实时转账模式
	PayUserId      string `json:"payUserId"`      //付款方会员编号，用户在商户系统中的唯一编号
	UserFeeAmt     string `json:"userFeeAmt"`     //用户服务费，商户向用户收取的服务费  元
	Postscript     string `json:"postscript"`     //附言
	ReceiveTimeOut string `json:"receiveTimeOut"` //超时回退时间，bizType为1时生效，格式：yyyyMMddHHmmss（默认7天，最大支持7天）
}
