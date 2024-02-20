package main

import (
	log "github.com/sirupsen/logrus"

	"CRM/internal/service/database/postgres"

	_ "github.com/lib/pq"
)

func main() {


	log.Info("---Starting Main---")


  postgres.GetCustomers()




}