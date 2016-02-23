package wxpay

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func mapToXmlString(params map[string]string) string {
	xml := "<xml>"
	for k, v := range params {
		xml += fmt.Sprintf("<%s>%s</%s>", k, v, k)
	}
	xml += "</xml>"

	return xml
}
