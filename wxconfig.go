package wxpay

import (
	"errors"
	"net/url"

	"github.com/xwjdsh/httphelper"
)

type WxConfig struct {
	AppId     string
	AppKey    string
	MchId     string
	NotifyUrl string
	TradeType string

	//config check
	checked bool
}

func (this *WxConfig) NewWxPay() (*WxPay, error) {
	if this.AppId == "" || this.AppKey == "" || this.MchId == "" || this.NotifyUrl == "" || this.TradeType == "" {
		return nil, errors.New("any config can't equals empty string")
	}
	this.checked = true
	helper := &httpHelper.HttpHelper{CommonHeader: url.Values{}}
	return &WxPay{
		Config: this,
		http:   helper,
	}, nil
}
