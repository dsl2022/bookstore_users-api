package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host = "mysql_users_host"
	mysql_users_schema = "mysql_users_schema"
)

func goDotEnvVariable(key string)string{
	var err =godotenv.Load(".env")
	if err!=nil{
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

var(
	Client *sql.DB
	username = goDotEnvVariable(mysql_users_username)
	password = goDotEnvVariable(mysql_users_password)
	host = goDotEnvVariable(mysql_users_host)
	schema = goDotEnvVariable(mysql_users_schema)
)


func init(){
	
	datasourceName:=fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
	username,password,host,schema)// "root","pass","127.0.0.1","users_db")
	var err error
	Client,err = sql.Open("mysql",datasourceName)
	if err!=nil{
		panic(err)
	}
	if err:= Client.Ping();err!=nil{
		panic(err)
	}	
	log.Println("database successful configured")
}