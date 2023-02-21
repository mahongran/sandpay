package elecaccountRequest

type OneClickAccountOpening struct {
	Mid             string `json:"mid"`             //商户号
	Timestamp       string `json:"timestamp"`       //时间戳
	Version         string `json:"version"`         //版本
	SignType        string `json:"signType"`        //标志类型
	EncryptType     string `json:"encryptType"`     //加密类型
	CustomerOrderNo string `json:"customerOrderNo"` //商户号下每次请求的唯一流水号
	BizUserNo       string `json:"bizUserNo"`       //用户在商户系统中的唯一编号
	NickName        string `json:"nickName"`        //会员昵称
	Name            string `json:"name"`            //会员姓名
	IdType          string `json:"idType"`          //01：身份证
	IdNo            string `json:"idNo"`            //身份证号
	Mobile          string `json:"mobile"`          //会员手机号
	NotifyUrl       string `json:"notifyUrl"`       //异步通知地址
	FrontUrl        string `json:"frontUrl"`        //前台通知地址
	EncryptKey      string `json:"encryptKey"`
	Sign            string `json:"sign"`
	Data            string `json:"data"`
}
