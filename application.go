package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var (
	host      = os.Getenv("RDS_HOSTNAME")           //localhost
	port, err = strconv.Atoi(os.Getenv("RDS_PORT")) //5432
	user      = os.Getenv("RDS_USERNAME")           //postgres
	password  = os.Getenv("RDS_PASSWORD")           //password
	dbname    = os.Getenv("RDS_DB_NAME")            //go_backend_development
)

func main() {
	db := connect()

	rows, err := db.Query("SELECT user_name FROM users LIMIT 10;")
	if err != nil {
		fmt.Printf("failed to execute query. Error: %s", err)
	}
	defer rows.Close()

	var usernames []string
	for rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			fmt.Printf("failed to scan rows. Error: %s", err)
		}
		usernames = append(usernames, username)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", usernames)
	})
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func connect() *sql.DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
