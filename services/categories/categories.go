package categories

import "time"

type Categories struct {
	ID          string     `db:"news_category_id" json:"id"`
	Name        string     `db:"name" json:"name"`
	Prefix      *string    `db:"category_prefix" json:"prefix"`
	Valid       string     `db:"valid_flag" json:"valid"`
	CreatedBy   string     `db:"created_by" json:"created_by"`
	CreatedDate time.Time  `db:"created_date" json:"created_date"`
	UpdatedBy   *string    `db:"updated_by" json:"updated_by"`
	UpdatedDate *time.Time `db:"updated_date" json:"updated_date"`
}
