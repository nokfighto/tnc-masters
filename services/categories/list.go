package categories

import (
	"strings"
	"tnc-masters/pkg/database"
)

func List(identifier, valid string) (interface{}, error) {
	// identifier is id or name
	if strings.Trim(identifier, " ") == "" {
		var c []Categories
		err := database.DB.Select(&c, "select * from tnc_news_category")
		if err != nil {
			return nil, err
		} else {
			return c, nil
		}
	}

	var args []interface{}
	args = append(args, identifier)
	args = append(args, "%"+identifier+"%")
	query := `select news_category_id, name, category_prefix, valid_flag
	      ,created_by, created_date, updated_by, updated_date
	      from tnc_news_category 
		  where (news_category_id = ? or name like ?) `

	if valid != "" {
		query += "and valid_flag = ? "
		args = append(args, valid)
	}
	//fmt.Println(args...)
	var c []Categories
	err := database.DB.Select(&c, query, args...)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func TopLevel() {

}

func Category() {

}

func SubCategory() {
	//query := `select `
}
