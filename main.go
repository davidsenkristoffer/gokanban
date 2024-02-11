package main

import (
	"database/sql"
	"fmt"
	"gokanban/db"
	"gokanban/route"

	"github.com/labstack/echo/v4"
)

type kanbanContext struct {
	echo.Context
	db *sql.DB
}

func main() {
	database, err := db.Connect()
	catch(err)

	e := route.Init(database)
	e.Logger.Fatal(e.Start(":1337"))
}

func catch(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
