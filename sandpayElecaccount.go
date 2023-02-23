package sandpay

import (
	"encoding/json"
	"fmt"
	"github.com/mahongran/sandpay/pay"
	"github.com/mahongran/sandpay/pay/elecaccountParams"
	"github.com/mahongran/sandpay/pay/elecaccountRequest"
	"github.com/mahongran/sandpay/util"
	"time"
)

// OneClickAccountOpening 云账户一键开户
func (sandPay *SandPay) OneClickAccountOpening(params elecaccountParams.OneClickAccountOpening) (string, error) {
	config := sandPay.Config
	timeString := time.Now().Format("20060102150405")
	body := elecaccountRequest.OneClickAccountOpening{
		Mid:             config.MerId,
		Version:         "1.0",
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
	dataMap := StructToMap(body)
	dataMap["signType"] = "SHA1WithRSA"
	dataMap["encryptType"] = "AES"
	dataMap["data"], _ = FormData(dataMap, key)
	dataMap["encryptKey"], _ = pay.FormEncryptKey(key)
	dataMap["sign"], _ = pay.FormSign(dataMap["data"].(string))

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

func FormData(paraMap interface{}, key string) (string, error) {

	dataJson, err := json.Marshal(paraMap)
	if err != nil {
		return "", err
	}
	aes := util.SandAES{}
	aes.Key = []byte(key)

	data := aes.Encypt5(dataJson)
	return data, nil
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
