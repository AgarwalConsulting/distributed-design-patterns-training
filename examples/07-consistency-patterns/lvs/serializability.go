package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

func main() {
	// Set up database connection
	conn, err := pgx.Connect(context.Background(), "postgresql://user:password@localhost:5432/mydb")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	// Start a transaction
	tx, err := conn.Begin(context.Background())
	if err != nil {
		log.Fatalf("Unable to begin transaction: %v\n", err)
	}
	defer tx.Rollback(context.Background()) // Ensure rollback on error

	// Execute some database operations
	_, err = tx.Exec(context.Background(), "UPDATE accounts SET balance = balance - 100 WHERE id = $1", 1)
	if err != nil {
		log.Fatalf("Failed to update account 1: %v\n", err)
	}

	_, err = tx.Exec(context.Background(), "UPDATE accounts SET balance = balance + 100 WHERE id = $1", 2)
	if err != nil {
		log.Fatalf("Failed to update account 2: %v\n", err)
	}

	// Commit the transaction, ensuring atomicity and serializability
	if err := tx.Commit(context.Background()); err != nil {
		log.Fatalf("Failed to commit transaction: %v\n", err)
	}

	fmt.Println("Transaction committed successfully")
}
