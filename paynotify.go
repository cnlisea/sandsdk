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
	Mid                 string `json:"mid"`
	OrderCode           string `json:"orderCode"`
	TotalAmount         string `json:"totalAmount"`
	OrderStatus         string `json:"orderStatus"`
	TradeNo             string `json:"tradeNo"`
	SettleAmount        string `json:"settleAmount"`
	BuyerPayAmount      string `json:"buyerPayAmount"`
	DiscAmount          string `json:"discAmount"`
	PayTime             string `json:"payTime"`
	ClearDate           string `json:"clearDate"`
	AccNo               string `json:"accNo"`
	MidFee              string `json:"midFee"`
	ExtraFee            string `json:"extraFee"`
	SpecialFee          string `json:"specialFee"`
	PlMidFee            string `json:"plMidFee"`
	BankSerial          string `json:"bankserial"`
	TxnCompleteTime     string `json:"txnCompleteTime"`
	PayOrderCode        string `json:"payordercode"`
	ExternalProductCode string `json:"externalProductCode"`
	CardNo              string `json:"cardNo"`
	CreditFlag          string `json:"creditFlag"`
	Bid                 string `json:"bid"`
	Extend              string `json:"extend"`
	FundBillList        string `json:"fundBillList"`
	PayDetail           string `json:"payDetail"`
	BenefitAmount       string `json:"benefitAmount"`
	RemittanceCode      string `json:"remittanceCode"`
	AccLogonNo          string `json:"accLogonNo"`
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
