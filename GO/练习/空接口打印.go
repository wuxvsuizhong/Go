package main

func main() {
	var (
		a int         = 0
		b int64       = 0
		c interface{} = int(0)
		d interface{} = int64(0)
	)

	println(c == 0)
	println(c == a)
	println(c == b)
	println(d == b)
	println(d == 0)
	println(c == 0)

}
