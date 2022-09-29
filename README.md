# go_restapi_gin
Belajar CRUD API dengan framework gin dengan menggunakan gorm mysql


#jalankan perintah "go run main.go" untuk menjalankan aplikasi

#disini saya membuat database dengan nama 'go_restapi_gin' dengan table 'product'

func ConnectDatabase() {

	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_restapi_gin"))

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})

	DB = database
}
