package infrastructure

import (
	"context"
	"log"
	"os"

	"github.com/directoryxx/fiber-clean-template/database/gen"
	"github.com/directoryxx/fiber-clean-template/database/gen/migrate"
	_ "github.com/go-sql-driver/mysql"
)

var client *gen.Client

// NewSQLHandler returns connection and methos which is related to database handling.
func NewSQLHandler(ctx context.Context) (*gen.Client, error) {
	client, err := gen.Open(os.Getenv("DB_TYPE"), os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}

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

func CloseSQLHandler() {
	client.Close()
}
