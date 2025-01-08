package main

import "fmt"

type ListaGenerica []interface{}

func (lista *ListaGenerica) RemoverIndice(indice int) interface{} {
	l := *lista
	removido := l[indice]
	*lista = append(l[0:indice], l[indice+1:]...)
	return removido
}

func (lista *ListaGenerica) RemoverInicio() interface{} {
	return lista.RemoverIndice(0)
}

func (lista *ListaGenerica) RemoverFim() interface{} {
	return lista.RemoverIndice(len(*lista) - 1)
}

func main() {
	lista := ListaGenerica{1, "Café", 42, true, 23, "Bola", 3.14, false}

	fmt.Println("Lista original:", lista)
	fmt.Println("Removendo do início:", lista.RemoverInicio(), lista)
	fmt.Println("Removendo do fim:", lista.RemoverFim(), lista)
	fmt.Println("Removendo do índice 3:", lista.RemoverIndice(3), lista)
	fmt.Println("Removendo do índice 0:", lista.RemoverIndice(0), lista)
	fmt.Println("Removendo do fim:", lista.RemoverIndice(len(lista)-1), lista)
}
