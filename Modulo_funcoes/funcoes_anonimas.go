// Função anônima
a:= func() int {
	z +=2
}

// Função dentro de outra função
func funcInsideFunc() func() int {
	x := 10
	return func() int {
		return x * x
	}
}

	

	