package migrations

import (
	"embed"
	"fmt"

	"github.com/uptrace/bun/migrate"
)

//go:embed *.sql
var sqlMigrations embed.FS

var Migrations = migrate.NewMigrations()

func init() {
	fmt.Println("Running migrations...")
	if err := Migrations.DiscoverCaller(); err != nil {
		panic(err)
	}
	fmt.Println("Migrations ran successfully")
}
