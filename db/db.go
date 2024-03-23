package db

import (
	"database/sql"
	// "os"
	// "regexp"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/joho/godotenv"
	// "subosito.com/go/gotenv"
)

var (
	DB             *sql.DB = Open()
	projectDirName string  = "museid_admin"
)

func dsn() string {
	// re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	// cwd, _ := os.Getwd()
	// rootPath := re.Find([]byte(cwd))
	//
	// err := godotenv.Load(string(rootPath) + `/.env`)
	// if err != nil {
	// 	panic("Problem loading .env file")
	// }

	// return fmt.Sprintf(
	// 	"%s:%s@tcp(%s)/%s",
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASS"),
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_NAME"),
	// )

	return "root:mypass@tcp(127.0.0.1:3306)/museid"
}

func Open() *sql.DB {
	mydb, err := sql.Open("mysql", dsn())
	if err != nil {
		panic(err)
	}

	return mydb
}
