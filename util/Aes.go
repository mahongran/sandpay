package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

type SandAES struct {
	Key []byte
}

func (s *SandAES) Pkcs7Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

//pkcs7UnPadding 填充的反向操作
func (s *SandAES) Pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误")
	}
	// 获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

func (s *SandAES) AESEncrypt(data []byte) ([]byte, error) {
	// 创建加密实例
	block, err := aes.NewCipher(s.Key)
	if err != nil {
		return nil, err
	}
	//判断加密块的大小
	blockSize := block.BlockSize()
	//填充
	encryptBytes := s.Pkcs7Padding(data, blockSize)
	// 初始化加密数据接受切片
	crypted := make([]byte, len(encryptBytes))
	// 使用cbc加密
	blockMod := cipher.NewCBCEncrypter(block, s.Key[:blockSize])

	// 执行加密
	blockMod.CryptBlocks(crypted, encryptBytes)
	return crypted, nil
}

//解密
func (s *SandAES) AESDecrypt(data []byte) ([]byte, error) {
	//创建实例
	block, err := aes.NewCipher(s.Key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, s.Key[:blockSize])

	//初始化数据
	crypted := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(crypted, data)
	//去处填充
	crypted, err = s.Pkcs7UnPadding(crypted)
	if err != nil {
		return nil, err
	}
	return crypted, nil
}

func (s *SandAES) EncryptByAES(data []byte) (string, error) {
	res, err := s.AESEncrypt(data)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(res), nil
}

//解密

func (s *SandAES) DecryptByAES(data string) ([]byte, error) {
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return s.AESDecrypt(dataByte)
}

func (s *SandAES) RandStr(n int) string {
	str := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	strData := strings.Builder{}
	for i := 0; i < n; i++ {
		index := rand.Intn(len(str))
		strData.WriteString(str[index])
	}
	res := strData.String()
	//fmt.Println("AES随机数 16,24,32 位，否则报错:", res)
	return res
}

func (s *SandAES) Encypt5(data []byte) string {
	block, _ := aes.NewCipher(s.Key)
	data = s.Pkcs7Padding(data, block.BlockSize())
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Encrypt(decrypted[bs:be], data[bs:be])
	}

	return base64.StdEncoding.EncodeToString(decrypted)
}

func (s *SandAES) EcbDecrypt(data, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}

	b, _ := s.Pkcs7UnPadding(decrypted)
	return b
}

// AesEcbPkcs5Padding AES ECB  Pkcs5Padding 加密
func (s *SandAES) AesEcbPkcs5Padding(key, plaintext string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	// 使用PKCS#5填充明文
	paddedPlaintext := pkcs5Padding([]byte(plaintext), block.BlockSize())

	// 加密
	ciphertext := make([]byte, len(paddedPlaintext))
	encrypter := NewECBEncrypter(block)
	encrypter.CryptBlocks(ciphertext, paddedPlaintext)

	// 输出加密结果和密钥
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// PKCS#5填充
func pkcs5Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// 自定义ECB模式加密器
type ecbEncrypter struct {
	b         cipher.Block
	blockSize int
}

func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return &ecbEncrypter{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

func (x *ecbEncrypter) BlockSize() int {
	return x.blockSize
}

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		fmt.Println("ciphertext length must be multiple of block size")
		return
	}
	if len(dst) < len(src) {
		fmt.Println("output smaller than input")
		return
	}

	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

// AesEcbPkcs5PaddingDecrypt AES 解密
func (s *SandAES) AesEcbPkcs5PaddingDecrypt(key, ciphertext string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	// 解码密文
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	// 解密
	plaintext := make([]byte, len(decodedCiphertext))
	decrypter := NewECBDecrypter(block)
	decrypter.CryptBlocks(plaintext, decodedCiphertext)

	// 去除填充
	unpaddedPlaintext := pkcs5Unpadding(plaintext)

	// 输出解密结果
	return string(unpaddedPlaintext), nil
}

// PKCS#5去除填充
func pkcs5Unpadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:length-unpadding]
}

// 自定义ECB模式解密器
type ecbDecrypter struct {
	b         cipher.Block
	blockSize int
}

func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return &ecbDecrypter{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

func (x *ecbDecrypter) BlockSize() int {
	return x.blockSize
}

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		fmt.Println("ciphertext length must be multiple of block size")
		return
	}
	if len(dst) < len(src) {
		fmt.Println("output smaller than input")
		return
	}

	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
