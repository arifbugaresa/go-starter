package connection

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"

	_ "github.com/lib/pq"
)

func Initiator() (dbConnection *sqlx.DB, err error) {
	log.Println("Initiating Database Connection....")

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("migration.db.postgres.db_host"),
		viper.GetInt("migration.db.postgres.db_port"),
		viper.GetString("migration.db.postgres.db_user"),
		viper.GetString("migration.db.postgres.db_password"),
		viper.GetString("migration.db.postgres.db_name"),
	)

	dbConnection, err = sqlx.Open("postgres", dsn)

	// check connection
	err = dbConnection.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected to database")

	return
}
