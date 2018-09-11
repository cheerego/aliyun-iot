package iot

import (
	"testing"
)

func TestRandStr(t *testing.T) {
	if len(RandStr(2)) == 2{
		t.Log("测试通过")
	}else {
		t.Error("测试不通过")
	}
}
