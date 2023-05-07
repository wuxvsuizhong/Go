package main

import (
  "fmt"
  "time"
)

func main() {
  now := time.Now() //获取当前时间
  fmt.Printf("now = %v,type : %T", now, now)

  //获取年月日，时分秒
  fmt.Println("年：", now.Year())
  fmt.Println("月：", now.Month())
  fmt.Println("月(数字)", int(now.Month()))
  fmt.Println("日：", now.Day())
  fmt.Println("时：", now.Hour())
  fmt.Println("分：", now.Minute())
  fmt.Println("秒：", now.Second())

  //时间格式化
  fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(),
    now.Hour(), now.Minute(), now.Second())

  //格式化时间字符串保存到变量中
  dataStr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(),
    now.Hour(), now.Minute(), now.Second())
  fmt.Printf("dataStr:%v\n", dataStr)

  //格式化日期时间的第二种方式 ,这种方式的时间格式是一组固定的数字
  fmt.Printf(now.Format("2006-01-02 15:04:05"))
  fmt.Println()
  fmt.Printf(now.Format("2006-01-02"))
  fmt.Println()
  fmt.Printf(now.Format("15:04:05"))
  fmt.Println()
  fmt.Printf(now.Format("2006"))
  fmt.Println("年")
  fmt.Printf(now.Format("01"))
  fmt.Println("月")
  fmt.Printf(now.Format("02"))
  fmt.Println("日")

}
