package elecaccountParams

const (
	C2BConsumption          = "04010001" //用户消费（C2B）
	C2CConsumption          = "04010003" //用户转账（C2C）
	C2CGuaranteeConsumption = "04010004" //担保消费  (C2C)
	MemberAccountOpening    = "00000001" //会员开户-协议签约
)

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
}

// CloudAccountPackage 云账户封装版请求参数
type CloudAccountPackage struct {
	//商户号下每次请求的唯一流水号
	OrderId string `json:"orderId"`
	//用户在商户系统中的唯一编号
	UserId string `json:"userId"`
	//会员昵称
	NickName string `json:"nickName"`
	//yyyyMMddHHmmss 例20180813142415，建议设置0.5～1小时
	ExpireTime string `json:"expireTime"`
	//异步通知地址
	NotifyUrl string `json:"notifyUrl"`
	//前台通知地址
	FrontUrl string `json:"frontUrl"`
	//ip
	CreateIp string `json:"createIp"`
}
