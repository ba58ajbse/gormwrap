package database

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/testcontainers/testcontainers-go"
	test_mysql "github.com/testcontainers/testcontainers-go/modules/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var mysqlContainer *test_mysql.MySQLContainer
var ctx context.Context

func TestContainerInit() *SQLHandler {
	var err error
	ctx = context.Background()
	mysqlContainer, err = test_mysql.RunContainer(ctx,
		testcontainers.WithImage("mysql:8.0.36"),
		test_mysql.WithDatabase("sample"),
		test_mysql.WithUsername("root"),
		test_mysql.WithPassword("password"),
	)

	if err != nil {
		panic(err.Error())
	}

	connectionString, _ := mysqlContainer.ConnectionString(ctx)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err.Error())
	}

	return &SQLHandler{
		DB: db,
	}
}

func Terminate() {
	err := mysqlContainer.Terminate(ctx)
	if err != nil {
		fmt.Printf("failed to terminate test container: %v\n", err)
	}
}
