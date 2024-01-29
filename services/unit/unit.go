package unit

import (
	"time"
)

type Unit struct {
	ID          string     `db:"unit_id" json:"id"`
	Name        string     `db:"unit_name" json:"unit_name"`
	Prefix      string     `db:"unit_prefix" json:"prefix"`
	Valid       string     `db:"valid_flag" json:"valid"`
	CreatedDate *time.Time `db:"created_date" json:"created_date"`
	CreatedBy   string     `db:"created_by" json:"created_by"`
	UpdatedDate *time.Time `db:"updated_date" json:"updated_date"`
	UpdatedBy   *string    `db:"updated_by" json:"updated_by"`
}

func List() {

}
