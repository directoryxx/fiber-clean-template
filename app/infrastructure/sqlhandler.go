package infrastructure

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/directoryxx/fiber-clean-template/database/gen"
	"github.com/directoryxx/fiber-clean-template/database/gen/migrate"
	_ "github.com/go-sql-driver/mysql"
)

// A SQLHandler belong to the infrastructure layer.
type SQLHandler struct {
	Conn *sql.DB
}

// A Tx belong to the infrastructure layer.
type Tx struct {
	Tx *sql.Tx
}

// A Result belong to the infrastructure layer.
type Result struct {
	Result sql.Result
}

// A Row belong to the infrastructure layer.
type Row struct {
	Rows *sql.Rows
}

// NewSQLHandler returns connection and methos which is related to database handling.
func NewSQLHandler() (*gen.Client, error) {
	client, err := gen.Open(os.Getenv("DB_TYPE"), os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run migration.
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client, nil
}
