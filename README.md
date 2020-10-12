# golang微信企业付款sdk
用于企业向微信用户个人付款，目前支持向指定微信用户的openid付款。  
微信官方文档：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_1

# 使用
## 企业付款到零钱
```
func TestXWXPay_Transfers(t *testing.T) {
	ctx  := context.Background()
	certPemFile := "apiclient_cert.pem"
	keyPemFile := "apiclient_key.pem"

	wxPayClient,err := NewXWXPay("apikey",certPemFile,keyPemFile)
	assert.NoError(t,err)
	param := &ParamPayRequest{
		MchAppid:       "MchAppid",
		Mchid:          "Mchid",
		PartnerTradeNo: "PartnerTradeNo",
		Openid:         "Openid",
		CheckName:      "NO_CHECK",
		Amount:         "100",
		Desc:           "test",
	}
	resp,err := wxPayClient.Transfers(ctx,param)
	assert.NoError(t,err)
	fmt.Printf("%+v",resp)
}
```

## 查询企业付款结果
```
func TestXWXPay_GetTransferInfo(t *testing.T) {
	ctx  := context.Background()
	certPemFile := "apiclient_cert.pem"
	keyPemFile := "apiclient_key.pem"

	wxPayClient,err := NewXWXPay("apikey",certPemFile,keyPemFile)
	assert.NoError(t,err)
	param := &ParamPayResultRequest{
		PartnerTradeNo: "PartnerTradeNo",
		Appid:          "Appid",
		MchId:          "MchId",
	}
	resp,err := wxPayClient.GetTransferInfo(ctx,param)
	assert.NoError(t,err)
	fmt.Printf("%+v",resp)
}
```

