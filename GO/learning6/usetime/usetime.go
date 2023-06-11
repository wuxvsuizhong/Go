package usetime

import (
	"fmt"
	"mypro/getfuncinfo"
	"time"
)

func TestTime() {
	getfuncinfo.PrintFuncName()
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	//时间戳
	fmt.Println(now.Unix())
	//事件戳(纳秒)
	fmt.Println("----------------纳秒时间戳----------------------")
	fmt.Println(now.UnixNano())
	//根据时间戳得到相应的时间
	//形参：秒，纳秒
	ret := time.Unix(1651671768, 0)
	//fmt.Println(ret.Year())
	fmt.Printf("%v年%v月%v号\n", ret.Year(), ret.Month(), ret.Day())

	//解析时间格式字符串，得到相应时间的时间戳
	tstr := "2022-05-04"
	timeobj, err := time.Parse("2006-01-02", tstr)
	if err != nil {
		fmt.Println("解析事件错误：", err)
		return
	}
	fmt.Printf("%s对应的事件戳：%d\n", tstr, timeobj.Unix())
	//fmt.Println(timeobj.Unix())

	//sleep
	//go 中sleep需要指明时间单位
	n := 5
	time.Sleep(5 * time.Second)
	time.Sleep(time.Duration(n) * time.Second)

}

func TestTimeSub() {
	//时间相减
	getfuncinfo.PrintFuncName()
	now := time.Now()
	fmt.Println(now)
	//按照字符串获取时间对应的时间戳(把传入的事件字符串视为是UTC对应时间,没有当地时区)
	timestr := "2022-05-04 20:00:00"
	timeobj, err := time.Parse("2006-01-02 15:04:05", timestr)
	fmt.Println(timestr, "对应的时间戳为：", timeobj)
	//按照字符串加载时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("解析location出错")
		return
	}
	//按照时区解析时间
	timeobj2, err := time.ParseInLocation("2006-01-02 15:04:05", timestr, location)
	if err != nil {
		fmt.Println("解析时区时间戳出错!")
		return
	}
	fmt.Println("按照本地时区解析,", timestr, "对应的时间戳为:", timeobj2)
	//时间对象相减
	timeduration := timeobj2.Sub(now)
	fmt.Println(timeduration)
}
