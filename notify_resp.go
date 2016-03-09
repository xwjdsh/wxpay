package wxpay

import "encoding/xml"

type NotifyResp struct {
	XMLName     xml.Name `xml:"xml"`
	ReturnCode  string   `xml:"return_code"`
	ReturnMsg   string   `xml:"return_msg"`
	AppId       string   `xml:"appid"`
	MchId       string   `xml:"mch_id"`
	DeviceInfo  string   `xml:"device_info"`
	NonceStr    string   `xml:"nonce_str"`
	Sign        string   `xml:"sign"`
	ResultCode  string   `xml:"result_code"`
	ErrCode     string   `xml:"err_code"`
	ErrCodeDesc string   `xml:"err_code_des"`
	OpenId      string   `xml:"openid"`
	IsSubscribe string   `xml:"is_subscribe"`
	TradeType   string   `xml:"trade_type"`
	BankType    string   `xml:"bank_type"`
	TotalFee    string   `xml:"total_fee"`
	FeeType     string   `xml:"fee_type"`
	CashFee     string   `xml:"cash_fee"`
	CashFeeType string   `xml:"cash_fee_type"`
	CouponFee   string   `xml:"coupon_fee"`
	CouponCount string   `xml:"coupon_count"`

	TransactionId string `xml:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	Attach        string `xml:"attach"`
	TimeEnd       string `xml:"time_end"`
}
