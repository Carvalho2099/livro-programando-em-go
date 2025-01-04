package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	palavras := os.Args[1:]
	estatisticas := colherEstatisticas(palavras)

	imprimir(estatisticas)
}

func colherEstatisticas(palavras []string) map[string]int {
	estastisticas := make(map[string]int)

	for _, palavra := range palavras {
		inicial := strings.ToUpper(string(palavra[0]))
		contador, encontrado := estastisticas[inicial]

		if encontrado {
			estastisticas[inicial] = contador + 1
		} else {
			estastisticas[inicial] = 1
		}
	}
	return estastisticas
}

func imprimir(estatisticas map[string]int) {
	fmt.Println("Contagem de palavras iniciadas em cada letra:")

	for inicial, contador := range estatisticas {
		fmt.Printf("%s = %d\n", inicial, contador)
	}
}
