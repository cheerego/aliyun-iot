package iot

import (
	"testing"
)

func TestRandStr(t *testing.T) {
	if len(GetRandomString(12)) == 12{
		t.Log("测试通过")
	}else {
		t.Error("测试不通过")
	}
}
