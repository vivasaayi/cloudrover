package dal

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vivasaayi/cloudrover/utililties"
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

func GetMySqlConnectionString() string {
	hostName := utililties.GetStringEnvVar("MYSQL_HOSTNAME", "localhost", false)
	port := utililties.GetStringEnvVar("MYSQL_PORT", "3306", false)
	userName := utililties.GetStringEnvVar("MYSQL_USERNAME", "", true)
	password := utililties.GetStringEnvVar("MYSQL_PASSWORD", "", true)
	databaseName := utililties.GetStringEnvVar("MYSQL_DB_NAME", "cloudrover", false)

	host_port := fmt.Sprintf("%s:%s", hostName, port)

	if hostName == "localhost" && port == "3306" {
		host_port = ""
	}

	cs := fmt.Sprintf("%s:%s@%s/%s", userName, password, host_port, databaseName)

	fmt.Println(cs)

	return cs
}
