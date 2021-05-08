package sandsdk

import (
	jsoniter "github.com/json-iterator/go"
	"net/url"
)

type PayNotify struct {
	Charset  string
	SignType string
	Sign     string
	dataStr  string
	Data     *PayNotifyData
}

type PayNotifyData struct {
	Head PayNotifyDataHead
	Body PayNotifyDataBody
}

type PayNotifyDataHead struct {
	Version  string `json:"version"`
	RespTime string `json:"respTime"`
	RespCode string `json:"respCode"`
	RespMsg  string `json:"respMsg"`
}

type PayNotifyDataBody struct {
	TotalAmount    string `json:"totalAmount"`
	ClearDate      string `json:"clearDate"`
	Credential     string `json:"credential"`
	TradeNo        string `json:"tradeNo"`
	PayTime        string `json:"payTime"`
	BuyerPayAmount string `json:"buyerPayAmount"`
	OrderCode      string `json:"orderCode"`
	DiscAmount     string `json:"discAmount"`
}

func (s *Service) PayParseNotify(data string) (*PayNotify, error) {
	vals, err := url.ParseQuery(data)
	if err != nil {
		return nil, err
	}

	var ret = new(PayNotify)
	ret.Charset = vals.Get("charset")
	ret.SignType = vals.Get("signType")
	ret.Sign = vals.Get("sign")
	ret.dataStr = vals.Get("data")
	ret.Data = new(PayNotifyData)
	if err = jsoniter.Unmarshal([]byte(ret.dataStr), ret.Data); err != nil {
		return nil, err
	}

	return ret, nil
}
