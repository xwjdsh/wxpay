package wxpay

import (
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/xwjdsh/httphelper"
)

type WxPay struct {
	Config *WxConfig
	http   *httpHelper.HttpHelper
}

const (
	NEW_ORDER_URL = "https://api.mch.weixin.qq.com/pay/unifiedorder"
)

func (this *WxPay) check() bool {
	return this.Config != nil && this.http != nil && this.Config.checked
}

func (this *WxPay) NewOrder(nonce_str, body, out_trade_no, total_fee, spbill_create_ip, trade_type string, param map[string]string) (*NewOrderResp, error) {
	if !this.check() {
		return nil, errors.New("wrong object!")
	}
	if err := checkNotEmpty(nonce_str, body, out_trade_no, total_fee, spbill_create_ip, trade_type); err != nil {
		return nil, err
	}
	if param == nil {
		param = make(map[string]string, 0)
	}
	//prepare params
	this.formatParam(nonce_str, body, out_trade_no, total_fee, spbill_create_ip, trade_type, param)
	//send new order request
	return this.createOrder(param)
}

func (this *WxPay) formatParam(nonce_str, body, out_trade_no, total_fee, spbill_create_ip, trade_type string, params map[string]string) {
	params["nonce_str"] = nonce_str
	params["body"] = body
	params["out_trade_no"] = out_trade_no
	params["total_fee"] = total_fee
	params["spbill_create_ip"] = spbill_create_ip
	params["trade_type"] = trade_type

	params["appid"] = this.Config.AppId
	params["mch_id"] = this.Config.MchId
	params["notify_url"] = this.Config.NotifyUrl
	params["sign"] = GenerateSign(params, this.Config.AppKey)
	fmt.Println("sign", params["sign"])
}

func checkNotEmpty(checkParams ...string) error {
	for _, param := range checkParams {
		if param == "" {
			return errors.New("必填参数存在空项!")
		}
	}
	return nil
}

func (this *WxPay) createOrder(params map[string]string) (*NewOrderResp, error) {
	fmt.Println(params)
	xmlString := mapToXmlString(params)
	fmt.Println(xmlString)
	resp, err := this.http.Send("POSTFORM", NEW_ORDER_URL, []byte(xmlString), nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(resp[:]))
	newOrderResp := NewOrderResp{}
	err = xml.Unmarshal(resp, &newOrderResp)
	if err != nil {
		return nil, err
	}
	return &newOrderResp, nil
}
