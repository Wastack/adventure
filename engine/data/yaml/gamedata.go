package yaml

import (
	"github.com/Wastack/adventure/engine"
	"log"
)

type InnerActionInfo struct {
	to *GameStateEntry
	engine.GameActionInfo
}

type GameStateEntry struct {
	actions map[engine.ActionId]*InnerActionInfo
	name    string
	story   engine.StoryContent
	is_dead bool
	is_end  bool
}

func (e *GameStateEntry) Next(id engine.ActionId) engine.GameNodeI {
	if v, ok := e.actions[id]; ok {
		if v.to == nil {
			return nil
		}
		return v.to
	}
	return nil
}

func (e *GameStateEntry) IsGameLost() bool {
	return e.is_dead
}

func (e *GameStateEntry) IsGameOver() bool {
	return e.is_end
}

func (e *GameStateEntry) Actions() map[engine.ActionId]engine.GameActionInfo {
	if e.actions == nil {
		log.Fatalf("Missing action map")
	}
	result := make(map[engine.ActionId]engine.GameActionInfo, len(e.actions))
	for k, v := range e.actions {
		result[k] = v.GameActionInfo
	}
	return result
}

func (e *GameStateEntry) Name() string {
	return e.name
}

func (e *GameStateEntry) Story() engine.StoryContent {
	return e.story
}

type GameData struct {
	entries    []*GameStateEntry
	start_node *GameStateEntry
}

func (g *GameData) Start() engine.GameNodeI {
	return g.start_node
}

func (g *GameData) GetNodeByString(name string) engine.GameNodeI {
	for i := range g.entries {
		if g.entries[i].name == name {
			return g.entries[i]
		}
	}
	log.Printf("Node not found: %s", name)
	return nil
}
