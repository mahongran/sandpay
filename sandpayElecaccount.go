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
	"io"
	"log"
	"time"
)

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
		NotifyUrl:       sandPay.Config.NotifyUrl,
		FrontUrl:        sandPay.Config.FrontUrl,
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
	resp, err := util.Do(config.CloudAccountApiHost+"/v4/elecaccount/ceas.elec.account.protocol.open", string(DataByte))
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
