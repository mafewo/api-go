package database

import (
	"database/sql"
	"log"
	//importa el driver de mysql
	_ "github.com/go-sql-driver/mysql"
	/*"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"*/)

var db *sql.DB

//var db *gorm.DB

//ConnString es el struct que contiene el string DNS de conexion
type ConnString struct {
	conn  string
	table string
}

//NewConn retorna el puntero de ConnString seteado
func NewConn(conn string, table string) *ConnString {
	return &ConnString{
		conn:  conn,
		table: table,
	}
}

//InitializeDatabase abre la conexion mysql
func (c *ConnString) InitializeDatabase() {
	log.Println("Initializing database ...")
	var err error
	db, err = sql.Open("mysql", c.conn)
	/*gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return c.table
	}
	db, err = gorm.Open("mysql", c.conn)
	db.SingularTable(true)*/
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
}
