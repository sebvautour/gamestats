package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/sebvautour/gamestats/internal/api"
)

type server struct {
	gameStatsAPI api.API
}

func New(gameStatsAPI api.API) *gin.Engine {
	r := gin.Default()

	svr := &server{gameStatsAPI: gameStatsAPI}

	tournamentsV1 := r.Group("/api/v1/tournaments")
	{
		tournamentsV1.GET("/", svr.GetTournaments)
		tournamentsV1.POST("/", svr.PostTournament)
		tournamentsV1.PUT("/:tournamentId", svr.PutTournament)
		tournamentsV1.GET("/:tournamentId/games", svr.GetTournamentGames)
		tournamentsV1.POST("/:tournamentId/games", svr.PostGame)
		tournamentsV1.GET("/:tournamentId/games/:gameId", svr.GetGame)
		tournamentsV1.PUT("/:tournamentId/games/:gameId", svr.PutGame)
	}

	return r
}
