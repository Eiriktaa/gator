package state

import (
	"database/sql"
	"log"

	"example.com/eiriktaa/gator/internal/config"
	"example.com/eiriktaa/gator/internal/database"
)

type State struct {
	Config *config.Config
	DB     *database.Queries
}

func InitalizeState() State {
	cfg := config.LoadConfiguration()
	db, err := sql.Open("postgres", cfg.DB_url)
	if err != nil {
		log.Fatal("failed to open DB", err)
	}
	dbQueries := database.New(db)
	return State{
		Config: &cfg,
		DB:     dbQueries,
	}
}

func (s *State) SetCurrentUser(username string) {
	s.Config.SetUser(username)
}
