res, err = http.Get("http://youtube.com") // função retorna 2 valores

/* Não vai ser executada: 
   Go Lang não permite declarar uma variável
   e não utiliza-la.
*/



// Tratar erro

if err != nil {
	//qualquer coisa ;)
}

/* Blank identifier: 
	(caso não queira tratar o erro)
	Faz com que o Go ignore o valor vazio
*/
res, _= http.Get("http://youtube.com")

