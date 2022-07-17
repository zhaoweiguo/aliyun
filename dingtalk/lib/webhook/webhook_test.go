package webhook

import (
	"testing"
)

func getRawMsg() []byte {
	return []byte(`{"text":{"content":"我就是我, @XXX 是不一样的烟火"},"msgtype":"text"}`)
}


func TestWebHook_SendRawMessage_WithoutSecret(t *testing.T) {
	access_token := "ding talk token"
	secret := ""
	debug := true
	webhooker := NewWebHook(access_token, secret, debug)
	webhooker.SendRawMessage(getRawMsg())
}

func TestWebHook_SendRawMessage_WithSecret(t *testing.T) {
	access_token := "ding talk token"
	secret := "ding talk secret"
	debug := true
	webhooker := NewWebHook(access_token, secret, debug)
	webhooker.SendRawMessage(getRawMsg())
}

