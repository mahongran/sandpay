package request

type OrderRefundBody struct {
	//商户订单号
	OrderCode string `json:"orderCode"`
	//2. 原商户订单号
	OriOrderCode string `json:"oriOrderCode"`
	//3. 退款金额
	RefundAmount string `json:"refundAmount"`
	//退账户金额 如产品编码为00002046，则必填
	RefundAccountAmt string `json:"refundAccountAmt"`
	//9. 异步通知地址 notifyUrl ANS0.256 M \
	NotifyUrl string `json:"notifyUrl"`
	//10. 退货原因
	RefundReason string `json:"refundReason"`
	//19. 扩展域
	Extends string `json:"extend"`
}
