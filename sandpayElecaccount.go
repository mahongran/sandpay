package sandpay

import (
	"encoding/json"
	"github.com/mahongran/sandpay/pay"
	"github.com/mahongran/sandpay/pay/elecaccountParams"
	"github.com/mahongran/sandpay/pay/elecaccountRequest"
	"github.com/mahongran/sandpay/util"
	"github.com/spf13/cast"
	"log"
	"net/url"
	"sort"
	"strings"
	"time"
)

// WithdrawApply 提现申请接口
func (sandPay *SandPay) WithdrawApply(params elecaccountParams.WithdrawApplyParams) (string, error) {
	config := sandPay.Config
	var body elecaccountRequest.WithdrawApplyRequest
	body.Mid = config.MerId
	body.SignType = "SHA1WithRSA"
	body.EncryptType = "AES"
	body.Version = "1.0.0"
	body.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	body.CustomerOrderNo = params.CustomerOrderNo
	body.BizUserNo = params.BizUserNo
	body.AccountType = params.AccountType
	body.OrderAmt = params.OrderAmt
	body.UserFeeAmt = params.UserFeeAmt
	body.RelatedCardNo = params.RelatedCardNo
	body.NotifyUrl = params.NotifyUrl
	body.FrontUrl = params.FrontUrl
	DataByte := AddSignature(body)
	resp, err := util.Do(params.ApiHost+"/v4/electrans/ceas.elec.trans.withdraw.apply", DataByte)
	if err != nil {
		return "", err
	}
	d := make(map[string]interface{})
	if err := json.Unmarshal(resp, &d); err != nil {
		return "", err
	}
	j, err := pay.CloudAccountVerification(d)
	if err != nil {
		return "", err
	}
	return j, nil
}

// FundOperationConfirmation 资金操作确认
func (sandPay *SandPay) FundOperationConfirmation(params elecaccountParams.FundOperationConfirmationParams) (string, error) {
	config := sandPay.Config
	var body elecaccountRequest.FundOperationConfirmationRequest
	body.Mid = config.MerId
	body.SignType = "SHA1WithRSA"
	body.EncryptType = "AES"
	body.Version = "1.0.0"
	body.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	body.CustomerOrderNo = params.CustomerOrderNo
	body.BizUserNo = params.BizUserNo
	body.OriCustomerOrderNo = params.OriCustomerOrderNo
	body.OriOrderAmt = params.OriOrderAmt
	body.SmsCode = params.SmsCode
	BeforeEncryption, _ := json.Marshal(body)
	log.Printf("明文请求参数：%v", string(BeforeEncryption))
	DataByte := AddSignature(body)
	log.Printf("密文请求参数：%v", DataByte)
	resp, err := util.Do(params.ApiHost+"/v4/electrans/ceas.elec.trans.order.confirm", DataByte)
	if err != nil {
		return "", err
	}
	d := make(map[string]interface{})
	if err := json.Unmarshal(resp, &d); err != nil {
		return "", err
	}
	j, err := pay.CloudAccountVerification(d)
	if err != nil {
		return "", err
	}
	return j, nil
}

// BackendRechargeOrderPlacement 云账户后台充值下单
func (sandPay *SandPay) BackendRechargeOrderPlacement(params elecaccountParams.BackendRechargeOrderPlacementParams) (string, error) {
	config := sandPay.Config
	var body elecaccountRequest.BackendRechargeOrderPlacementRequest
	body.Mid = config.MerId
	body.SignType = "SHA1WithRSA"
	body.EncryptType = "AES"
	body.Version = "1.0.0"
	body.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	body.CustomerOrderNo = params.CustomerOrderNo
	body.BizUserNo = params.BizUserNo
	body.NotifyUrl = params.NotifyUrl
	body.FrontUrl = params.FrontUrl

	body.OrderTimeOut = params.OrderTimeOut
	body.PayTool = params.PayTool
	body.PayExtend = params.PayExtend
	body.WalletAmt = params.WalletAmt
	//!!!此参数不可传 传了资金操作确认接口会报错!!!
	//body.Extend = params.Extend
	DataByte := AddSignature(body)

	resp, err := util.Do(params.ApiHost+"/v4/electrans/ceas.elec.trans.third.payment.deposit", DataByte)
	if err != nil {
		return "", err
	}
	d := make(map[string]interface{})
	if err := json.Unmarshal(resp, &d); err != nil {
		return "", err
	}
	j, err := pay.CloudAccountVerification(d)
	if err != nil {
		return "", err
	}
	return j, nil
}

// BalanceQuery 查询用户余额
func (sandPay *SandPay) BalanceQuery(params elecaccountParams.BalanceQueryParams) (string, error) {
	config := sandPay.Config
	var body elecaccountRequest.BalanceQueryRequest
	body.Mid = config.MerId
	body.SignType = "SHA1WithRSA"
	body.EncryptType = "AES"
	body.Version = "1.0.0"
	body.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	body.CustomerOrderNo = params.CustomerOrderNo
	body.BizUserNo = params.BizUserNo
	body.AccountType = params.AccountType
	DataByte := AddSignature(body)
	resp, err := util.Do(params.ApiHost+"/v4/elecaccount/ceas.elec.account.balance.query", DataByte)
	if err != nil {
		return "", err
	}
	d := make(map[string]interface{})
	if err := json.Unmarshal(resp, &d); err != nil {
		return "", err
	}
	j, err := pay.CloudAccountVerification(d)
	if err != nil {
		return "", err
	}
	return j, nil
}

// PasswordManagement 密码管理
func (sandPay *SandPay) PasswordManagement(params elecaccountParams.PasswordManagementParams) (string, error) {
	config := sandPay.Config
	var body elecaccountRequest.PasswordManagementRequest
	body.Mid = config.MerId
	body.SignType = "SHA1WithRSA"
	body.EncryptType = "AES"
	body.Version = "1.0.0"
	body.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	body.CustomerOrderNo = params.CustomerOrderNo
	body.BizUserNo = params.BizUserNo
	body.PageType = params.PageType
	body.ManagementType = params.ManagementType
	body.Extend = params.Extend
	body.NotifyUrl = params.NotifyUrl
	body.FrontUrl = params.FrontUrl
	DataByte := AddSignature(body)
	resp, err := util.Do(params.ApiHost+"/v4/elecaccount/ceas.elec.account.pay.password.manage", DataByte)
	if err != nil {
		return "", err
	}
	d := make(map[string]interface{})
	if err := json.Unmarshal(resp, &d); err != nil {
		return "", err
	}
	j, err := pay.CloudAccountVerification(d)
	if err != nil {
		return "", err
	}
	return j, nil
}

// UnbindAssociatedCards 解绑关联卡
func (sandPay *SandPay) UnbindAssociatedCards(params elecaccountParams.UnbindAssociatedCardsParams) (string, error) {
	config := sandPay.Config
	var body elecaccountRequest.UnbindAssociatedCardsRequest
	body.Mid = config.MerId
	body.SignType = "SHA1WithRSA"
	body.EncryptType = "AES"
	body.Version = "1.0.0"
	body.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	body.CustomerOrderNo = params.CustomerOrderNo
	body.BizUserNo = params.BizUserNo
	body.RelatedCardNo = params.RelatedCardNo
	body.NotifyUrl = params.NotifyUrl
	DataByte := AddSignature(body)
	resp, err := util.Do(params.ApiHost+"/v4/elecaccount/ceas.elec.unbind.card", DataByte)
	if err != nil {
		return "", err
	}
	d := make(map[string]interface{})
	if err := json.Unmarshal(resp, &d); err != nil {
		return "", err
	}
	j, err := pay.CloudAccountVerification(d)
	if err != nil {
		return "", err
	}
	return j, nil
}

// SetAssociatedBankCardConfirm 设置关联银行卡确认
func (sandPay *SandPay) SetAssociatedBankCardConfirm(params elecaccountParams.SetAssociatedBankCardConfirmParams) (string, error) {
	config := sandPay.Config
	var body elecaccountRequest.SetAssociatedBankCardConfirmRequest
	body.Mid = config.MerId
	body.SignType = "SHA1WithRSA"
	body.EncryptType = "AES"
	body.Version = "1.0.0"
	body.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	body.CustomerOrderNo = params.CustomerOrderNo
	body.BizUserNo = params.BizUserNo
	body.OriCustomerOrderNo = params.OriCustomerOrderNo
	body.SmsCode = params.SmsCode
	body.NotifyUrl = params.NotifyUrl
	DataByte := AddSignature(body)
	resp, err := util.Do(params.ApiHost+"/v4/elecaccount/ceas.elec.bind.card.confirm", DataByte)
	if err != nil {
		return "", err
	}
	d := make(map[string]interface{})
	if err := json.Unmarshal(resp, &d); err != nil {
		return "", err
	}
	j, err := pay.CloudAccountVerification(d)
	if err != nil {
		return "", err
	}
	return j, nil
}

// SetAssociatedBankCard 设置关联银行卡
func (sandPay *SandPay) SetAssociatedBankCard(params elecaccountParams.SetAssociatedBankCardParams) (string, error) {
	config := sandPay.Config
	var body elecaccountRequest.SetAssociatedBankCardRequest
	body.Mid = config.MerId
	body.SignType = "SHA1WithRSA"
	body.EncryptType = "AES"
	body.Version = "1.0.0"
	body.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	body.CustomerOrderNo = params.CustomerOrderNo
	body.BizUserNo = params.BizUserNo
	body.CardNo = params.CardNo
	body.BankMobile = params.BankMobile
	body.RelatedCardType = params.RelatedCardType
	body.NotifyUrl = params.NotifyUrl
	body.FrontUrl = params.FrontUrl
	DataByte := AddSignature(body)
	resp, err := util.Do(params.ApiHost+"/v4/elecaccount/ceas.elec.bind.card", DataByte)
	if err != nil {
		return "", err
	}
	d := make(map[string]interface{})
	if err := json.Unmarshal(resp, &d); err != nil {
		return "", err
	}
	j, err := pay.CloudAccountVerification(d)
	if err != nil {
		return "", err
	}
	return j, nil
}

// AssociatedCardQuery 关联卡查询
func (sandPay *SandPay) AssociatedCardQuery(params elecaccountParams.AssociatedCardQueryParams) (string, error) {
	config := sandPay.Config
	var body elecaccountRequest.AssociatedCardQueryRequest
	body.Mid = config.MerId
	body.SignType = "SHA1WithRSA"
	body.EncryptType = "AES"
	body.Version = "1.0.0"
	body.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	body.CustomerOrderNo = params.CustomerOrderNo
	body.BizUserNo = params.BizUserNo
	body.RelatedCardNo = params.RelatedCardNo
	body.NotifyUrl = params.NotifyUrl
	body.FrontUrl = params.FrontUrl
	DataByte := AddSignature(body)
	resp, err := util.Do(params.ApiHost+"/v4/elecaccount/ceas.elec.bind.card.query", DataByte)
	if err != nil {
		return "", err
	}
	d := make(map[string]interface{})
	if err := json.Unmarshal(resp, &d); err != nil {
		return "", err
	}
	j, err := pay.CloudAccountVerification(d)
	if err != nil {
		return "", err
	}
	return j, nil
}

// IsSetPayPassword 是否设置支付密码
func (sandPay *SandPay) IsSetPayPassword(params elecaccountParams.IsSetPayPasswordParams) (string, error) {
	config := sandPay.Config
	var body elecaccountRequest.IsSetPayPasswordRequest
	body.Mid = config.MerId
	body.SignType = "SHA1WithRSA"
	body.EncryptType = "AES"
	body.Version = "1.0.0"
	body.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	body.CustomerOrderNo = params.CustomerOrderNo
	body.BizUserNo = params.BizUserNo
	body.NotifyUrl = params.NotifyUrl
	body.FrontUrl = params.FrontUrl
	DataByte := AddSignature(body)
	resp, err := util.Do(params.ApiHost+"/v4/elecaccount/ceas.elec.account.pay.password.query", DataByte)
	if err != nil {
		return "", err
	}
	d := make(map[string]interface{})
	if err := json.Unmarshal(resp, &d); err != nil {
		return "", err
	}
	j, err := pay.CloudAccountVerification(d)
	if err != nil {
		return "", err
	}
	return j, nil
}

// WithdrawalApplication 云账户提现申请
func (sandPay *SandPay) WithdrawalApplication(params elecaccountParams.WithdrawalApplicationParam) (string, error) {
	config := sandPay.Config
	var body elecaccountRequest.WithdrawalApplicationRequest
	body.Mid = config.MerId
	body.SignType = "SHA1WithRSA"
	body.EncryptType = "AES"
	body.Version = "1.0.0"
	body.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	body.CustomerOrderNo = params.CustomerOrderNo
	body.BizUserNo = params.BizUserNo
	body.NotifyUrl = params.NotifyUrl
	body.FrontUrl = params.FrontUrl
	body.AccountType = params.AccountType
	body.OrderAmt = params.OrderAmt
	body.RelatedCardNo = params.RelatedCardNo
	body.Remark = params.Remark

	DataByte := AddSignature(body)

	resp, err := util.Do(params.ApiHost+"/v4/elecaccount/ceas.elec.trans.withdraw.apply", DataByte)
	if err != nil {
		return "", err
	}

	d := make(map[string]interface{})
	if err := json.Unmarshal(resp, &d); err != nil {
		return "", err
	}
	j, err := pay.CloudAccountVerification(d)
	if err != nil {
		return "", err
	}
	return j, nil
}

// AgreementSigning 协议签约
func (sandPay *SandPay) AgreementSigning(params elecaccountParams.AgreementSigningParam) (string, error) {
	config := sandPay.Config
	var body elecaccountRequest.AgreementSigningRequest
	body.Mid = config.MerId
	body.SignType = "SHA1WithRSA"
	body.EncryptType = "AES"
	body.Version = "1.0.0"
	body.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	body.CustomerOrderNo = params.CustomerOrderNo
	body.BizUserNo = params.BizUserNo
	body.NotifyUrl = params.NotifyUrl
	body.FrontUrl = params.FrontUrl
	var SignProtocolParam elecaccountRequest.SignProtocolParam
	SignProtocolParam.ProtocolNo = "XY001"
	body.SignProtocol = SignProtocolParam
	DataByte := AddSignature(body)

	resp, err := util.Do(params.ApiHost+"/v4/elecaccount/ceas.elec.account.protocol.sign", DataByte)
	if err != nil {
		return "", err
	}

	d := make(map[string]interface{})
	if err := json.Unmarshal(resp, &d); err != nil {
		return "", err
	}
	j, err := pay.CloudAccountVerification(d)
	if err != nil {
		return "", err
	}
	return j, nil
}

// BindCardToOpenAnAccount 云账户开户&&绑卡
func (sandPay *SandPay) BindCardToOpenAnAccount(params elecaccountParams.BindCardToOpenAnAccountParam) (string, error) {
	config := sandPay.Config
	var body elecaccountRequest.OneClickAccountOpeningRequest

	body.Mid = config.MerId
	body.SignType = "SHA1WithRSA"
	body.EncryptType = "AES"
	body.Version = "1.0.0"
	body.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	body.CustomerOrderNo = params.CustomerOrderNo
	body.BizUserNo = params.BizUserNo
	body.NickName = params.NickName
	body.Name = params.Name
	body.IdType = params.IdType
	body.IdNo = params.IdNo
	body.Mobile = params.Mobile
	body.NotifyUrl = params.NotifyUrl
	body.FrontUrl = params.FrontUrl
	var BankInfoParam elecaccountRequest.BankInfoParam
	BankInfoParam.CardNo = params.CardNo

	body.BankInfo = BankInfoParam
	DataByte := AddSignature(body)

	resp, err := util.Do(params.ApiHost+"/v4/elecaccount/ceas.elec.account.quick.bindcard.open", DataByte)
	if err != nil {
		return "", err
	}

	d := make(map[string]interface{})
	if err := json.Unmarshal(resp, &d); err != nil {
		return "", err
	}
	j, err := pay.CloudAccountVerification(d)
	if err != nil {
		return "", err
	}
	return j, nil
}

// OneClickAccountOpening 云账户一键开户
func (sandPay *SandPay) OneClickAccountOpening(params elecaccountParams.OneClickAccountOpening) (string, error) {
	config := sandPay.Config
	var body elecaccountRequest.OneClickAccountOpening

	body.Mid = config.MerId
	body.SignType = "SHA1WithRSA"
	body.EncryptType = "AES"
	body.Version = "1.0.0"
	body.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	body.CustomerOrderNo = params.CustomerOrderNo
	body.BizUserNo = params.BizUserNo
	body.NickName = params.NickName
	body.Name = params.Name
	body.IdType = params.IdType
	body.IdNo = params.IdNo
	body.Mobile = params.Mobile
	body.NotifyUrl = params.NotifyUrl
	body.FrontUrl = params.FrontUrl
	var SignPageRequiredParam elecaccountRequest.SignPageRequiredParam
	SignPageRequiredParam.PwdRequired = true
	SignPageRequiredParam.IdFileRequired = false
	SignPageRequiredParam.BindCardRequired = true
	body.SignPageRequired = SignPageRequiredParam
	DataByte := AddSignature(body)

	resp, err := util.Do(params.ApiHost+"/v4/elecaccount/ceas.elec.account.protocol.open", DataByte)
	if err != nil {
		return "", err
	}

	d := make(map[string]interface{})
	if err := json.Unmarshal(resp, &d); err != nil {
		return "", err
	}
	j, err := pay.CloudAccountVerification(d)
	if err != nil {
		return "", err
	}
	return j, nil
}

// CloudAccountCancellationConfirm 云账户注销确认接口
func (sandPay *SandPay) CloudAccountCancellationConfirm(params elecaccountParams.CloudAccountCancellationConfirmParams) (string, error) {
	config := sandPay.Config
	var body elecaccountRequest.CloudAccountCancellationConfirmRequest
	body.Mid = config.MerId
	body.SignType = "SHA1WithRSA"
	body.EncryptType = "AES"
	body.Version = "1.0.0"
	body.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	body.CustomerOrderNo = params.CustomerOrderNo
	body.BizUserNo = params.BizUserNo
	body.ApiHost = params.ApiHost
	body.OriCustomerOrderNo = params.OriCustomerOrderNo
	body.SmsCode = params.SmsCode
	body.NotifyUrl = params.NotifyUrl
	body.FrontUrl = params.FrontUrl

	DataByte := AddSignature(body)
	log.Printf("请求参数：%v", DataByte)
	resp, err := util.Do(params.ApiHost+"/v4/elecaccount/ceas.elec.account.member.modify.confirm", DataByte)
	if err != nil {
		return "", err
	}
	log.Println(string(resp))
	d := make(map[string]interface{})
	if err := json.Unmarshal(resp, &d); err != nil {
		return "", err
	}
	j, err := pay.CloudAccountVerification(d)
	if err != nil {
		return "", err
	}
	return j, nil
}

// CloudAccountCancellation 云账户用户注销
func (sandPay *SandPay) CloudAccountCancellation(params elecaccountParams.CloudAccountCancellationParams) (string, error) {
	config := sandPay.Config
	var body elecaccountRequest.CloudAccountCancellationRequest
	body.Mid = config.MerId
	body.SignType = "SHA1WithRSA"
	body.EncryptType = "AES"
	body.Version = "1.0.0"
	body.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	body.CustomerOrderNo = params.CustomerOrderNo
	body.BizUserNo = params.BizUserNo
	body.ApiHost = params.ApiHost
	body.BizType = params.BizType
	body.NotifyUrl = params.NotifyUrl
	body.FrontUrl = params.FrontUrl
	body.Remark = params.Remark

	DataByte := AddSignature(body)
	log.Printf("请求参数：%v", DataByte)
	resp, err := util.Do(params.ApiHost+"/v4/elecaccount/ceas.elec.account.member.status.modify", DataByte)
	if err != nil {
		return "", err
	}
	log.Println(string(resp))
	d := make(map[string]interface{})
	if err := json.Unmarshal(resp, &d); err != nil {
		return "", err
	}
	j, err := pay.CloudAccountVerification(d)
	if err != nil {
		return "", err
	}
	return j, nil
}

// CloudAccountUserInfo 云账户用户信息详情
func (sandPay *SandPay) CloudAccountUserInfo(params elecaccountParams.CloudAccountUserInfoParams) (string, error) {
	config := sandPay.Config
	var body elecaccountRequest.CloudAccountUserInfoRequest
	body.Mid = config.MerId
	body.SignType = "SHA1WithRSA"
	body.EncryptType = "AES"
	body.Version = "1.0.0"
	body.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	body.CustomerOrderNo = params.CustomerOrderNo
	body.BizUserNo = params.BizUserNo

	DataByte := AddSignature(body)

	log.Printf("请求参数：%v", DataByte)
	resp, err := util.Do(params.ApiHost+"/v4/elecaccount/ceas.elec.member.status.query", DataByte)
	if err != nil {
		return "", err
	}
	log.Println(string(resp))
	d := make(map[string]interface{})
	if err := json.Unmarshal(resp, &d); err != nil {
		return "", err
	}
	j, err := pay.CloudAccountVerification(d)
	if err != nil {
		return "", err
	}
	return j, nil
}

// CloudAccountTransfer 转账（企业转个人）
func (sandPay *SandPay) CloudAccountTransfer(params elecaccountParams.CloudAccountTransferParams) (string, error) {
	config := sandPay.Config
	var body elecaccountRequest.CloudAccountTransferRequest
	body.Mid = config.MerId
	body.SignType = "SHA1WithRSA"
	body.EncryptType = "AES"
	body.Version = "1.0.0"
	body.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	body.CustomerOrderNo = params.CustomerOrderNo
	body.AccountType = params.AccountType
	body.OrderAmt = params.OrderAmt
	//收款方信息
	var PayeeJSONObject elecaccountRequest.PayeeJSONObject
	PayeeJSONObject.Name = params.Name
	PayeeJSONObject.BizUserNo = params.BizUserNo
	body.Payee = PayeeJSONObject

	body.Postscript = params.Postscript
	body.Remark = params.Remark
	body.NotifyUrl = params.NotifyUrl
	DataByte := AddSignature(body)

	log.Printf("请求参数：%v", DataByte)
	resp, err := util.Do(params.ApiHost+"/v4/electrans/ceas.elec.trans.corp.transfer", DataByte)
	if err != nil {
		return "", err
	}
	log.Println(string(resp))
	d := make(map[string]interface{})
	if err := json.Unmarshal(resp, &d); err != nil {
		return "", err
	}
	j, err := pay.CloudAccountVerification(d)
	if err != nil {
		return "", err
	}
	return j, nil
}

// CloudAccountPackage 云账户封装版
func (sandPay *SandPay) CloudAccountPackage(params elecaccountParams.CloudAccountPackage) (string, error) {
	config := sandPay.Config
	body := elecaccountRequest.CloudAccountPackage{}
	body.Version = "10"
	body.MerNo = config.MerId
	body.CreateTime = time.Now().Format("20060102150405")
	body.MerOrderNo = params.OrderId
	body.OrderAmt = params.OrderAmt
	body.NotifyUrl = params.NotifyUrl
	body.FrontUrl = params.FrontUrl
	body.CreateIp = strings.Replace(params.CreateIp, ".", "_", -1)
	body.PayExtra = params.PayExtra
	body.AccsplitFlag = "NO"
	body.SignType = "RSA"
	body.StoreId = "000000"
	body.ExpireTime = params.ExpireTime
	body.GoodsName = params.GoodsName
	body.ProductCode = params.ProductCode
	body.Extends = params.Extends
	body.ClearCycle = "3"
	body.JumpScheme = "sandcash://scpay"
	body.MetaOption = `[{"s":"Android","n":"","id":"","sc":""},{"s":"IOS","n":"","id":"","sc":""}]` //固定值
	//参与验签的字段
	keysToSign := []string{
		"version",
		"mer_no",
		"mer_order_no",
		"create_time",
		"order_amt",
		"notify_url",
		"return_url",
		"create_ip",
		"pay_extra",
		"accsplit_flag",
		"sign_type",
		"store_id",
		"activity_no",
		"benefit_amount",
		"extend",
		"merch_extend_params",
	}
	dataMap := StructToMapString(body)
	sign, _ := pay.CloudAccountPackageSign(dataMap, keysToSign)
	dataMap["sign"] = sign
	//需要url decode 转码的key
	keysToUrlDecode := []string{"goods_name", "notify_url", "return_url", "pay_extra", "meta_option", "extend", "merch_extend_params", "sign"}
	u := params.ApiHost + "?" + HttpBuildQuery(dataMap, keysToUrlDecode)
	return u, nil
}

// HttpBuildQuery url encode
func HttpBuildQuery(params map[string]string, keysToUrlEncode []string) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var parts []string
	for _, k := range keys {
		value := cast.ToString(params[k])
		if shouldUrlEncode(k, keysToUrlEncode) {
			value = url.QueryEscape(value)
		}
		part := url.QueryEscape(k) + "=" + value
		parts = append(parts, part)
	}

	return strings.Join(parts, "&")
}

func shouldUrlEncode(key string, keysToUrlEncode []string) bool {
	for _, k := range keysToUrlEncode {
		if key == k {
			return true
		}
	}
	return false
}

func StructToMapString(p interface{}) (list map[string]string) {
	data, err := json.Marshal(p)
	if err != nil {
		return list
	}
	var m map[string]string
	if err := json.Unmarshal(data, &m); err != nil {
		return list
	}
	return m
}
func StructToMap(p interface{}) (list map[string]interface{}) {
	data, err := json.Marshal(p)
	if err != nil {
		return list
	}
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return list
	}
	return m
}
func AddSignature(body interface{}) string {
	sanDe := util.SandAES{}
	key := sanDe.RandStr(16)
	dataMap := StructToMap(body)
	plaintext, _ := json.Marshal(dataMap)
	dataMap["data"], _ = sanDe.AesEcbPkcs5Padding(key, string(plaintext))
	dataMap["encryptKey"], _ = pay.FormEncryptKey(key)
	sign, _ := pay.PrivateSha1SignData(dataMap["data"].(string))
	dataMap["sign"] = sign

	lists := make(map[string]interface{})

	lists["mid"] = dataMap["mid"]
	lists["sign"] = dataMap["sign"]
	lists["timestamp"] = dataMap["timestamp"]
	lists["version"] = dataMap["version"]
	lists["customerOrderNo"] = dataMap["customerOrderNo"]
	lists["signType"] = dataMap["signType"]
	lists["encryptType"] = dataMap["encryptType"]
	lists["encryptKey"] = dataMap["encryptKey"]
	lists["data"] = dataMap["data"]
	DataByte, _ := json.Marshal(lists)
	return string(DataByte)
}
