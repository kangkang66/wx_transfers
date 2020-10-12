package wx_transfers

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sort"
	"strings"
)

type XWXPay struct {
	httpTlsClient *http.Client
	apiKey		string
	baseRandStr string
}

//apiKey:API密钥
//certFile,keyFile:微信支付商户平台证书路径
func NewXWXPay(apiKey, certFile, keyFile string) (client *XWXPay, err error) {
	if apiKey == "" || certFile == "" || keyFile == "" {
		err = errors.New("wxpay param is empty")
		return
	}

	client = new(XWXPay)
	client.baseRandStr = "abcdefghijkmnpqrstuvwxyz23456789"
	client.apiKey = apiKey
	cert, err := ioutil.ReadFile(certFile)
	if err != nil {
		return
	}
	key, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return
	}
	tlsCert, err := tls.X509KeyPair(cert, key)
	if err != nil {
		return
	}
	client.httpTlsClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{tlsCert},
			},
		},
	}
	return
}

//企业付款
func (x *XWXPay) Transfers(ctx context.Context, param *ParamPayRequest) (payResponse *ParamPayResponse,err error) {
	urlAddr := "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers"
	response,err := x.post(ctx, urlAddr, param)
	if err != nil {
		return
	}
	payResponse = new(ParamPayResponse)
	err = xml.Unmarshal(response, payResponse)
	if err != nil {
		return
	}
	return
}

//查询企业付款
func (x *XWXPay) GetTransferInfo(ctx context.Context, param *ParamPayResultRequest) (payResultResponse *ParamPayResultResponse,err error) {
	urlAddr := "https://api.mch.weixin.qq.com/mmpaymkttransfers/gettransferinfo"
	response,err := x.post(ctx, urlAddr, param)
	if err != nil {
		return
	}
	payResultResponse = new(ParamPayResultResponse)
	err = xml.Unmarshal(response, payResultResponse)
	if err != nil {
		return
	}
	return
}

func (x *XWXPay) post(ctx context.Context, urlAddr string, param ParamRequester) (response []byte,err error) {
	//检查有没有设置随机值
	if param.GetNonceStr() == "" {
		param.SetNonceStr(x.createRandStr(32))
	}
	//检查有没有签名
	if param.GetSign() == "" {
		var sign string
		sign,err = x.createSign(ctx,param)
		if err != nil {
			return
		}
		param.SetSign(sign)
	}
	//转换成xml
	xmlByte,err := param.ToXml()
	if err != nil {
		return
	}
	//fmt.Println(string(xmlByte))

	//发送post
	resp,err := x.httpTlsClient.Post(urlAddr, "application/xml; charset=utf-8", bytes.NewReader(xmlByte))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	response,err = ioutil.ReadAll(resp.Body)
	return
}

func (x *XWXPay) createRandStr(length int) (string) {
	l := len(x.baseRandStr)
	randChar := make([]byte,0,length)
	for i:=0;i<length;i++ {
		randChar = append(randChar, x.baseRandStr[rand.Intn(l)])
	}
	return string(randChar)
}

func (x *XWXPay) createSign(ctx context.Context, param ParamRequester) (sign string, err error) {
	jbByte,err := json.Marshal(param)
	if err != nil {
		return
	}
	paramMap := map[string]string{}
	err = json.Unmarshal(jbByte, &paramMap)
	if err != nil {
		return
	}
	//fmt.Println(paramMap)

	var keys = make([]string, 0, len(paramMap))
	for k, v := range paramMap {
		if v != "" && k != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	var buf bytes.Buffer
	for _, k := range keys {
		buf.WriteString(k)
		buf.WriteString(`=`)
		buf.WriteString(paramMap[k])
		buf.WriteString(`&`)
	}
	buf.WriteString(`key=`)
	buf.WriteString(x.apiKey)

	//fmt.Println(buf.String())

	sum := md5.Sum(buf.Bytes())
	str := hex.EncodeToString(sum[:])
	sign = strings.ToUpper(str)
	return
}
