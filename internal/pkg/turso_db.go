package pkg

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/tursodatabase/go-libsql"
)

func InitDB() (*sql.DB, error) {
	dbName := os.Getenv("TURSO_DB_NAME")
	primaryUrl := os.Getenv("TURSO_DATABASE_URL")
	authToken := os.Getenv("TURSO_AUTH_TOKEN")
	dir, err := os.MkdirTemp("", "libsql-*")
	if err != nil {
		log.Fatal("Error creating temporary directory:", err)
		os.Exit(1)
		return nil, errors.New("Error creating temporary directory")
	}

	dbPath := filepath.Join(dir, dbName)

	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, primaryUrl,
		libsql.WithAuthToken(authToken),
	)
	if err != nil {
		log.Fatal("Error creating connector:", err)
		return nil, errors.New("Error creating connector")
	}

	db := sql.OpenDB(connector)
	err = createTables(db)

	if err != nil {
		return nil, errors.New("Error creating tables")
	}

	return db, nil
}

func createTables(db *sql.DB) error {

	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		last_name TEXT,
		password TEXT NOT NULL
	);
`)
	if err != nil {
		log.Fatal("Error creating users table:", err)
		return err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS keys (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		pass TEXT NOT NULL,
		alias TEXT UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		valid_until TIMESTAMP,
		user_id INTEGER NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
`)
	if err != nil {
		log.Fatal("Error creating keys table:", err)
		return err
	}

	return nil
}
