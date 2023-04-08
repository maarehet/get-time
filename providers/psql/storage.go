package psql

import (
	"context"
	"database/sql"
	"fmt"
	"get-time/providers/psql/migrations"
	"get-time/providers/psql/schema"
	"github.com/jmoiron/sqlx"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/migrate"
	"sync"
)

type Storage struct {
	sync.RWMutex
	db        *sqlx.DB
	dbBun     *bun.DB
	inmemory  map[string]string
	dbAddress string
}

// CloseDb closes DB connection.
func (s Storage) CloseDb() error {
	if s.db == nil {
		return nil
	}
	return s.db.Close()
}

func (s Storage) CloseBun() error {
	if s.dbBun == nil {
		return nil
	}
	return s.dbBun.Close()
}

func (st Storage) UpdateSchema(ctx context.Context) error {
	//logger := st.Logger(withOperation("migration"))
	migrations, err := migrations.GetMigrations()
	if err != nil {
		return err
	}
	migration := migrate.NewMigrator(st.dbBun, migrations)
	if err := migration.Init(ctx); err != nil {
		return fmt.Errorf("initialising migration: %w", err)
	}

	res, err := migration.Migrate(ctx)
	if err != nil {
		return fmt.Errorf("performing migration: %w", err)
	}

	fmt.Printf("Migration applied: %s", res.Migrations.LastGroup().String())
	return nil
}

func NewStorage(dbAddress string) (*Storage, error) {
	storage := Storage{
		inmemory:  make(map[string]string),
		dbAddress: dbAddress,
	}
	clientBan := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(storage.dbAddress)))
	clientSqlx, err := sqlx.Open("pgx", storage.dbAddress)
	if err != nil {
		fmt.Errorf("client sqlx: %v", err)
	}
	storage.dbBun = bun.NewDB(clientBan, pgdialect.New())
	storage.db = clientSqlx
	storage.dbBun.RegisterModel((*schema.TimeDb)(nil))
	if err := storage.dbBun.Ping(); err != nil {
		return nil, fmt.Errorf("ping for DSN (%s) failed: %w", storage.dbAddress, err)
	}
	return &storage, nil
}
