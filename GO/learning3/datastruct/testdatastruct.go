package datastruct

import "fmt"

func TestDataStruct(){
	var nums [5]int
	nums[0] = 23
	nums[1] = 11
	nums[2] = 30
	nums[3] = 49
	nums[4] = 3

	var sum int
	for i:=0;i<len(nums);i++ {
		sum += nums[i]
	}
	avg := sum/len(nums)
	fmt.Printf("avg:%v\n",avg)
}

func TestDatastru2(){
	var scores [5]int
	for i:=0;i<len(scores);i++ {
		fmt.Printf("输入第%v个学生成绩:",i)
		fmt.Scanln(&scores[i])
	}

	for i := 0;i < len(scores);i++ {
		fmt.Printf("第%v个学生成绩为:%v\n",i,scores[i])
	}

	fmt.Println("----------------------------")
	for key,val := range scores {
		fmt.Printf("第%d个学生成绩为:%d\n",key,val)
	}

}
