package sandsdk

import (
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type WebPayChannel uint8

const (
	WebPayChannelH5       = WebPayChannel(iota) // 浏览器h5
	WebPayChannelWeChatMp                       // 微信公众号
)

type WebPayParam struct {
	Title     string            // 标题
	Info      map[string]string // 信息
	OrderId   string            // 订单号
	Amount    uint32            // 金额，单位：分
	NotifyUrl string            // 异步通知地址
	ResultUrl string            // 同步返回地址
	Channel   WebPayChannel     // 支付渠道
}

func (s *Service) WebPay(param *WebPayParam) (string, error) {
	if param == nil {
		return "", errors.New("param is nil")
	}

	if len(param.Info) == 0 {
		return "", errors.New("param info is nil")
	}

	info, err := jsoniter.Marshal(param.Info)
	if err != nil {
		return "", err
	}

	var (
		payModeList string
		userAgent   string
	)
	switch param.Channel {
	case WebPayChannelWeChatMp:
		payModeList = "[wxjsbridge]"
		userAgent = "MicroMessenger"
	default:
		payModeList = "[alipay]"
		userAgent = "Android iPhone"
	}

	body := map[string]map[string]string{
		"head": {
			"version":     "1.0",
			"method":      "sandpay.trade.orderCreate",
			"productId":   "00002000",
			"accessType":  "1",
			"mid":         s.MerId,
			"channelType": "07",
			"reqTime":     time.Now().Format("20060102150405"),
		},
		"body": {
			"orderCode":   param.OrderId,
			"totalAmount": fmt.Sprintf("%012d", param.Amount),
			"subject":     param.Title,
			"body":        string(info),
			"payModeList": payModeList,
			"notifyUrl":   param.NotifyUrl,
			"frontUrl":    param.ResultUrl,
		},
	}

	data, err := jsoniter.Marshal(body)
	if err != nil {
		return "", err
	}

	sign, err := s.Sign(data)
	if err != nil {
		return "", err
	}

	val := make(url.Values)
	val.Set("charset", "utf-8")
	val.Set("signType", "01")
	val.Set("data", string(data))
	val.Set("sign", sign)
	req, err := http.NewRequest(http.MethodPost, "https://cashier.sandpay.com.cn/gw/web/order/create", strings.NewReader(val.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", userAgent)

	c := &http.Client{
		Timeout: 3 * time.Second,
	}

	res, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	return res.Request.Header.Get("Referer"), nil
}
