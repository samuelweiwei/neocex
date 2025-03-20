package initialize

import (
	"fmt"
	"neocex/v2/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormPgSQL() (*gorm.DB, error) {
	x := config.GeneralDB{Port: "5432", Path: "localhost", Username: "postgres", Password: "postgres"}
	p := config.Pgsql{GeneralDB: x}
	dsn := p.Dsn("postgres")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		er := fmt.Errorf("The error is: %w", err)
		return nil, er
	}
	return db, nil
}
