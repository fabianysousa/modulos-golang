func variadicFunc(x ...int) int {

}

variadicFunc(1,2,5,6,19,4,5)

func variadicFunc(x ...int) int {
	var res int
	for _, v := range x {
	res += v
	}
	return res
}
	