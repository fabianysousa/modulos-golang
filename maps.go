m := make(map[string]int)
m["a"] = 19
m["b"] = 20

// Deletar valor map
delete(m,"b")


// Verificar existencia de valor no mapa
_, exists := m["b"]
fmt.Println(exists) // false


// Outra forma de verificar valor
var x = map[string]int{}

// ou 3 

x := map[string]int{"a":5,"b":10}

