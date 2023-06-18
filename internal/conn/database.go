package conn

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func InitDB() (*bun.DB, error) {
	sqldb, err := sql.Open(sqliteshim.ShimName, "data/data.db")
	if err != nil {
		return nil, err
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	return db, nil
}
