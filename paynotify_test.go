package sandsdk

import "testing"

func TestService_PayNotifyVerifySign(t *testing.T) {
	s := NewService("", "", "")
	payNotify, err := s.PayParseNotify("")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(payNotify)

	if err = s.PayNotifyVerifySign(payNotify); err != nil {
		t.Fatal(err)
	}
	t.Log("verify success")
}
