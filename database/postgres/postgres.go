package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // import postgers dependencies
	"log"
	"os"
	"os/signal"
	"syscall"
)

// NewConnection establishes a new connection with Postgres
func NewConnection(user, password, host, dbname, port string) *sql.DB {
	source := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	db, err := sql.Open("postgres", source)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	go close(db)

	return db
}

func close(db *sql.DB) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs

	fmt.Println("Gracefully closing db connection")
	if err := db.Close(); err != nil {
		log.Printf("Failed to properly close database connection: %v", err)
	}
}
