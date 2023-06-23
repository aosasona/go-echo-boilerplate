package conn

import (
	"context"
	"database/sql"
	"fmt"
	"gopi/internal/logger"
	"gopi/migrations"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/migrate"
	"go.uber.org/zap"
)

func InitDB(isDev bool) (*bun.DB, error) {
	sqldb, err := sql.Open(sqliteshim.ShimName, "data/data.db")
	if err != nil {
		return nil, err
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true), bundebug.WithEnabled(isDev)))

	if err = db.Ping(); err != nil {
		panic(fmt.Sprintf("failed to ping database: %v", err))
	}

	ctx := context.Background()
	if err = runPragmas(ctx, db); err != nil {
		logger.Logger.Error("failed to run pragmas", zap.Error(err))
	}

	// Run all migrations before starting the app
	if err = runMigration(db); err != nil {
		panic(fmt.Sprintf("failed to run migrations: %v", err))
	}

	return db, nil
}

func runMigration(db *bun.DB) error {
	ctx := context.Background()
	migrator := migrate.NewMigrator(db, migrations.Migrations)
	if err := migrator.Init(ctx); err != nil {
		return err
	}

	if err := migrator.Lock(ctx); err != nil {
		return err
	}
	defer migrator.Unlock(ctx)

	groups, err := migrator.Migrate(ctx)
	if err != nil {
		return err
	}

	if groups.IsZero() {
		fmt.Println("no new migrations to run")
		return nil
	}

	fmt.Printf("migrated to %s\n", groups.String())
	return nil
}

func runPragmas(ctx context.Context, tx bun.IDB) error {
	if _, err := tx.NewRaw("PRAGMA busy_timeout = 10000;").Exec(ctx); err != nil {
		return err
	}

	if _, err := tx.NewRaw("PRAGMA synchronous = NORMAL;").Exec(ctx); err != nil {
		return err
	}

	if _, err := tx.NewRaw("PRAGMA foreign_keys = ON;").Exec(ctx); err != nil {
		return err
	}

	return nil
}
