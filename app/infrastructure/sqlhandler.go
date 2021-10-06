package infrastructure

import (
	"context"
	"log"
	"os"
	"time"

	"entgo.io/ent/dialect/sql"

	"github.com/casbin/casbin/v2"
	"github.com/directoryxx/fiber-clean-template/database/gen"
	"github.com/directoryxx/fiber-clean-template/database/gen/migrate"
	_ "github.com/go-sql-driver/mysql"
)

var client *gen.Client
var dsn string
var driver string

// NewSQLHandler returns connection and methos which is related to database handling.
func NewSQLHandler(ctx context.Context) (*gen.Client, *casbin.Enforcer, error) {
	driver = os.Getenv("DB_TYPE")
	dsn = os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?parseTime=true"

	client, err := gen.Open(driver, dsn)
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

	client.Close()
	enforcer := CasbinLoad(driver, dsn)

	// fmt.Println(enforcer)

	return nil, enforcer, nil
}

func Open() (*gen.Client, error) {
	driver = os.Getenv("DB_TYPE")
	dsn = os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?parseTime=true"

	drv, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	// defer db.Close()
	return gen.NewClient(gen.Driver(drv)), nil
}

func CloseSQLHandler() {
	client.Close()
}
