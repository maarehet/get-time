package psql

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"get-time/pkg/providers/psql/schema"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/migrate"
	"sync"
)

//go:embed migrations
var migrationFS embed.FS

type Storage struct {
	sync.RWMutex
	db        *sqlx.DB
	dbBun     *bun.DB
	dbAddress string
	log       *logrus.Entry
}

func getMigrations() (*migrate.Migrations, error) {
	migrations := migrate.NewMigrations()
	if err := migrations.Discover(migrationFS); err != nil {
		return nil, fmt.Errorf("discovering migrations by caller: %w", err)
	}

	return migrations, nil
}

// CloseDb closes DB connection.
func (s *Storage) CloseDb() error {
	if s.db == nil {
		return nil
	}
	return s.db.Close()
}

func (s *Storage) CloseBun() error {
	if s.dbBun == nil {
		return nil
	}
	return s.dbBun.Close()
}

func (s *Storage) UpdateSchema(ctx context.Context) error {
	migrations, err := getMigrations()
	if err != nil {
		return err
	}
	migration := migrate.NewMigrator(s.dbBun, migrations)
	if err := migration.Init(ctx); err != nil {
		return fmt.Errorf("initialising migration: %w", err)
	}

	res, err := migration.Migrate(ctx)
	if err != nil {
		return fmt.Errorf("performing migration: %w", err)
	}

	s.log.Infof("Migration applied: %s", res.Migrations.LastGroup().String())
	return nil
}

func NewStorage(log *logrus.Logger, dbAddress string) (*Storage, error) {
	storage := Storage{
		dbAddress: dbAddress,
		log:       log.WithField("module", "storage"),
	}
	clientBan := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(storage.dbAddress)))
	clientSqlx, err := sqlx.Open("pgx", storage.dbAddress)
	if err != nil {
		return nil, fmt.Errorf("client sqlx: %v", err)
	}
	storage.dbBun = bun.NewDB(clientBan, pgdialect.New())
	storage.db = clientSqlx
	storage.dbBun.RegisterModel((*schema.TimeDb)(nil))
	if err := storage.dbBun.Ping(); err != nil {
		return nil, fmt.Errorf("ping for DSN (%s) failed: %w", storage.dbAddress, err)
	}
	return &storage, nil
}
