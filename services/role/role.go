package role

import (
	"time"
	"tnc-masters/pkg/database"
)

type Role struct {
	ID          string     `db:"role_id" json:"id"`
	Name        string     `db:"name" json:"name"`
	Key         string     `db:"key" json:""`
	Valid       string     `db:"valid_flag" json:"valid"`
	CreatedDate time.Time  `db:"created_date" json:"created_date"`
	CreatedBy   string     `db:"created_by" json:"created_by"`
	UpdatedDate *time.Time `db:"updated_date" json:"updated_date"`
	UpdatedBy   *string    `db:"updated_by" json:"updated_by"`
}

func List() (interface{}, error) {
	query := "select role_id, name, `key`, valid_flag, created_by, created_date, updated_by, updated_date "
	query += "from tnc_role"

	var rs []Role
	if err := database.DB.Select(&rs, query); err != nil {
		return nil, err
	}
	return rs, nil
}
