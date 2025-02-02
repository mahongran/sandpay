package request

type OrderPayBody struct {
	//账户金额
	AccountAmt string `json:"accountAmt"`
	//云账户用户唯一id
	MasterAccount string `json:"masterAccount"`
	//商户订单号
	OrderCode string `json:"orderCode"`
	//2. 订单金额
	TotalAmount string `json:"totalAmount"`
	//3. 订单标题
	Subject string `json:"subject"`
	//4. 订单描述
	Body string `json:"body"`
	//5. 订单超时时间
	TxnTimeOut string `json:"txnTimeOut,omitempty"`
	//支付模式
	PayMode string `json:"payMode,omitempty"`
	//支付方式列表
	PayModeList string `json:"payModeList,omitempty"`
	UserId      string `json:"userId,omitempty"`
	//	7. 支付扩展域  ANS0.1024 C 具体格式根据 payMode 确定,
	//PayExtra PayExtra `json:"payExtra"`
	PayExtra string `json:"payExtra,omitempty"`
	//	8. 客户端 IP
	ClientIp string `json:"clientIp,omitempty"`
	//9. 异步通知地址 notifyUrl ANS0.256 M \
	NotifyUrl string `json:"notifyUrl"`
	//10. 前台通知地址
	FrontUrl string `json:"frontUrl"`
	//ANS0.256 M 11. 商户门店编号
	StoreId string `json:"storeId"`
	//12. 商户终端编号
	TerminalId string `json:"terminalId"`
	//13. 操作员编号
	OperatorId string `json:"operatorId"`
	//14. 清算模式
	ClearCycle string `json:"clearCycle,omitempty"`
	//	分账信息
	RoyaltyInfo string `json:"royaltyInfo,omitempty"`
	//	16. 风控信息域
	RiskRateInfo string `json:"riskRateInfo"`
	//	17. 业务扩展参数
	BizExtendParams string `json:"bizExtendParams"`
	//	18. 商户扩展参数
	MerchExtendParams string `json:"merchExtendParams"`
	//19. 扩展域
	Extends string `json:"extend"`
	//支付工具
	PayTool string `json:"payTool,omitempty"`
	// 商户上送的订单时间
	OrderTime string `json:"orderTime,omitempty"`
	// 币种
	CurrencyCode string `json:"currencyCode,omitempty"`
	// 卡号
	CardNo string `json:"cardNo,omitempty"`
}
