package elecaccountParams

const (
	C2BConsumption          = "04010001" //用户消费（C2B）
	C2CConsumption          = "04010003" //用户转账（C2C）
	C2CGuaranteeConsumption = "04010004" //担保消费  (C2C)
	MemberAccountOpening    = "00000001" //会员开户-协议签约
)

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
