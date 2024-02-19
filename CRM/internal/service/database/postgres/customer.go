package postgres

import (
	models "CRM/dbmodels"

	"fmt"

	"github.com/google/uuid"
)



func (env *Env) GetCustomers(memberKey uuid.UUID) (records []models.Customer, err error) {



	 sql := `SELECT customer_user_key, full_name, email, phonenumber
	FROM public.customer`;

	temp,_ := env.db.Query(sql)

	fmt.Println(temp,"Temp")

	// postgres.Open()


	// rows , err := db.query(sql)


	fmt.Println(records)

	return records, nil
}
