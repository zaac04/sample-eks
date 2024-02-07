package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Pg Postgres

type Postgres struct {
	Client *gorm.DB
}

func (pg *Postgres) Connect() {
	var err error
	dsn := os.Getenv("PSQL_URL")
	pg.Client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(err)
}