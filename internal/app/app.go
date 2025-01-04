package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/tursodatabase/go-libsql"

	"github.com/gin-gonic/gin"
)

func Start() {
	loadDependencies()

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Run(":8080")
}

// momentario
func loadDependencies() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	initDb()
}

func initDb() *sql.DB {
	dbName := os.Getenv("TURSO_DB_NAME")
	primaryUrl := os.Getenv("TURSO_DATABASE_URL")
	authToken := os.Getenv("TURSO_AUTH_TOKEN")
	dir, err := os.MkdirTemp("", "libsql-*")
	if err != nil {
		fmt.Println("Error creating temporary directory:", err)
		os.Exit(1)
	}
	defer os.RemoveAll(dir)

	dbPath := filepath.Join(dir, dbName)

	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, primaryUrl,
		libsql.WithAuthToken(authToken),
	)
	if err != nil {
		fmt.Println("Error creating connector:", err)
		os.Exit(1)
	}
	defer connector.Close()
	db := sql.OpenDB(connector)
	return db
}
