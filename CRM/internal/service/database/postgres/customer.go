package postgres

//package CRM/internal/service/database/postgres/dbmodels is not in std (C:\Program Files\Go\src\CRM\internal\service\database\postgres\dbmodels) (compile)go-staticcheck
import (
	"CRM/internal/service/database/dbmodels"
	"fmt"

	"github.com/google/uuid"
)



func GetCustomers(memberKey uuid.UUID) (records []dbmodels.Customer, err error) {



	 sql := `SELECT customer_user_key, full_name, email, phonenumber
	FROM public.customer`;

	temp,_ := env.db.Query(sql)

	fmt.Println(temp,"Temp")

	// postgres.Open()


	// rows , err := db.query(sql)


	fmt.Println(records)

	return records, nil
}
