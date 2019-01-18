package iot

import (
	"testing"
)

func TestRandStringBytesMaskImprSrc(t *testing.T) {
	m := make(map[string]string)
	ch := make(chan string, 100)
	for i := 0; i < 100; i++ {
		go genRand(ch)
	}
	for i := 0; i < 100; i++ {
		s := <-ch
		m[s] = s
	}
	if len(m) == 100 {
		t.Log("同一时刻生成随机字符串重复测试成功")
	} else {
		t.Error("同一时刻生成随机字符串重复测试失败")
	}
}
func genRand(ch chan string) {
	ch <- RandStringBytesMaskImprSrc(14)
}

func BenchmarkRandStringBytesMaskImprSrc(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
