package comm

var num int

var NumStruct struct {
	n int
	m int
}

func GetNum() *int {
	return &num
}

var InnerNum *NumStruct = &NumStruct{}

func GetInneNum() *NumStruct {
	return InnerNum
}
