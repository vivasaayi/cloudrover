package repositories

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Alerts struct {
	Id   int
	Date int
}

func GetAllAlerts() []Alerts {
	db, err := sql.Open("mysql", "root:root@/cloudrover")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for reading data
	rows, err := db.Query("SELECT id, date_happened FROM alerts")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	result := []Alerts{}

	for rows.Next() {
		a := Alerts{}
		// get RawBytes from data
		err = rows.Scan(&a.Id, &a.Date)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		result = append(result, a)
	}

	return result
}
