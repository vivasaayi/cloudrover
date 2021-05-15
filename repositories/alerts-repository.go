package repositories

import (
	"database/sql"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	_ "github.com/go-sql-driver/mysql"
)

type Alerts struct {
	Id   int64
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

func getStringValue(val *string) string {
	if val == nil {
		return ""
	}

	return *val
}

func InsertDataDogAlert(event *datadog.Event, tagsJson string) {
	db, err := sql.Open("mysql", "root:root@/cloudrover")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for reading data
	fmt.Println(*event.Id)
	query := `
		INSERT INTO alerts (
			id, date_happened, device_name, alert_type, 
			title, url, host, payload, priority, text, tags) 
		values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	fmt.Println("Executing Query")

	rows, err := db.Query(query,
		*event.Id,
		*event.DateHappened,
		getStringValue(event.DeviceName),
		getStringValue((*string)(event.AlertType)),
		getStringValue(event.Title),
		getStringValue(event.Url),
		getStringValue(event.Host),
		getStringValue(event.Payload),
		getStringValue((*string)(event.Priority)),
		getStringValue(event.Text),
		tagsJson,
	)

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
}
