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
}

func (this *WxConfig) NewWxPay() (*WxPay, error) {
	if this.AppId == "" || this.AppKey == "" || this.MchId == "" || this.NotifyUrl == "" || this.TradeType == "" {
		return nil, errors.New("any config can't equals empty string")
	}
	helper := &httpHelper.HttpHelper{CommonHeader: url.Values{}}
	return &WxPay{
		Config: this,
		Http:   helper,
	}, nil
}
