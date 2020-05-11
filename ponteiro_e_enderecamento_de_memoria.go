// Acesso direto ao endereçaento de memória
x := 10
fmt.Println(&x) // 0xc82000a288


// Criar e apontar variável
y := &x
fmt.Println(y) // 0xc82000a288


// Exibir valor atribuido para o endereçamento de memoria
fmt.Println(*y) // 10

// Alterar valor atribuído na memória
*y = 20
fmt.Println(x) // 20

// Declarar variável definindo tipo
var z *int = &x
fmt.Println(z) // 0xc82000a288
fmt.Println(*z) // 20

// Ponteiro dentro de funções
func xpto(a *int) int {
	*a = 100
	return *a
}

b := 10
xpto(&b) // 100
fmt.Println(b) // 100

