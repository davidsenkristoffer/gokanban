package main

import (
	"fmt"
	"gokanban/db"
	"gokanban/route"
	"gokanban/services"
)

func main() {
	database, err := db.Connect()
	catch(err)

	ps := new(services.ProjectService)

	e := route.Init(database, *ps)
	e.Logger.Fatal(e.Start(":1337"))
}

func catch(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
