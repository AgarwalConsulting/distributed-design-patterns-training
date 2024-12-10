package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var leaderDB *sql.DB
var followerDB *sql.DB

func init() {
	// Connect to the leader DB
	var err error
	leaderDB, err = sql.Open("postgres", "user=leader password=leaderpass dbname=mydb host=leader-db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the follower DB
	followerDB, err = sql.Open("postgres", "user=follower password=followerpass dbname=mydb host=follower-db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func readFromFollower(query string) {
	rows, err := followerDB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var value string
		if err := rows.Scan(&id, &value); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Follower read: %d, %s\n", id, value)
	}
}

func writeToLeader(query string, args ...interface{}) {
	_, err := leaderDB.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Write operation to the leader
	writeToLeader("INSERT INTO my_table (id, value) VALUES ($1, $2)", 1, "value1")

	// Read operation from the follower
	readFromFollower("SELECT id, value FROM my_table")
}
