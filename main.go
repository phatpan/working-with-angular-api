package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/phatpan/working-with-angular-api/income"
	"github.com/phatpan/working-with-angular-api/logs"
	"github.com/phatpan/working-with-angular-api/outcome"
	"github.com/phatpan/working-with-angular-api/user"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func main() {
	var env, port, dsn string
	flag.StringVar(&env, "env", "dev", "running environment")
	flag.StringVar(&port, "port", "9003", "running port number")
	flag.StringVar(&dsn, "dsn", "work:Work@1q2w3e4r5t@tcp(103.74.254.157:3306)/ohmymoney?parseTime=true", "datasource name")
	flag.Parse()

	db := makeDBConnection(dsn)
	defer db.Close()

	logger := makeLogger(env)
	fieldLogger := &logs.FieldLog{Logger: logger}

	e := echo.New()
	e.HideBanner = true

	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.POST, echo.PUT, echo.GET, echo.DELETE},
	}))

	user.NewHandler(e, db, fieldLogger)
	income.NewHandler(e, db, fieldLogger)
	outcome.NewHandler(e, db, fieldLogger)

	logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

func makeDBConnection(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	if err = db.Ping(); err != nil {
		panic(err.Error())
	}
	return db
}

func makeLogger(env string) (logger *logrus.Logger) {
	logger = logrus.New()
	logger.Out = os.Stdout
	if env == "dev" {
		logger.SetFormatter(&prefixed.TextFormatter{FullTimestamp: true})
		logger.SetLevel(logrus.DebugLevel)
		return
	}
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)
	return
}
