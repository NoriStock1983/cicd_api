package companymst_dto

import (
	"time"
)

/* companymstの全項目取得 */
type Companymst struct {
	Companycd        string    `json:"companycd"`
	Company_name     string    `json:"company_name"`
	Company_nm_short string    `json:"company_nm_short"`
	Active_from_date time.Time `json:"active_from_date"`
	Active_to_date   time.Time `json:"active_to_date"`
	Created_by       string    `json:"created_by"`
	Created_date     time.Time `json:"created_date"`
	Updated_by       string    `json:"updated_by"`
	Updated_date     time.Time `json:"updated_date"`
	Update_counter   uint      `json:"update_counter"`
}
