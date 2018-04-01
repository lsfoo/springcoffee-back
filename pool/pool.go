package pool

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"encoding/json"
)

var db *sql.DB

func init() {

	var err error
	db, err = sql.Open("mysql", "root:lsf000000@/cxj")
	if err != nil {
		panic(err)
	}
	db.Ping()
}
func DsToJSON(sqlString string) (string, error) {

	stmt, err := db.Prepare(sqlString)
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	//fmt.Println(string(jsonData))
	return string(jsonData), nil
}
