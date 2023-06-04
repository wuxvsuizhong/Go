package usestruct

/*
继承结构体的方法
*/
import "fmt"

type newperson struct {
	name string
	age  int
}

type newworker struct {
	addr string
	city string
	newperson
}

func (p newperson) getinfo() {
	fmt.Printf("%s的年级%d\n", p.name, p.age)
}

func TestInherit() {
	w1 := newworker{addr: "陕西", city: "西安", newperson: newperson{"张三", 30}}
	fmt.Println(w1)
	w1.getinfo()
	// newworker 继承了newperson 的方法getinfo
}
