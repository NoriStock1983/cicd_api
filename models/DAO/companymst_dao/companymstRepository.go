package companymstdao

import (
	"cicd_api/commons/dbaccess"
	"cicd_api/models/DTO/companymst_dto"
)

func GetAllCompanymst() ([]companymst_dto.Companymst, error) {
	var companymst []companymst_dto.Companymst
	db, error := dbaccess.DBAccess()

	if error != nil {
		return nil, error
	}
	defer db.Close()

	rows, error := db.Query("SELECT * FROM companymst")

	if error != nil {
		return nil, error
	}
	defer rows.Close()

	for rows.Next() {
		company := &companymst_dto.Companymst{}
		rows.Scan(&company.Companycd, &company.Company_name, &company.Company_nm_short, &company.Active_from_date, &company.Active_to_date, &company.Created_by, &company.Created_date, &company.Updated_by, &company.Updated_date, &company.Update_counter)

		companymst = append(companymst, *company)
	}
	return companymst, nil
}
