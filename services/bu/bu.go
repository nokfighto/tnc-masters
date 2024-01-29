package bu

import (
	"time"
	"tnc-masters/pkg/database"
)

type BU struct {
	ID          string     `db:"bu_id" json:"id"`
	Name        string     `db:"name" json:"name"`
	Description string     `db:"description" json:"description"`
	Valid       string     `db:"valid_flag" json:"valid"`
	CreatedDate time.Time  `db:"created_date" json:"created_date"`
	CreatedBy   string     `db:"created_by" json:"created_by"`
	UpdatedDate *time.Time `db:"updated_date" json:"updated_date"`
	UpdatedBy   *string    `db:"updated_by" json:"updated_by"`
}

func List() (interface{}, error) {
	query := `select bu_id, name, description, valid_flag, created_date, created_by, updated_date, updated_by
			  from tnc_business_unit`

	var b []BU
	if err := database.DB.Select(&b, query); err != nil {
		return nil, err
	}
	return b, nil
}
