package postgres

//package CRM/internal/service/database/postgres/dbmodels is not in std (C:\Program Files\Go\src\CRM\internal\service\database\postgres\dbmodels) (compile)go-staticcheck
import (
	"CRM/internal/service/database/dbmodels"
)



func GetCustomers() ([]dbmodels.Customer, error) {


	var records []dbmodels.Customer

	for _,_ = range []int{1}{

	records = append(records, dbmodels.Customer{

		FirstName: "John",
		LastName: "Doe",
		PhoneNumber: 111-111-1110,
		Email: "JohnDoe@emaill.com",

	})
}




	return records, nil
}
