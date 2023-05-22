package pay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/mahongran/sandpay/util"
	"io/ioutil"
	"net/url"
	"sort"
	"strings"
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
	// 加密证书 pro
	ProEncryptCert *x509.Certificate
	// pro 公钥 云账户账户侧加密验签使用
	ProPublic *rsa.PublicKey
	// 加密pro公钥ID
	ProEncryptId string
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
	// 杉德公钥
	EncryptCertPath string
	// 杉德pro公钥
	ProEncryptCertPath string
	//API 网关地址
	ApiHost string

	NotifyUrl string
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
	certData.ProEncryptCert, err = LoadPublicKey(info.ProEncryptCertPath)
	if err != nil {
		err = fmt.Errorf("ProEncryptCert ERR:%v", err)
		return
	}
	certData.ProEncryptId = fmt.Sprintf("%v", certData.ProEncryptCert.SerialNumber)
	certData.ProPublic = certData.ProEncryptCert.PublicKey.(*rsa.PublicKey)
	//log.Println("	certData.Public", certData.Public)
	certData.Private, err = ParsePrivateFromFile(info.PrivatePath)
	//log.Println("certData.Private", certData.Private)
	certData.Cert, err = ParseCertificateFromFile(info.CertPath)
	//log.Println("certData.Cert", certData.Cert, err)
	certData.CertId = fmt.Sprintf("%v", certData.Cert.SerialNumber)
	//fmt.Println("certData.CertId ", certData.CertId)
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
		//fmt.Println("PrivateSha1SignData Error  from signing: %s\n", err)
		return "", err
	}
	return Base64Encode(signer), nil
}
func shouldSign(key string, keysToUrlEncode []string) bool {
	for _, k := range keysToUrlEncode {
		if key == k {
			return true
		}
	}
	return false
}
func CloudAccountPackageSign(params map[string]string, keysToSign []string) (string, error) {
	// 1. 筛选并排序
	var keys []string
	for k := range params {
		if params[k] != "" && shouldSign(k, keysToSign) {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	// 2. 拼接
	var signStrings []string
	for _, k := range keys {
		signStrings = append(signStrings, fmt.Sprintf("%s=%s", k, params[k]))
	}
	signString := strings.Join(signStrings, "&")
	//fmt.Println("请求参数:" + string(signString))

	// 3. 调用签名函数
	hashed := sha1.Sum([]byte(signString))
	signature, err := rsa.SignPKCS1v15(rand.Reader, certData.Private, crypto.SHA1, hashed[:])
	if err != nil {
		return "", err
	}

	// 4. base64 编码签名结果
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)

	return signatureBase64, nil
}

// CloudAccountPublicSha1Verify 云账户回调验签
func CloudAccountPublicSha1Verify(signature, str string) (ok bool, err error) {

	//log.Println("vals", vals)

	hash := crypto.Hash.New(crypto.SHA1)
	hash.Write([]byte(str))
	hashed := hash.Sum(nil)

	var inSign []byte
	inSign, err1 := Base64Decode(signature)
	if err1 != nil {
		return false, fmt.Errorf("解析返回signature失败1 %v", err1)
	}
	err = rsa.VerifyPKCS1v15(certData.ProPublic, crypto.SHA1, hashed, inSign)
	if err != nil {
		return false, fmt.Errorf("解析返回signature失败2 %v", err1)
	}
	return true, nil
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

// FormEncryptKey 云账户加签
func FormEncryptKey(key string) (string, error) {
	return util.RsaEncrypt(key, certData.ProPublic)
}

// CloudAccountVerification 云账户验签
func CloudAccountVerification(d map[string]interface{}) (string, error) {
	data := d["data"].(string)
	sign := d["sign"].(string)
	encryptKey := d["encryptKey"].(string)
	// step8: 使用公钥验签报文
	ok, err := NewPublicSha1Verify1(data, sign, certData.ProPublic)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", errors.New("验签失败")
	}
	// step9: 使用私钥解密AESKey
	decryptAESKey, err := util.RsaDecrypt(encryptKey, certData.Private)
	if err != nil {
		return "", err
	}
	sanDe := util.SandAES{}
	//用key 解密 data 获得json
	jsonString, err := sanDe.AesEcbPkcs5PaddingDecrypt(decryptAESKey, data)
	if err != nil {
		return "", err
	}
	return jsonString, nil
}

// NewPublicSha1Verify 验签
func NewPublicSha1Verify1(signature, str string, SandPublicKey *rsa.PublicKey) (ok bool, err error) {
	hash := crypto.Hash.New(crypto.SHA1)
	hash.Write([]byte(str))
	hashed := hash.Sum(nil)
	var inSign []byte
	inSign, err1 := Base64Decode(signature)
	if err1 != nil {
		return false, fmt.Errorf("解析返回signature失败1 %v", err1)
	}
	err = rsa.VerifyPKCS1v15(SandPublicKey, crypto.SHA1, hashed, inSign)
	if err != nil {
		return false, fmt.Errorf("解析返回signature失败2" + err.Error())
	}
	return true, nil
}
