package wxpay

import (
	"encoding/xml"
	"fmt"
	"sort"
	"strings"

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

func (this *WxPay) NewOrder(param map[string]string) {
	if !this.check() {
		return
	}
	//prepare params
	this.formatParam(param)
	//send new order request
	this.createOrder(param)
}

func (this *WxPay) formatParam(params map[string]string) {
	if params == nil {
		params = make(map[string]string, 0)
	}
	params["appid"] = this.Config.AppId
	params["mch_id"] = this.Config.MchId
	params["notify_url"] = this.Config.NotifyUrl
	params["trade_type"] = this.Config.TradeType

	keys := []string{}
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var result string
	for i, k := range keys {
		if i == 1 {
			result += fmt.Sprintf("%s=%s", k, params[k])
		}
		result += fmt.Sprintf("&%s=%s", k, params[k])
	}
	result += fmt.Sprintf("&%s=%s", "key", this.Config.AppKey)
	params["sign"] = strings.ToUpper(getMD5Hash(result))
}

func (this *WxPay) createOrder(params map[string]string) (*NewOrderResp, error) {
	xmlString := mapToXmlString(params)
	resp, err := this.http.Send("POSTFORM", NEW_ORDER_URL, []byte(xmlString), nil)
	if err != nil {
		return nil, err
	}
	newOrderResp := NewOrderResp{}
	err = xml.Unmarshal(resp, &newOrderResp)
	if err != nil {
		return nil, err
	}
	return &newOrderResp, nil
}
