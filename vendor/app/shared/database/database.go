package database

import (
	//"encoding/json"
	"fmt"
	"log"
	//"time"
	//"database/sql"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// SQL wrapper
	SQL *gorm.DB
)

type MySQLInfo struct {
	Username  string
	Password  string
	Name      string
	Hostname  string
	Port      int
	Parameter string
}

// DSN returns the Data Source Name
func DSN(ci MySQLInfo) string {
	// Example: root:@tcp(localhost:3306)/test
	return ci.Username +
		":" +
		ci.Password +
		"@tcp(" +
		ci.Hostname +
		":" +
		fmt.Sprintf("%d", ci.Port) +
		")/" +
		ci.Name + ci.Parameter
}

func Connect(d MySQLInfo) {
	var err error
	if SQL, err = gorm.Open("mysql", "root@tcp(localhost:3306)/todo_dev"); err != nil {
		log.Println("SQL Driver Error", err)
	}

}
