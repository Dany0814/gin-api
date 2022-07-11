package domain

import (
	"database/sql"
	"fmt"

	"github.com/Dany0814/gin-api/internal/platform/server"
	"github.com/Dany0814/gin-api/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "julio"
	dbPass = "julio"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "julio"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)

	srv := server.New(host, port, courseRepository)
	return srv.Run()
}
