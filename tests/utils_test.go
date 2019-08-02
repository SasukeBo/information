package utils

import (
	"testing"

	"github.com/SasukeBo/information/utils"
)

func Test_GenSmsCode(t *testing.T) {
	code := utils.GenSmsCode()
	if code != "117650" {
		t.Error("生成验证码: ", code)
	}
}
