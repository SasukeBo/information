package utils

import (
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
