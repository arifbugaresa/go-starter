package migration

import (
	"embed"
	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
	"log"
)

//go:embed sql_migrations/*.sql
var dbMigrations embed.FS

func Initiator(dbParam *sqlx.DB) {
	log.Println("Initiating Migrations....")

	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "sql_migrations",
	}

	n, errs := migrate.Exec(dbParam.DB, "postgres", migrations, migrate.Up)
	if errs != nil {
		log.Fatalf("Migration failed: %v", errs)
	}

	log.Println("Migration success, applied", n, "migrations!")
}
