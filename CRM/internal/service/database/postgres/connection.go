package postgres

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)



type Env struct {
    db *sql.DB
}


func Open() (*sql.DB, error) {
	log.Info("Connecting to the Database")
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "kemner123"
		dbname   = "postgres"
	  )

	  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	  "password=%s dbname=%s sslmode=disable",
	  host, port, user, password, dbname)

	  db, err := sql.Open("postgres", psqlInfo)
if err != nil {
  panic(err)
}
defer db.Close()

err = db.Ping()

if err != nil {
  return nil,err
}

log.Info("Successfully connected to the database")

return db,nil


}
