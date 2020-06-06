package test

import (
	"context"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/qreasio/go-starter-kit/pkg/log"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	dbUsername string = "root"
	dbPassword string = "password"
	dbName     string = "test"
)

var (
	db *sqlx.DB
)

func SetupMySQLDBContainer(logger log.Logger) (func(), *sqlx.DB, error) {
	logger.Info("setup MySQL Container")
	ctx := context.Background()

	seedDataPath, err := os.Getwd()
	if err != nil {
		logger.Errorf("error get working directoryr %s", err)
		panic(fmt.Sprintf("%v", err))
	}
	mountPath := seedDataPath + "/../../test/integration"

	req := testcontainers.ContainerRequest{
		Image:        "mysql:latest",
		ExposedPorts: []string{"3306/tcp", "33060/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": dbPassword,
			"MYSQL_DATABASE":      dbName,
		},
		BindMounts: map[string]string{
			mountPath: "/docker-entrypoint-initdb.d",
		},
		WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL"),
	}

	mysqlC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		logger.Errorf("error starting mysql container %s", err)
		panic(fmt.Sprintf("%v", err))
	}

	closeContainer := func() {
		logger.Info(">>>>>> terminating container")
		err := mysqlC.Terminate(ctx)
		if err != nil {
			logger.Errorf("error terminating mysql container %s", err)
			panic(fmt.Sprintf("%v", err))
		}
	}

	host, _ := mysqlC.Host(ctx)
	p, _ := mysqlC.MappedPort(ctx, "3306/tcp")
	port := p.Int()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=skip-verify&parseTime=true&multiStatements=true",
		dbUsername, dbPassword, host, port, dbName)

	db, err = sqlx.Connect("mysql", connectionString)
	if err != nil {
		logger.Info("error Connect db: %+v\n", err)
		return closeContainer, db, err
	}

	if err = db.Ping(); err != nil {
		logger.Infof("error pinging db: %+v\n", err)
		return closeContainer, db, err
	}

	return closeContainer, db, nil
}
