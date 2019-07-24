package utils

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

// GenSmsCode 生成
func GenSmsCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	smsCodeLen, _ := strconv.ParseInt(beego.AppConfig.String("smscodelen"), 10, 0)
	intCode := int(math.Floor(r.Float64() * math.Pow10(int(smsCodeLen))))
	strCode := fmt.Sprintf("%d", intCode)

	if codeLen := len(strCode); codeLen < int(smsCodeLen) {
		prefixZero := strings.Repeat("0", int(smsCodeLen)-codeLen)
		return strings.Join([]string{prefixZero, strCode}, "")
	}
	return strCode
}

// Encrypt md5 + base64
func Encrypt(data string) string {
	r := md5.Sum([]byte(data))
	rs := fmt.Sprintf("%x", r)
	str := base64.StdEncoding.EncodeToString([]byte(rs))

	return str
}
