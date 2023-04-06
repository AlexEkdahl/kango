package repository

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlexEkdahl/kango/config"
	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
)

type DAO interface {
	NewTaskQuery() TaskQuery
	NewTaskMutation() TaskMutation
}

type dao struct{}

var DB *sql.DB

func pgQb() squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(DB)
}

func NewDAO() DAO {
	return &dao{}
}

func (d *dao) NewTaskQuery() TaskQuery {
	return &taskQuery{}
}

func (d *dao) NewTaskMutation() TaskMutation {
	return &taskMutation{}
}

func NewDB(c config.Config) error {
	var err error

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	DB, err = createPostgresDB(dbUser, dbPassword, dbHost, dbPort, dbName)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	return nil
}

func NewLocalDB(c config.Config) error {
	var err error

	DB, err = createSQLiteDB(c.DBPath, c.DBName)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	return nil
}

func createSQLiteDB(dbPath, dbName string) (*sql.DB, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to find user home directory: %v", err)
	}

	// Replace ~ symbol with the user's home directory
	dbPath = strings.Replace(dbPath, "~", homeDir, 1)

	dbFilePath := filepath.Join(dbPath, dbName)
	shouldCreateTable := false

	if _, err := os.Stat(dbFilePath); os.IsNotExist(err) {
		err := os.MkdirAll(dbPath, 0o755)
		if err != nil {
			return nil, fmt.Errorf("failed to create the database directory: %v", err)
		}

		_, err = os.Create(dbFilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to create SQLite database file: %v", err)
		}
		shouldCreateTable = true
	}

	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return nil, err
	}

	if shouldCreateTable {
		if err = createTaskTable(db); err != nil {
			return nil, err
		}
	}

	return db, nil
}

func createTaskTable(db *sql.DB) error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS task (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    status INTEGER NOT NULL,
	    subject VARCHAR(255) NOT NULL,
	    description TEXT
	);
	`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create task table: %v", err)
	}

	return nil
}

func createPostgresDB(dbUser, dbPassword, dbHost, dbPort, dbName string) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("postgres", connStr)
	return db, err
}
