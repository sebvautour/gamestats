package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sebvautour/gamestats/internal/model"
)

func (s *impl) CreateTournament(ctx context.Context, t model.Tournament) error {
	keys := []string{dbId, dbObj}

	objBytes, err := json.Marshal(t)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	return s.db.insertMany(ctx, dbTournamentsTable, keys, []interface{}{t.ID, string(objBytes)})
}

func (s *impl) UpdateTournament(ctx context.Context, t model.Tournament) error {
	return s.db.updateObj(ctx, dbTournamentsTable, t.ID, t)
}

func (s *impl) Tournaments(ctx context.Context) ([]model.Tournament, error) {
	rows, err := s.db.db.QueryContext(ctx, fmt.Sprintf("SELECT %s FROM %s", dbObj, dbTournamentsTable))
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	var tournaments []model.Tournament
	for rows.Next() {
		var obj string
		rows.Scan()
		if err = rows.Scan(&obj); err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		var t model.Tournament
		if err := json.Unmarshal([]byte(obj), &t); err != nil {
			return nil, fmt.Errorf("json.Unarshal: %w", err)
		}
		tournaments = append(tournaments, t)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows: %w", err)
	}

	return tournaments, nil
}
