// Slices = Arrays com esteróides

//Não possuem tamanho fixo e é dinâmico

slice := make([]int)


// Slice com quantidade definida de repetições
		
slice := make([]int, 5) // 5 posições padrão.
slice[0] = 10
fmt.Println(slice[0]) // 10
slice = append(slice, 1,2,3,47,5)


// Outra forma de declarar slices
sliceString := []string {
	"Hello",
	"World"
}

fmt.Println(slideString[0]) // "Hello"


//Recursos:
	// Navegação pelo ídices

sliceString := []string {
	"Hello",
	"World",
	"Much",
	"Better"
}

fmt.Println(sliceString[:2]) // "Hello World"
fmt.Println(sliceString[1:2]) // "World" - A partir do índice um até a segunda posição.
fmt.Println(sliceString[2:4]) // "Much Better" - A partir do índice 2 até a quarta posição.
fmt.Println(sliceString[2:]) // "Much Better" - A partir do índice 2 até ao final.
