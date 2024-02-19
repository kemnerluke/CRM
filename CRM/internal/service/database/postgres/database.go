package postgres

//package CRM/internal/service/database/postgres/dbmodels is not in std (C:\Program Files\Go\src\CRM\internal\service\database\postgres\dbmodels) (compile)go-staticcheck
import (
	"CRM/internal/service/database/postgres/dbmodels"

	"github.com/google/uuid"
)

type Database interface {
	GetCustomers(memberKey uuid.UUID) (records []dbmodels.Customer, err error)
}