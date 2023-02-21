package sandpay

import (
	"bytes"
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
	"sort"
	"time"
)

func (sandPay *SandPay) ValidateSign(data, signStr string) error {
	return pay.ValidateSign(data, signStr)
}

// OneClickAccountOpening 云账户一键开户
func (sandPay *SandPay) OneClickAccountOpening(params elecaccountParams.OneClickAccountOpening) (string, error) {
	config := sandPay.Config
	timeString := time.Now().Format("20060102150405")
	body := elecaccountRequest.OneClickAccountOpening{
		Mid:             config.MerId,
		Version:         "1.0",
		SignType:        "SHA1WithRSA",
		EncryptType:     "AES",
		Timestamp:       timeString,
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
	body.Data, _ = AESEncrypt(StructToMap(body), []byte(key))
	body.EncryptKey, _ = pay.FormEncryptKey(key)
	body.Sign, _ = pay.FormSign(body.Data)
	DataByte, _ := json.Marshal(body)
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

func AESEncrypt(plainText map[string]interface{}, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	// Sort the plainText map by key
	var keys []string
	for k := range plainText {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	sortedPlainText := make(map[string]interface{})
	for _, k := range keys {
		sortedPlainText[k] = plainText[k]
	}

	plainTextBytes, err := json.Marshal(sortedPlainText)
	if err != nil {
		return "", err
	}

	blockSize := aes.BlockSize
	padding := blockSize - len(plainTextBytes)%blockSize
	paddedPlainTextBytes := append(plainTextBytes, bytes.Repeat([]byte{byte(padding)}, padding)...)

	result := make([]byte, aes.BlockSize+len(paddedPlainTextBytes))
	iv := result[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(result[aes.BlockSize:], paddedPlainTextBytes)

	return base64.StdEncoding.EncodeToString(result), nil

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
