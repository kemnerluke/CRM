package main

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	"CRM/internal/service/database/postgres"

	_ "github.com/lib/pq"
)

func main() {

	type Env struct {
		db *sql.DB
	}
 

//  db , err:= postgres.Open()

//  env := &Env{db: db}

//  log.Info(env)


  postgres.GetCustomers()


 if err != nil {
	log.Fatal(err)
}

}