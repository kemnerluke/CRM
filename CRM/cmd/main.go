package main

import (
	log "github.com/sirupsen/logrus"

	"CRM/internal/service/database/postgres"
	"CRM/internal/service/datatable"

	_ "github.com/lib/pq"
)

func main() {


	log.Info("---Starting Main---")


  records,_ := postgres.GetCustomers()

  datatable.MapRecordsToDatatable(records)


}