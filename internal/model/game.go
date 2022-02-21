package model

import (
	"time"

	"github.com/google/uuid"
)

type Game struct {
	ID           uuid.UUID
	TournamentID uuid.UUID
	Name         string
	Date         time.Time
	HomeTeam     GameTeam
	AwayTeam     GameTeam
	Stats        GameStats
}

type GameTeam struct {
	Name     string
	Logo     string
	PlayedBy string
	Type     string
}

type GameStats struct {
	HomeScore int
	AwayScore int
}
