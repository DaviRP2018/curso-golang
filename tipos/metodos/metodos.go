package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getFullName() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setFullName(fullName string) {
	parts := strings.Split(fullName, " ")
	p.nome = parts[0]
	p.sobrenome = parts[1]
}

func main() {
	p1 := pessoa{"Pedro", "Silva"}
	fmt.Println(p1.getFullName())

	p1.setFullName("Maria Pereira")
	fmt.Println(p1.getFullName())
}
