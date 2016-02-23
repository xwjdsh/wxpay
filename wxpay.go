package wxpay

import (
	"fmt"
	"sort"
	"strings"

	"github.com/xwjdsh/httphelper"
)

type WxPay struct {
	Config *WxConfig
	Http   *httpHelper.HttpHelper
}

const (
	NEW_ORDER_URL = "https://api.mch.weixin.qq.com/pay/unifiedorder"
)

func (this *WxPay) NewOrder(param map[string]string) {
	this.formatParam(param)
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

func (this *WxPay) createOrder(params map[string]string) {
	xml := mapToXmlString(params)
	resp, err := this.Http.Send("POSTFORM", NEW_ORDER_URL, []byte(xml), nil)
	if err != nil {
		fmt.Println(resp)
	} else {
		fmt.Println("err:" + err.Error())
	}

}
