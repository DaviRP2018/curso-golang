package main

import "fmt"

func init() {
	// convenção em go, todo arquivo pode ter uma função init()
	// Todas as init() serão executadas em ORDEM ALFABÉTICA do
	// nome dos arquivos antes da função main()
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
