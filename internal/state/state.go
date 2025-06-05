package state

import "example.com/eiriktaa/gator/internal/config"

type State struct {
	Config *config.Config
}

func InitalizeState() State {
	cfg := config.LoadConfiguration()
	return State{
		Config: &cfg,
	}
}

func (s *State) SetCurrentUser(username string) {
	s.Config.SetUser(username)
}
