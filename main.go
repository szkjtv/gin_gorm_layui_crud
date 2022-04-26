package main

import (
	"time"

	"example.com/m/v2/mysql"
	"example.com/m/v2/router"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	mysql.Dbinit()
	router.Router()
	for {
		time.Sleep(0 * time.Microsecond)
	}
}
