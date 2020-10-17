package main

import (
	"fmt"
	"log"
	"testing"

	_ "database/sql"

	"github.com/aiaoyang/resourceManager/apis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// var (
// 	host   = "127.0.0.1"
// 	port   = 5432
// 	user   = "postgres"
// 	dbname = "testdb"
// )

func Test_jm(t *testing.T) {
	m := &apis.MyResource{}
	// m.Value.Add("key", "value")
	pgsqlStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := gorm.Open("postgres", pgsqlStr)
	if err != nil {
		log.Fatal(err)
	}

	db.Debug().AutoMigrate(m)
	// db.CreateTable(m)
	// db.Model(apis.MyResource{}).AutoMigrate()
}
