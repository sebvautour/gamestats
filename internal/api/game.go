package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/sebvautour/gamestats/internal/model"
)

func (s *impl) CreateGames(ctx context.Context, g ...model.Game) error {
	keys := []string{dbId, dbTournamentId, dbObj}
	values := make([][]interface{}, len(g))
	for i := range g {
		objBytes, err := json.Marshal(g[i])
		if err != nil {
			return fmt.Errorf("json.Marshal: %w", err)
		}
		values[i] = []interface{}{g[i].ID, g[i].TournamentID, objBytes}
	}

	return s.db.insertMany(ctx, dbGamesTable, keys, values...)
}

func (s *impl) UpdateGame(ctx context.Context, g model.Game) error {
	return s.db.updateObj(ctx, dbTournamentsTable, g.ID, g)
}

func (s *impl) Games(ctx context.Context, tournamentID uuid.UUID) ([]model.Game, error) {
	rows, err := s.db.db.QueryContext(ctx, fmt.Sprintf("SELECT %s FROM %s WHERE %s = '%s';", dbObj, dbGamesTable, dbTournamentId, tournamentID.String()))
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	var games []model.Game
	for rows.Next() {
		var obj string
		if err = rows.Scan(&obj); err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		var g model.Game
		if err := json.Unmarshal([]byte(obj), &g); err != nil {
			return nil, fmt.Errorf("json.Unmarshal: %w", err)
		}
		games = append(games, g)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows: %w", err)
	}

	return games, nil
}

func (s *impl) Game(ctx context.Context, gameID uuid.UUID) (*model.Game, error) {

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s = '%s';", dbObj, dbGamesTable, dbId, gameID.String())
	log.Println(query)
	row := s.db.db.QueryRowContext(ctx, query)
	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	var obj string
	if err := row.Scan(&obj); err != nil {
		return nil, fmt.Errorf("scan: %w", err)
	}

	log.Println(obj)
	var g model.Game
	if err := json.Unmarshal([]byte(obj), &g); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}
	return &g, nil
}
