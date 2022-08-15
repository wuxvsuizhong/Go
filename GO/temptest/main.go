package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main1() {
	fmt.Println(time.Now().Unix())
	ftime := time.Now().Format("20060102150405.000")
	fmt.Println(ftime)
	var timeStr string = "20220508202700.536"
	t, _ := time.Parse("20060102150405.000", timeStr)
	fmt.Println(t.UnixMilli())
	fname := "logtest.log20220508205354.314"
	basefname := "logtest.log"
	fmt.Println(strings.Split(fname, "."))
	var arr []string
	arr = append(arr, "10086")
	arr = append(arr, "10010")
	arr = append(arr, "12345")
	arr = append(arr, "11111")

	// fmt.Println(arr[:2])
	fmt.Println(strings.Join(arr[:2], "-"))
	fmt.Println("abcd" == "abcd")
	fprefix := "logtest.log"
	re, _ := regexp.Compile(`^` + fprefix + `\d+`)
	fmt.Println(re.Match([]byte(fname)))
	fmt.Println(re.Match([]byte(basefname)))

	var flist []string
	flist = append(flist, "logtest.log.error20220508201637.231")
	flist = append(flist, "logtest.log.error20220508201622.309")
	flist = append(flist, "logtest.log.error20220508201906.002")
	flist = append(flist, "logtest.log.error20220508201954.352")
	flist = append(flist, "logtest.log.error20220508204834.152")
	flist = append(flist, "logtest.log.error20220508201939.223")
	fmt.Println(flist)

	sort.Slice(flist, func(i, j int) bool {
		return flist[i] < flist[j]
	})

	fmt.Println(flist)

	for _, val := range []string{"abc", "def"} {
		println(val)
	}
}

func main() {
	s := "1234567890"
	fmt.Println(s[3:5])
	s = fmt.Sprintf("%d", 'a')
	fmt.Println(s)
	ret, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	fmt.Println(ret)
	ret = ret - 96
	fmt.Println(ret)

	ecodeMap := map[int]string{}
	for i := 1; i <= 26; i++ {
		s = fmt.Sprintf("%c", i+96)
		ecodeMap[i] = s
	}
	fmt.Println(ecodeMap)

	ret, _ = strconv.Atoi("098")
	fmt.Printf("%d\n", ret)
}
