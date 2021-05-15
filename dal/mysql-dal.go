package dal

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Alerts struct {
	id            int
	date_happened int
}

func ExecuteSelectQuery() {
	db, err := sql.Open("mysql", "root:root@db/cloudrover")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for reading data
	rows, err := db.Query("SELECT id, date_happened FROM alerts")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for rows.Next() {
		a := Alerts{}
		// get RawBytes from data
		err = rows.Scan(&a.id, &a.date_happened)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}
}
