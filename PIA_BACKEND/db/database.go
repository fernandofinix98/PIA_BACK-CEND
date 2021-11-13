package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//username:password@tcp(localhost:3306)/database
const url = "root:022528Sepmaydic@tcp(localhost:3306)/pia_b"

//Guarda la conexion
var db *sql.DB

//realiza la conexion
func Connect() {

	conection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion exitosa")
	db = conection
}

//cierra la conexion
func Close() {
	db.Close()
}

//Verificar la conexion
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

//Verificar si una tabla existe o no
func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Error", err)
	}
	return rows.Next()
}

//Crea una tabla
func CreateTable(schema string, name string) {
	if !ExistsTable(name) {
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println(err)
		}

	}
}

//Reiniciar el registro de una tabla
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)

}

//Polimorfismo de Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	resust, err := db.Exec(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}
	return resust, err
}

//Polimorfismo de Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	rows, err := db.Query(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}
	return rows, err

}
