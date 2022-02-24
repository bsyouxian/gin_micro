package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
	"github.com/gin-gonic/gin"
)

type name struct {
	
}
var DB *gorm.DB

func Database(connString string)  {
	db,err:=gorm.Open("mysql",connString)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	if gin.Mode()=="release"{
		db.LogMode(false)
	}
	db.SingularTable(true)
	//设置连接池
	db.DB().SetMaxIdleConns(20)
	//打开
	db.DB().SetMaxIdleConns(100)
	//超市
	db.DB().SetConnMaxLifetime(time.Second*30)
	DB = db
	migration()
}