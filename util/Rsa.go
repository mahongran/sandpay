package util

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"log"
)

func RsaEncrypt(value string, rsaPublicKey *rsa.PublicKey) (string, error) {
	log.Printf("RSA 加密前：%v", value)
	buffer, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, []byte(value))
	if err != nil {
		return "", err
	}
	s := base64.StdEncoding.EncodeToString(buffer)
	log.Printf("RSA 加密后：%v", s)

	return s, nil
}

//RSA解密
func RsaDecrypt(value string, privateKey *rsa.PrivateKey) (string, error) {

	valueBytes, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	buffer, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, valueBytes)
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}
