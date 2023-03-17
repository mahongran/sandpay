package pay

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/mahongran/sandpay/util"
	"io"
	"io/ioutil"
	"log"
	"net/url"
)

var certData *Cert

// 证书信息结构体
type Cert struct {
	// 私钥 签名使用
	Private *rsa.PrivateKey
	// 证书 与私钥为一套
	Cert *x509.Certificate
	// 签名证书ID
	CertId string
	// 加密证书
	EncryptCert *x509.Certificate
	// 公钥 加密验签使用
	Public *rsa.PublicKey
	// 加密公钥ID
	EncryptId string
}

//初始使用的配置
type Config struct {
	// 版本号 默认5.1.0
	Version string
	// 商户代码
	MerId string
	// 前台通知地址
	FrontUrl string
	// 验签私钥证书地址，传入pfx此路径可不传
	// openssl pkcs12 -in xxxx.pfx -nodes -out server.pem 生成为原生格式pem 私钥
	// openssl rsa -in server.pem -out server.key  生成为rsa格式私钥文件
	PrivatePath string
	// 验签证书地址,传入pfx此路径可以不传
	// openssl pkcs12 -in xxxx.pfx -clcerts -nokeys -out key.cert
	CertPath string
	// wind导出的加密证书地址
	EncryptCertPath string
	//API 网关地址
	ApiHost             string
	CloudAccountApiHost string
	NotifyUrl           string
}

func LoadCertInfo(info *Config) (err error) {
	certData = &Cert{}
	certData.EncryptCert, err = LoadPublicKey(info.EncryptCertPath)
	if err != nil {
		err = fmt.Errorf("encryptCert ERR:%v", err)
		return
	}
	certData.EncryptId = fmt.Sprintf("%v", certData.EncryptCert.SerialNumber)
	certData.Public = certData.EncryptCert.PublicKey.(*rsa.PublicKey)
	log.Println("	certData.Public", certData.Public)
	certData.Private, err = ParsePrivateFromFile(info.PrivatePath)
	log.Println("certData.Private", certData.Private)
	certData.Cert, err = ParseCertificateFromFile(info.CertPath)
	log.Println("certData.Cert", certData.Cert, err)
	certData.CertId = fmt.Sprintf("%v", certData.Cert.SerialNumber)
	fmt.Println("certData.CertId ", certData.CertId)
	return
}

// 根据文件名解析出证书
// openssl pkcs12 -in xxxx.pfx -clcerts -nokeys -out key.cert
func ParseCertificateFromFile(path string) (cert *x509.Certificate, err error) {
	// Read the verify sign certification key
	pemData, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	// Extract the PEM-encoded data block
	block, _ := pem.Decode(pemData)
	if block == nil {
		err = fmt.Errorf("bad key data: %s", "not PEM-encoded")
		return
	}
	if got, want := block.Type, "CERTIFICATE"; got != want {
		err = fmt.Errorf("unknown key type %q, want %q", got, want)
		return
	}

	// Decode the certification
	cert, err = x509.ParseCertificate(block.Bytes)
	if err != nil {
		err = fmt.Errorf("bad private key: %s", err)
		return
	}
	return
}

// 根据文件名解析出私钥 ,文件必须是rsa 私钥格式。
// openssl pkcs12 -in xxxx.pfx -nodes -out server.pem 生成为原生格式pem 私钥
// openssl rsa -in server.pem -out server.key  生成为rsa格式私钥文件
func ParsePrivateFromFile(path string) (private *rsa.PrivateKey, err error) {
	// Read the private key\
	//fmt.Println("path")
	pemData, err := ioutil.ReadFile(path)

	//pemData := generatePem(pemDataByte)
	if err != nil {
		err = fmt.Errorf("read key file: %s", err)
		return
	}

	// Extract the PEM-encoded data block
	block, _ := pem.Decode(pemData)
	if block == nil {
		err = fmt.Errorf("bad key data: %s", "not PEM-encoded")
		return
	}
	//fmt.Println("err unknown key type", block)
	if got, want := block.Type, "RSA PRIVATE KEY"; got != want {
		err = fmt.Errorf("unknown key type %q, want %q", got, want)
		return
	}

	// Decode the RSA private key
	private, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		err = fmt.Errorf("bad private key: %s", err)
		return
	}
	return
}

//格式转化
func ChunkSplit(body string, chunklen uint, end string) string {
	if end == "" {
		end = "\r\n"
	}
	runes, erunes := []rune(body), []rune(end)
	l := uint(len(runes))
	if l <= 1 || l < chunklen {
		return body + end
	}
	ns := make([]rune, 0, len(runes)+len(erunes))
	var i uint
	for i = 0; i < l; i += chunklen {
		if i+chunklen > l {
			ns = append(ns, runes[i:]...)
		} else {
			ns = append(ns, runes[i:i+chunklen]...)
		}
		ns = append(ns, erunes...)
	}
	return string(ns)
}

//将windows导出的私钥转化为pem格式
func generatePem(data []byte) string {
	base64 := Base64Encode(data)
	cert := ChunkSplit(base64, 64, "\n")
	cert = "-----BEGIN CERTIFICATE-----\n" + cert + "-----END CERTIFICATE-----\n"
	return cert
}

//加载公钥
func LoadPublicKey(path string) (cert *x509.Certificate, err error) {
	publicKeyData, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("read key file: %s", err)
		return
	}
	publicKeyDataBase := generatePem(publicKeyData)
	block, _ := pem.Decode([]byte(publicKeyDataBase))

	if block == nil {
		err = fmt.Errorf("bad key data: %s", "not PEM-encoded")
		return
	}
	if got, want := block.Type, "CERTIFICATE"; got != want {
		err = fmt.Errorf("unknown key type %q, want %q", got, want)
		return
	}

	//Decode the certification
	cert, err = x509.ParseCertificate(block.Bytes)
	if err != nil {
		err = fmt.Errorf("bad private key: %s", err)
		return
	}
	return
}
func NewPublicSha1Verify(signature, str string) (ok bool, err error) {

	//log.Println("vals", vals)

	hash := crypto.Hash.New(crypto.SHA1)
	hash.Write([]byte(str))
	hashed := hash.Sum(nil)

	var inSign []byte
	inSign, err1 := Base64Decode(signature)
	if err1 != nil {
		return false, fmt.Errorf("解析返回signature失败1 %v", err1)
	}
	err = rsa.VerifyPKCS1v15(certData.Public, crypto.SHA1, hashed, inSign)
	if err != nil {
		return false, fmt.Errorf("解析返回signature失败2 %v", err1)
	}
	return true, nil
}

// 返回数据验签
func PublicSha1Verify(vals url.Values) (res interface{}, err error) {
	var signature string
	var str string
	length := len(vals) - 1
	kvs := make(map[string]string, length)
	for k := range vals {
		if k == "sign" {
			signature = vals.Get(k)
			continue
		}
		if k == "data" {
			str = vals.Get(k)
		}
		if vals.Get(k) == "" {
			continue
		}
		kvs[k] = vals.Get(k)
	}

	//log.Println("vals", vals)

	hash := crypto.Hash.New(crypto.SHA1)
	hash.Write([]byte(str))
	hashed := hash.Sum(nil)

	var inSign []byte
	inSign, err1 := Base64Decode(signature)
	if err1 != nil {
		return nil, fmt.Errorf("解析返回signature失败 %v", err1)
	}

	//ffmt.P(hashed)
	err = rsa.VerifyPKCS1v15(certData.Public, crypto.SHA1, hashed, inSign)
	//log.Println("PublicSha1Verify  111 ", err)
	if err != nil {
		//log.Println("PublicSha1Verify Error from signing: %s ", err)
		//return "", err
	}
	return kvs, nil
}

// sign 做签
func PrivateSha1SignData(signData string) (string, error) {

	h := crypto.Hash.New(crypto.SHA1)
	h.Write([]byte(signData))
	hashed := h.Sum(nil)

	signer, err := rsa.SignPKCS1v15(rand.Reader, certData.Private,
		crypto.SHA1, hashed)
	if err != nil {
		fmt.Println("PrivateSha1SignData Error  from signing: %s\n", err)
		return "", err
	}
	return Base64Encode(signer), nil
}

// FormEncryptKey 云账户验签
func FormEncryptKey(key string) (string, error) {
	return util.RsaEncrypt(key, certData.Public)
}

// FormSign 云账户验签
func FormSign(data string) (string, error) {
	return util.SignSand(certData.Private, data)
}

func ValidateSign(data, signStr string) error {
	return util.Verification(data, signStr, certData.Public)
}
func Encrypt(paramsJson map[string]interface{}) (ResJson map[string]interface{}, err error) {
	aesKey, err := genRandomStringByLength(16)
	if err != nil {
		return ResJson, err
	}
	aesKeyBytes := []byte(aesKey)
	log.Printf("生成加密随机数：%v", aesKey)
	plainValue, err := json.Marshal(paramsJson)
	if err != nil {
		return ResJson, err
	}
	log.Printf("AES加密前值：%v", string(plainValue))
	encryptValueBytes, err := encryptAES(string(plainValue), aesKeyBytes)
	if err != nil {
		return ResJson, err
	}
	encryptValue := base64.StdEncoding.EncodeToString([]byte(encryptValueBytes))
	log.Printf("AES加密后值：%v", encryptValue)
	paramsJson["data"] = encryptValue
	encryptKeyBytes, err := encryptRSA(aesKeyBytes)
	if err != nil {
		return ResJson, err
	}
	sandEncryptKey := base64.StdEncoding.EncodeToString(encryptKeyBytes)
	paramsJson["encryptKey"] = sandEncryptKey
	paramsJson["encryptType"] = "AES"
	return paramsJson, nil
}

func genRandomStringByLength(length int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, length)
	_, err := io.ReadFull(rand.Reader, bytes)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}
func encryptAES(plainValue string, aesKeyBytes []byte) (string, error) {
	// 创建一个新的AES加密块
	block, err := aes.NewCipher(aesKeyBytes)
	if err != nil {
		return "", fmt.Errorf("无法创建AES加密块：%v", err)
	} // 对明文数据进行填充
	paddingLength := block.BlockSize() - len(plainValue)%block.BlockSize()
	padding := make([]byte, paddingLength)
	for i := range padding {
		padding[i] = byte(paddingLength)
	}
	paddedData := append([]byte(plainValue), padding...)

	// 使用AES加密算法进行加密
	encryptValue := make([]byte, len(paddedData))
	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCBCEncrypter(block, iv)
	stream.CryptBlocks(encryptValue, paddedData)

	// 对加密结果进行Base64编码
	encryptValueBase64 := base64.StdEncoding.EncodeToString(encryptValue)
	return encryptValueBase64, nil
}
func encryptRSA(aesKeyBytes []byte) ([]byte, error) {
	label := []byte("")
	hash := sha256.New()
	encryptKeyBytes, err := rsa.EncryptOAEP(hash, rand.Reader, certData.Public, aesKeyBytes, label)
	if err != nil {
		return nil, err
	}
	return encryptKeyBytes, nil
}
func Sign(paramsJson map[string]interface{}) (ResJson map[string]interface{}, err error) {
	// 获取需要签名的明文数据
	plainText, ok := paramsJson["data"].(string)
	if !ok {
		return ResJson, fmt.Errorf("data字段不是字符串类型")
	}
	// 打印待签名报文
	log.Printf("待签名报文：%s", plainText)

	// 使用SHA1WithRSA算法进行签名
	h := sha1.New()
	h.Write([]byte(plainText))
	hashed := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, certData.Private, crypto.SHA1, hashed)
	if err != nil {
		return ResJson, err
	}
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)

	// 将签名和签名算法添加到JSON对象中
	paramsJson["sign"] = signatureBase64
	paramsJson["signType"] = "SHA1WithRSA"
	return paramsJson, nil
}
