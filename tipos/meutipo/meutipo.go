package main

import "fmt"

type nota float64
type meuint int

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	switch {
	case n.entre(9.0, 10.0):
		return "A"
	case n.entre(7.0, 8.99):
		return "B"
	case n.entre(5.0, 7.99):
		return "C"
	case n.entre(3.00, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.9))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))

	oi := 1.7627662347263476237
	fmt.Println(meuint(oi))
}
