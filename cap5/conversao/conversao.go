package main

import "fmt"

type ListaCompras []string

func imprimirSlice(slice []string) {
	fmt.Println("Slice:", slice)
}

func imprimirLista(lista ListaCompras) {
	fmt.Println("Lista de compras:", lista)
}

func main() {
	lista := ListaCompras{"Alface", "Atum", "Azeite"}
	slice := []string{"Alface", "Atum", "Azeite"}

	imprimirSlice([]string(lista))
	imprimirLista(ListaCompras(slice))
}
