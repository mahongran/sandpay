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
	DataByte, _ := json.Marshal(dataMap)
	return string(DataByte)
}
