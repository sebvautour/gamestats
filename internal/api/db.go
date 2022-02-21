package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

const (
	dbTournamentsTable = "tournaments"
	dbGamesTable       = "games"
	dbId               = "id"
	dbTournamentId     = "tournament_id"
	dbObj              = "obj"
)

type dbWrapper struct {
	db *sql.DB
}

func (w *dbWrapper) init(ctx context.Context) error {
	createTablesStmt := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (%s text not null primary key, %s text);
	CREATE TABLE IF NOT EXISTS %s (%s text not null primary key, %s text not null, %s text);
	`,
		dbTournamentsTable,
		dbId,
		dbObj,
		dbGamesTable,
		dbId,
		dbTournamentId,
		dbObj,
	)

	if _, err := w.db.ExecContext(ctx, createTablesStmt); err != nil {
		return err
	}
	return nil
}

func (w *dbWrapper) insertMany(ctx context.Context, table string, fields []string, values ...[]interface{}) error {
	tx, err := w.db.Begin()
	if err != nil {
		return err
	}
	insertValues := make([]string, len(fields))
	for i := range insertValues {
		insertValues[i] = "?"
	}
	stmt, err := tx.PrepareContext(ctx, fmt.Sprintf(`INSERT INTO %s(%s) values(%s)`,
		table,
		strings.Join(fields, ", "),
		strings.Join(insertValues, ", "),
	))
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, v := range values {
		if _, err = stmt.Exec(v...); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (w *dbWrapper) updateObj(ctx context.Context, table string, id uuid.UUID, obj interface{}) error {
	objBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	stmt := fmt.Sprintf(`UPDATE %s SET %s = '%s' WHERE %s = '%s';`,
		table,
		dbObj,
		string(objBytes),
		dbId,
		id,
	)

	if _, err := w.db.ExecContext(ctx, stmt); err != nil {
		return err
	}
	return nil
}
