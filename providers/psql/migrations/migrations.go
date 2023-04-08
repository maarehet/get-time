package migrations

import (
	"fmt"
	"github.com/uptrace/bun/migrate"
)

func GetMigrations() (*migrate.Migrations, error) {
	migrations := migrate.NewMigrations()
	if err := migrations.DiscoverCaller(); err != nil {
		return nil, fmt.Errorf("discovering migrations by caller: %w", err)
	}

	return migrations, nil
}
