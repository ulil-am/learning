package models

import (
	bookStruct "learning/advanced-cloud-native-go-gin/Frameworks/Gin-Web/structs"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq" //PostgreSQL Driver
)

var ormObject orm.Ormer

// ConnectToDb - Initializes the ORM and Connection to the postgres DB
func ConnectToDb() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=postgres dbname=postgres host=192.168.99.100:5432 sslmode=disable")
	orm.RegisterModel(new(bookStruct.Book))
	ormObject = orm.NewOrm()
}

// GetOrmObject - Getter function for the ORM object with which we can query the database
func GetOrmObject() orm.Ormer {
	return ormObject
}
