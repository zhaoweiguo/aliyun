package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// 自定义机器人安全设置: https://open.dingtalk.com/document/robots/customize-robot-security-settings

func getSign(timestamp int64, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	str := fmt.Sprintf("%d\n%s", timestamp, secret)
	h.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}



