package server

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"putrafirman.com/playground/privy-code-test/service"
)

type Server struct {
}

func (s *Server) StartServer() {

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	//setup db connection
	db := dbConn()
	e.Logger.Info(db.Ping())
	defer db.Close()

	//migrate
	migrate(db)

	//handler
	e.GET("/ping", service.GetPingAndPong(db))
	e.GET("/cakes", service.GetAllCake(db))
	e.GET("/cakes/:id", service.GetCakeDetail(db))
	e.POST("/cakes", service.CreateCake(db))
	e.PUT("/cakes/:id", service.UpdateCake(db))
	e.DELETE("/cakes/:id", service.DeleteCake(db))

	e.Logger.Info("Starting echo server...")
	e.Logger.Fatal(e.Start(":8080"))

}

func dbConn() (db *sql.DB) {
	//must be in env var for production use. for intial will be hardcoded
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "!1Password"
	dbName := "cakestore"
	dbHost := "x-internal.pyx.my.id:3306"

	dbParam := dbUser + ":" + dbPass + "@tcp(" + dbHost + ")/" + dbName + "?parseTime=true"
	db, err := sql.Open(dbDriver, dbParam)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS cakes(
        id INTEGER NOT NULL AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL DEFAULT '',
		description VARCHAR(255) NOT NULL DEFAULT '',
		rating float(3) NOT NULL DEFAULT 0,
		image VARCHAR(255) NOT NULL DEFAULT '',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY(id)
    );
    `

	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}
