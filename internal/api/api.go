package api

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/sebvautour/gamestats/internal/model"
)

type API interface {
	CreateTournament(ctx context.Context, t model.Tournament) error
	UpdateTournament(ctx context.Context, t model.Tournament) error
	Tournaments(ctx context.Context) ([]model.Tournament, error)

	CreateGames(ctx context.Context, g ...model.Game) error
	UpdateGame(ctx context.Context, g model.Game) error
	Games(ctx context.Context, tournamentID uuid.UUID) ([]model.Game, error)
	Game(ctx context.Context, gameID uuid.UUID) (*model.Game, error)
}

type impl struct {
	db *dbWrapper
}

func New(ctx context.Context, db *sql.DB) (API, error) {
	s := &impl{db: &dbWrapper{db: db}}
	return s, s.db.init(ctx)
}
