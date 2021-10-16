package infrastructure

import (
	"context"
	"os"

	"github.com/casbin/casbin/v2"
	"github.com/directoryxx/fiber-clean-template/app/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn string
var driver string

// NewSQLHandler returns connection and methos which is related to database handling.
func NewSQLHandler(ctx context.Context) (*casbin.Enforcer, error) {
	driver = os.Getenv("DB_TYPE")
	dsn = os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&domain.User{}, &domain.Role{}, &domain.PhoneBook{})

	enforcer := CasbinLoad(driver, dsn)

	// fmt.Println(enforcer)

	return enforcer, nil
}

func Open() (gormConn *gorm.DB, err error) {
	driver = os.Getenv("DB_TYPE")
	dsn = os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?parseTime=true"

	drv, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Get the underlying sql.DB object of the driver.
	// db, errDB := drv.DB()
	// db.SetMaxIdleConns(10)
	// db.SetMaxOpenConns(100)
	// db.SetConnMaxLifetime(time.Hour)
	// defer db.Close()
	return drv, err
}
