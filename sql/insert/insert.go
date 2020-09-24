package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/curso")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO usuarios(nome) VALUES(?)")
	stmt.Exec("Maria")
	stmt.Exec("José")

	res, _ := stmt.Exec("Pedro")

	id, _ := res.LastInsertId()
	fmt.Printf("ID do novo Pedro %d\n", id)

	linhas, _ := res.RowsAffected()
	fmt.Printf("Linhas afetadas %d\n", linhas)
}
