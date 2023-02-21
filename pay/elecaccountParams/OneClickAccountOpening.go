package elecaccountParams

//OneClickAccountOpening 支付参数定义
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
