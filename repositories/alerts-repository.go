package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vivasaayi/cloudrover/dal"
)

type Alerts struct {
	Id             int64
	DateHappened   int
	DeviceName     string
	Host           string
	AlertType      string
	Payload        string
	Priority       string
	SourceTypeName string
	Text           string
	Title          string
	Url            string
}

type Report struct {
	Id           int64
	DateHappened int
	Name         string
	Report       string
	ParsedJson   interface{}
}

func GetAllAlerts() []Alerts {
	db, err := sql.Open("mysql", dal.GetMySqlConnectionString())
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	query := `
		SELECT 
			id, date_happened, device_name, alert_type, title, url,
			host, payload, priority, text
		FROM alerts
	`

	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	result := []Alerts{}

	for rows.Next() {
		a := Alerts{}
		// get RawBytes from data
		err = rows.Scan(
			&a.Id, &a.DateHappened, &a.DeviceName, &a.AlertType, &a.Title,
			&a.Url, &a.Host, &a.Payload, &a.Priority, &a.Text,
		)

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
	db, err := sql.Open("mysql", dal.GetMySqlConnectionString())
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

	_, err = db.Query(query,
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
}

func InsertDataDogReport(name string, dateHappened int, reportJson string) {
	db, err := sql.Open("mysql", dal.GetMySqlConnectionString())
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for reading data
	query := `
		INSERT INTO past_reports (name, date_happened, report) 
		values(?, ?, ?)
	`

	fmt.Println("Executing Query")

	_, err = db.Query(query,
		name,
		dateHappened,
		reportJson,
	)

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

func GetDataDogReport(name string, pj interface{}) []Report {
	db, err := sql.Open("mysql", dal.GetMySqlConnectionString())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	query := `
		select * from cloudrover.past_reports
		where name = '` + name + `'
		order by date_happened desc
		limit 1	
	`

	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	result := []Report{}

	for rows.Next() {
		r := Report{}
		// get RawBytes from data
		err = rows.Scan(
			&r.Id, &r.DateHappened, &r.Name, &r.Report,
		)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		err := json.Unmarshal([]byte(r.Report), &pj)

		if err != nil {
			fmt.Println(err)

		} else {
			r.ParsedJson = pj
		}

		result = append(result, r)
	}

	return result
}
