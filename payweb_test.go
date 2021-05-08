package sandsdk

import "testing"

func TestService_WebPay(t *testing.T) {
	s := NewService("", "", "")
	url, err := s.WebPay(&WebPayParam{
		Title: "测试充值",
		Info: map[string]string{
			"goodsDesc": "测试充值",
		},
		OrderId:   "11111111111111111118",
		Amount:    1,
		NotifyUrl: "",
		ResultUrl: "",
		Channel:   WebPayChannelH5,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(url)
}
