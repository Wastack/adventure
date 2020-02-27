package yaml

import (
	"github.com/Wastack/adventure/engine"
)

type GameActionInfo struct {
	to   *GameStateEntry
	name string
}

type GameStateEntry struct {
	actions map[engine.GameAction]GameActionInfo
	name    string
	story   engine.StoryContent
	is_dead bool
	is_end  bool
}

func (e *GameStateEntry) Next(a engine.GameAction) engine.GameNodeI {
	return e.actions[a].to
}

func (e *GameStateEntry) IsGameLost() bool {
	return e.is_dead
}

func (e *GameStateEntry) IsGameOver() bool {
	return e.is_end
}

func (e *GameStateEntry) Actions() map[engine.GameAction]string {
	result := make(map[engine.GameAction]string)
	for k, v := range e.actions {
		result[k] = v.name
	}
	return result
}

func (e *GameStateEntry) Story() engine.StoryContent {
	return e.story
}

type GameData struct {
	entries    []GameStateEntry
	start_node *GameStateEntry
}

func (g *GameData) Start() engine.GameNodeI {
	return g.start_node
}
