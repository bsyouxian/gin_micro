package model

func migration()  {
	DB.Set("gorm:table_options","charset=utf8&mb4").AutoMigrate(&User{})
}