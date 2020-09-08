package main

import "fmt"

// Não tem operador ternário
// Exemplo: return nota >= 6.2 ? "Aprovado" : "Reprovado"
// Exemplo: return "Aprovado" if nota >= 6.2 else "Reprovado"
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.2))
}
