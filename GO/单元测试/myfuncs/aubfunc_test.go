package myfuncs

import "testing"

func TestSub(t *testing.T) {
  //  调用
  ret := Sub(10, 5)
  if ret != 5 {
    t.Fatalf("测试未通过Sub(10,5)的期望结果:%v,实际是：%v\n", 5, ret)
  }

  t.Logf("Sub(10,5)的测试结果是:%v\n", ret)
}
