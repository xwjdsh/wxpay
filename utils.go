package wxpay

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
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

func RandStr(strlen int) string {
	rand.Seed(time.Now().Unix())
	data := make([]byte, strlen)
	var num int
	for i := 0; i < strlen; i++ {
		num = rand.Intn(57) + 65
		for {
			if num > 90 && num < 97 {
				num = rand.Intn(57) + 65
			} else {
				break
			}
		}
		data[i] = byte(num)
	}
	return string(data)
}

func GenerateSign(params map[string]string, key string) string {
	keys := []string{}
	for k, _ := range params {
		if k != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	var result string
	for i, k := range keys {
		if i == 0 {
			result += fmt.Sprintf("%s=%s", k, params[k])
		} else {
			result += fmt.Sprintf("&%s=%s", k, params[k])
		}
	}
	result += fmt.Sprintf("&%s=%s", "key", key)
	return strings.ToUpper(getMD5Hash(result))
}
