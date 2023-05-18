# sandpay
>  杉德支付 
  网银支付
  代下单服务!

## 提供的pfx证书在golang中解析会失败 需要转为pem证书解析pfx转pem方法,需要安装openssl
  
## 解析pem证书的公钥和私钥
```shell script
  openssl pkcs12 -in xxxx.pfx -nodes -out server.pem  #生成为原生格式pem 私钥
  openssl rsa -in server.pem -out server.key          #生成为rsa格式私钥文件
  openssl x509 -in server.pem  -out server.crt
  openssl pkcs12 -in xxxx.pfx -clcerts -nokeys -out key.cert
```

```go 
        var config SandPayConfig
	config.MerId = "xxxxxxxxxxxxx"      //商户号
	config.PrivatePath = "server.key"      //私钥
	config.CertPath = "key.cert"        //公钥
	config.EncryptCertPath = "sand.cer" //杉德公钥
```
###  在目录下生成server.key即为私钥文件.
###  在目录下生成server.crt即为公钥文件.
  
