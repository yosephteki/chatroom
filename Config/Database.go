package Config

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	host     string
	port     string
	user     string
	dbname   string
	password string
	sslmode  string
}

func BuildConfig() *DBConfig {
	dbConfig := DBConfig{
		// host:     os.Getenv("DB_host"),
		// port:     os.Getenv("DB_port"),
		// user:     os.Getenv("DB_user"),
		// password: os.Getenv("DB_password"),
		// dbname:   os.Getenv("DB_name"),
		// sslmode:  os.Getenv("DB_sslmode"),
		host:     "ec2-54-166-167-192.compute-1.amazonaws.com",
		port:     "5432",
		user:     "nrojvokwxfufac",
		password: "25bdb986f6d7469052ab3468b88b900886a70a510aaf19d5b580b65bba36b624",
		dbname:   "d1ihg08luvjgbb",
		sslmode:  "require",
	}
	return &dbConfig
}

func DBURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.host,
		dbConfig.port,
		dbConfig.user,
		dbConfig.password,
		dbConfig.dbname,
		dbConfig.sslmode,
	)
}
