package sandpay

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/mahongran/sandpay/pay"
	"github.com/mahongran/sandpay/pay/elecaccountParams"
	"github.com/mahongran/sandpay/pay/elecaccountRequest"
	"github.com/mahongran/sandpay/util"
	"github.com/spf13/cast"
	"io"
	"log"
	"net/url"
	"sort"
	"strings"
	"time"
)

// CloudAccountPackage 云账户封装版
func (sandPay *SandPay) CloudAccountPackage(params elecaccountParams.CloudAccountPackage) (string, error) {
	var PayExtraMemberAccountOpening elecaccountRequest.PayExtraMemberAccountOpening
	PayExtraMemberAccountOpening.UserId = params.UserId
	PayExtraMemberAccountOpening.NickName = params.NickName
	PayExtraString, _ := json.Marshal(&PayExtraMemberAccountOpening)
	config := sandPay.Config
	body := elecaccountRequest.CloudAccountPackage{}
	body.Version = "10"
	body.MerNo = config.MerId
	body.CreateTime = time.Now().Format("20060102150405")
	body.MerOrderNo = params.OrderId
	body.OrderAmt = "0.11"
	body.NotifyUrl = params.NotifyUrl
	body.FrontUrl = params.FrontUrl
	body.CreateIp = strings.Replace(params.CreateIp, ".", "_", -1)
	body.PayExtra = string(PayExtraString)
	body.AccsplitFlag = "NO"
	body.SignType = "RSA"
	body.StoreId = "000000"
	body.ExpireTime = params.ExpireTime
	body.GoodsName = "开户"
	body.ProductCode = elecaccountParams.MemberAccountOpening
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
		"merch_extend_params",
	}
	dataMap := StructToMapString(body)
	sign, _ := pay.CloudAccountPackageSign(dataMap, keysToSign)
	dataMap["sign"] = sign
	//需要url decode 转码的key
	keysToUrlDecode := []string{"goods_name", "notify_url", "return_url", "pay_extra", "meta_option", "extend", "merch_extend_params", "sign"}
	u := config.CloudAccountApiHost + "?" + HttpBuildQuery(dataMap, keysToUrlDecode)
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
	body := elecaccountRequest.OneClickAccountOpening{
		Mid:             config.MerId,
		SignType:        "SHA1WithRSA",
		EncryptType:     "AES",
		Version:         "1.0.0",
		Timestamp:       time.Now().Format("2006-01-02 15:04:05"),
		CustomerOrderNo: params.CustomerOrderNo,
		BizUserNo:       params.BizUserNo,
		NickName:        params.NickName,
		Name:            params.Name,
		IdType:          params.IdType,
		IdNo:            params.IdNo,
		Mobile:          params.Mobile,
		NotifyUrl:       params.NotifyUrl,
		FrontUrl:        params.FrontUrl,
	}
	sanDe := util.SandAES{}
	key := sanDe.RandStr(16)
	dataMap := StructToMap(body)
	dataMap["data"], _ = FormData(dataMap, key)
	dataMap["encryptKey"], _ = pay.FormEncryptKey(key)
	sign, _ := pay.PrivateSha1SignData(dataMap["data"].(string))
	dataMap["sign"] = sign
	DataByte, _ := json.Marshal(dataMap)
	fmt.Println("请求参数:" + string(DataByte))
	resp, err := util.Do(config.CloudAccountEncapsulationHost+"/v4/elecaccount/ceas.elec.account.protocol.open", string(DataByte))
	if err != nil {
		return "", err
	}
	d := make(map[string]interface{})
	if err := json.Unmarshal(resp, &d); err != nil {
		return "", err
	}
	fmt.Println("杉德回调解析结果:" + string(resp))
	return string(resp), nil
}

func FormData(paraMap interface{}, keys string) (string, error) {

	plaintext, err := json.Marshal(paraMap)
	if err != nil {
		return "", err
	}
	log.Printf("data 加密前：%v", string(plaintext))
	// 16 位随机数作为 AES 密钥
	key := []byte(keys)

	// 创建一个 AES 块密码
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 创建一个随机的初始化向量 IV
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 创建一个 CTR 加密器
	stream := cipher.NewCTR(block, iv)

	// 加密数据
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, []byte(plaintext))

	// 将初始化向量和密文连接起来
	ciphertext = append(iv, ciphertext...)

	// Base64 编码加密结果
	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	log.Printf("data 加密后：%v", encoded)
	return encoded, nil
}
func StructToMapString(p interface{}) (list map[string]string) {
	data, err := json.Marshal(p)
	if err != nil {
		return list
	}
	// Unmarshal JSON into map[string]interface{}
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
	// Unmarshal JSON into map[string]interface{}
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return list
	}
	return m
}
