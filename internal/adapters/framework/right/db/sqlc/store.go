package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
}

// Store provides all function to execure db queries and transactons
type PGStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &PGStore{
		Queries: New(db),
		db:      db,
	}
}

func (store *PGStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {

			return fmt.Errorf("tx err: %v, rollback err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
