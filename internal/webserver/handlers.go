package webserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sebvautour/gamestats/internal/model"
)

func (s *server) GetTournaments(c *gin.Context) {
	res, err := s.gameStatsAPI.Tournaments(c.Request.Context())
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to retrieve tournaments: %s", err))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *server) PostTournament(c *gin.Context) {
	var t model.Tournament
	if err := c.Bind(&t); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("failed to bind JSON body: %s", err))
		return
	}
	t.ID = uuid.New()

	if err := s.gameStatsAPI.CreateTournament(c.Request.Context(), t); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to create tournament: %s", err))
		return
	}
	c.JSON(http.StatusCreated, t)
}

func (s *server) PutTournament(c *gin.Context) {
	tournamentId, err := uuid.Parse(c.Param("tournamentId"))
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid tournamentId: %s", err))
		return
	}
	var t model.Tournament
	if err := c.Bind(&t); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("failed to bind JSON body: %s", err))
		return
	}
	t.ID = tournamentId

	if err := s.gameStatsAPI.UpdateTournament(c.Request.Context(), t); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to update tournament: %s", err))
		return
	}
	c.JSON(http.StatusOK, t)

}

func (s *server) GetTournamentGames(c *gin.Context) {
	tournamentId, err := uuid.Parse(c.Param("tournamentId"))
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid tournamentId: %s", err))
		return
	}

	res, err := s.gameStatsAPI.Games(c.Request.Context(), tournamentId)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to retrieve tournament games: %s", err))
		return
	}
	c.JSON(http.StatusOK, res)

}

func (s *server) GetGame(c *gin.Context) {
	gameId, err := uuid.Parse(c.Param("gameId"))
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid gameId: %s", err))
		return
	}

	res, err := s.gameStatsAPI.Game(c.Request.Context(), gameId)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to retrieve tournament game: %s", err))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *server) PostGame(c *gin.Context) {
	tournamentId, err := uuid.Parse(c.Param("tournamentId"))
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid tournamentId: %s", err))
		return
	}

	var g model.Game
	if err := c.Bind(&g); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("failed to bind JSON body: %s", err))
		return
	}
	g.ID = uuid.New()
	g.TournamentID = tournamentId

	if err := s.gameStatsAPI.CreateGames(c.Request.Context(), g); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to create tournament game: %s", err))
		return
	}
	c.JSON(http.StatusCreated, g)
}

func (s *server) PutGame(c *gin.Context) {
	tournamentId, err := uuid.Parse(c.Param("tournamentId"))
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid tournamentId: %s", err))
		return
	}
	gameId, err := uuid.Parse(c.Param("gameId"))
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid gameId: %s", err))
		return
	}

	var g model.Game
	if err := c.Bind(&g); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("failed to bind JSON body: %s", err))
		return
	}
	g.ID = gameId
	g.TournamentID = tournamentId

	if err := s.gameStatsAPI.UpdateGame(c.Request.Context(), g); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("failed to update tournament game: %s", err))
		return
	}
	c.JSON(http.StatusOK, g)
}
