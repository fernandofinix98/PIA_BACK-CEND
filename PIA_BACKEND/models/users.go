package models

import "apirest/db"

type User struct {
	Id       int64  `json:"id"`
	Empresa  string `json:"empresa"`
	Grupo    string `json:"grupo"`
	Miembros string `json:"miembros"`
}
type Users []User

//Comando de mysql implementado en Go para crear una tabla
const UserSchema string = `CREATE TABLE catalogo (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	empresa VARCHAR(30) NOT NULL,
	grupo VARCHAR(30) NOT NULL,
	miembros VARCHAR(150),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

//contruir usuario
func NewUser(empresa, grupo, miembros string) *User {
	user := &User{Empresa: empresa, Grupo: grupo, Miembros: miembros}
	return user
}

//Crear usuario e inserta bd
func CreateUser(empresa, grupo, miembros string) *User {
	user := NewUser(empresa, grupo, miembros)
	user.Save()
	return user

}

//Insertar Registro
func (user *User) insert() {
	sql := "INSERT catalogo SET empresa=?, grupo=?, miembros=?"
	result, _ := db.Exec(sql, user.Empresa, user.Grupo, user.Miembros)
	user.Id, _ = result.LastInsertId()

}

//Listar todos los registros
func ListUsers() (Users, error) {
	sql := "SELECT id, empresa, grupo, miembros FROM catalogo"
	users := Users{}
	rows, err := db.Query(sql)
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Empresa, &user.Grupo, &user.Miembros)
		users = append(users, user)
	}
	return users, err
}

//Obtener un registro
func GetUser(id int) (*User, error) {
	user := NewUser("", "", "")

	sql := "SELECT id, empresa, grupo, miembros FROM catalogo WHERE id=?"
	if rows, err := db.Query(sql, id); err != nil {
		return nil, err
	} else {
		for rows.Next() {
			rows.Scan(&user.Id, &user.Empresa, &user.Grupo, &user.Miembros)
		}
	}
	return user, nil
}

//Actualizar registro
func (user *User) update() {
	sql := "UPDATE catalogo SET empresa=?, grupo=?, miembros=? WHERE id=?"
	db.Exec(sql, user.Empresa, user.Grupo, user.Miembros, user.Id)
}

//Guardar o editar registro
func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}
}

//Eliminar un registro
func (user *User) Delete() {
	sql := "DELETE FROM catalogo WHERE id=?"
	db.Exec(sql, user.Id)
}
