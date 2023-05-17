package elecaccountRequest

type OneClickAccountOpening struct {
	Mid             string `json:"mid"`             //商户号
	Timestamp       string `json:"timestamp"`       //时间戳
	Version         string `json:"version"`         //版本
	CustomerOrderNo string `json:"customerOrderNo"` //商户号下每次请求的唯一流水号
	BizUserNo       string `json:"bizUserNo"`       //用户在商户系统中的唯一编号
	NickName        string `json:"nickName"`        //会员昵称
	Name            string `json:"name"`            //会员姓名
	IdType          string `json:"idType"`          //01：身份证
	IdNo            string `json:"idNo"`            //身份证号
	Mobile          string `json:"mobile"`          //会员手机号
	NotifyUrl       string `json:"notifyUrl"`       //异步通知地址
	FrontUrl        string `json:"frontUrl"`        //前台通知地址
	SignType        string `json:"signType"`        //签名方式
	EncryptType     string `json:"encryptType"`     //加密方式
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

}

// PayExtraMemberAccountOpening 支付扩展域用户开户
type PayExtraMemberAccountOpening struct {
	UserId   string `json:"userId"`
	NickName string `json:"nickName"`
}
