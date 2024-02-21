package datatable

import "CRM/internal/service/database/dbmodels"

type DataTable struct {
	THeader THeader
	TRow    TRow

}

type THeader struct{
	Columns []DataTableColumn

}

type TRow struct {
	Rows []DataTableRow


}

type DataTableColumn struct {
	Name string
	Show bool

}

type DataTableRow struct {
	Cells DataTableCells

}

type DataTableCells struct{

}

func MapRecordsToDatatable([]dbmodels.Customer) DataTable {

	var dataTable = DataTable{}



	return dataTable
}