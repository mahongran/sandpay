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
	paramsJson, err := pay.Encrypt(StructToMap(body))
	if err != nil {
		return "", err
	}
	paramsJson["sign"], _ = pay.FormSign(paramsJson["data"].(string))
	paramsJson["signType"] = "SHA1WithRSA"
	//paramsJson, err = pay.Sign(paramsJson)
	//
	//if err != nil {
	//	return "", err
	//}
	DataByte, _ := json.Marshal(paramsJson)
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
