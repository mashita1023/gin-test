package models

import (
			 "fmt"
			 "github.com/jinzhu/gorm"
			 _ "github.com/go-sql-driver/mysql"
)

const (
			DB = "mysql"
			User = "root"
			Pass = "pwd"
			Protocol = "tcp(127.0.0.1:5432)"
			Name = "test"
)

func SetupGorm() *gorm.DB {
		 connect := fmt.Sprintf("%s:%s@%s/%s", User, Pass, Protocol, Name)
		 db, err := gorm.Open(DB, connect)

		 if err != nil {
		 		fmt.Print(err.Error())
		 }

		 return db
}
