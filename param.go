package wx_transfers

import "encoding/xml"

type ParamRequester interface {
	//转换为xml
	ToXml() ([]byte,error)
	//随机值
	GetNonceStr() string
	SetNonceStr(string)
	//签名
	GetSign() string
	SetSign(string)
}

//参数解释:https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2
//请求支付，给用户打款
type ParamPayRequest struct {
	XMLName        xml.Name `xml:"xml" json:"-"`
	//商户账号appid,必填
	MchAppid       string   `xml:"mch_appid" json:"mch_appid"`
	//商户号,必填
	Mchid          string   `xml:"mchid" json:"mchid"`
	//设备号
	//DeviceInfo	   string 	`xml:"device_info" json:"device_info"`
	//随机字符串,必填
	NonceStr       string   `xml:"nonce_str" json:"nonce_str"`
	//商户订单号,必填
	PartnerTradeNo string   `xml:"partner_trade_no" json:"partner_trade_no"`
	//用户openid,必填
	Openid         string   `xml:"openid" json:"openid"`
	//校验用户姓名选项,必填 (NO_CHECK：不校验真实姓名,FORCE_CHECK：强校验真实姓名)
	CheckName      string   `xml:"check_name" json:"check_name"`
	//收款用户姓名
	//ReUserName     string   `xml:"re_user_name" json:"re_user_name"`
	//金额(分),必填
	Amount         string   `xml:"amount" json:"amount"`
	//企业付款备注,必填
	Desc           string   `xml:"desc" json:"desc"`
	//Ip地址
	//SpbillCreateIp string   `xml:"spbill_create_ip" json:"spbill_create_ip"`
	//签名,必填
	Sign           string   `xml:"sign" json:"-"`
}

//转换为xml
func (p *ParamPayRequest) ToXml() (xmlData []byte,err error) {
	xmlData,err = xml.Marshal(p)
	return
}
//获取随机值
func (p *ParamPayRequest) GetNonceStr() string {
	return p.NonceStr
}
//设置随机值
func (p *ParamPayRequest) SetNonceStr(nonceStr string) {
	p.NonceStr = nonceStr
}
//获取签名
func (p *ParamPayRequest) GetSign() string {
	return p.Sign
}
//设置签名
func (p *ParamPayRequest) SetSign(sign string) {
	p.Sign = sign
}

//请求支付响应结果
type ParamPayResponse struct {
	XMLName        xml.Name `xml:"xml"`
	//返回状态码,SUCCESS/FAIL,此字段是通信标识，非付款标识，付款是否成功需要查看result_code来判断
	ReturnCode     string   `xml:"return_code" json:"return_code"`
	//返回信息,返回信息，如非空，为错误原因
	ReturnMsg      string   `xml:"return_msg" json:"return_msg"`
	/***********以下字段在return_code为SUCCESS的时候有返回*********/
	//商户appid
	MchAppid       string   `xml:"mch_appid" json:"mch_appid"`
	//商户号
	Mchid          string   `xml:"mchid" json:"mchid"`
	//设备号
	DeviceInfo     string   `xml:"device_info" json:"device_info"`
	//随机字符串
	NonceStr       string   `xml:"nonce_str" json:"nonce_str"`
	//业务结果,SUCCESS/FAIL，注意：当状态为FAIL时，存在业务结果未明确的情况。如果状态为FAIL，请务必关注错误代码（err_code字段），通过查询查询接口确认此次付款的结果。
	ResultCode     string   `xml:"result_code" json:"result_code"`
	//错误码信息，注意：出现未明确的错误码时（SYSTEMERROR等），请务必用原商户订单号重试，或通过查询接口确认此次付款的结果。
	ErrCode		   string 	`xml:"err_code" json:"err_code"`
	//错误代码描述,结果信息描述
	ErrCodeDes	   string   `xml:"err_code_des" json:"err_code_des"`
	/*********以下字段在return_code 和result_code都为SUCCESS的时候有返回*********/
	//商户订单号
	PartnerTradeNo string   `xml:"partner_trade_no" json:"partner_trade_no"`
	//微信付款单号
	PaymentNo      string   `xml:"payment_no" json:"payment_no"`
	//付款成功时间
	PaymentTime    string   `xml:"payment_time" json:"payment_time"`
}
//转换为xml
func (p *ParamPayResponse) ToXml() (xmlData []byte,err error) {
	xmlData,err = xml.Marshal(p)
	return
}

//参数解释:https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_3
//查询支付状态请求
type ParamPayResultRequest struct {
	XMLName        xml.Name `xml:"xml" json:"-"`
	NonceStr       string   `xml:"nonce_str" json:"nonce_str"`
	Sign           string   `xml:"sign" json:"-"`
	//商户订单号,必填
	PartnerTradeNo string   `xml:"partner_trade_no" json:"partner_trade_no"`
	//商户账号appid,必填
	Appid          string 	`xml:"appid" json:"appid"`
	//商户号,必填
	MchId          string   `xml:"mch_id" json:"mch_id"`
}

//转换为xml
func (p *ParamPayResultRequest) ToXml() (xmlData []byte,err error) {
	xmlData,err = xml.Marshal(p)
	return
}
//获取随机值
func (p *ParamPayResultRequest) GetNonceStr() string {
	return p.NonceStr
}
//设置随机值
func (p *ParamPayResultRequest) SetNonceStr(nonceStr string) {
	p.NonceStr = nonceStr
}
//获取签名
func (p *ParamPayResultRequest) GetSign() string {
	return p.Sign
}
//设置签名
func (p *ParamPayResultRequest) SetSign(sign string) {
	p.Sign = sign
}


//查询支付状态响应
type ParamPayResultResponse struct {
	XMLName        xml.Name `xml:"xml"`
	//返回状态码,SUCCESS/FAIL,此字段是通信标识，非付款标识，付款是否成功需要查看result_code来判断
	ReturnCode     string   `xml:"return_code" json:"return_code"`
	//返回信息,如非空，为错误原因
	ReturnMsg      string   `xml:"return_msg" json:"return_msg"`

	/**********以下字段在return_code为SUCCESS的时候有返回*******/
	//业务结果, SUCCESS/FAIL ，非付款标识，付款是否成功需要查看status字段来判断
	ResultCode     string   `xml:"result_code" json:"result_code"`
	//错误代码,错误码信息
	ErrCode		   string 	`xml:"err_code" json:"err_code"`
	//错误代码描述,结果信息描述
	ErrCodeDes	   string   `xml:"err_code_des" json:"err_code_des"`

	/**********以下字段在return_code 和result_code都为SUCCESS的时候有返回*******/
	//商户单号
	PartnerTradeNo string   `xml:"partner_trade_no" json:"partner_trade_no"`
	//商户号
	MchID          string   `xml:"mch_id" json:"mch_id"`
	//商户号的appid
	Appid          string   `xml:"appid" json:"appid"`
	//付款单号,调用企业付款API时，微信系统内部产生的单号
	DetailID       string   `xml:"detail_id" json:"detail_id"`
	//转账状态,SUCCESS:转账成功,FAILED:转账失败,PROCESSING:处理中
	Status         string   `xml:"status" json:"status"`
	//失败原因,如果失败则有失败原因
	Reason         string   `xml:"reason" json:"reason"`
	//付款金额
	PaymentAmount  string   `xml:"payment_amount" json:"payment_amount"`
	//收款用户openid
	Openid         string   `xml:"openid" json:"openid"`
	//转账时间,发起转账的时间
	TransferTime   string   `xml:"transfer_time" json:"transfer_time"`
	//付款成功时间,企业付款成功时间
	PaymentTime   string   `xml:"payment_time" json:"payment_time"`
	//收款用户姓名
	TransferName   string   `xml:"transfer_name" json:"transfer_name"`
	//企业付款备注
	Desc           string   `xml:"desc" json:"desc"`
}
//转换为xml
func (p *ParamPayResultResponse) ToXml() (xmlData []byte,err error) {
	xmlData,err = xml.Marshal(p)
	return
}
