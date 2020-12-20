package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "bkdir"
	password = "password"
	dbname   = "go_backend_development"
)

func main() {
	db := connect()
	fmt.Println("db")
	fmt.Println(db)

	rows, err := db.Query("SELECT user_name FROM public.users LIMIT 10")
	fmt.Println("rows:")
	fmt.Println(rows)
	if err != nil {
		fmt.Printf("failed to execute query. Error: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			fmt.Printf("failed to scan rows. Error: %s", err)
		}
		fmt.Println(username)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	})
	//log.Fatal(http.ListenAndServe(":5000", nil))
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
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
