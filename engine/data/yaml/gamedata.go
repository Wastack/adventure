package yaml

import (
	"github.com/Wastack/adventure/engine"
	"log"
)

type GameActionInfo struct {
	to          *GameStateEntry
	action_name string
	story       string
	action_id   string
}

type GameStateEntry struct {
	actions []*GameActionInfo
	name    string
	story   engine.StoryContent
	is_dead bool
	is_end  bool
}

func (e *GameStateEntry) Next(id string) engine.GameNodeI {
	for i := range e.actions {
		if e.actions[i].action_id == id {
			return e.actions[i].to
		}
	}
	return nil
}

func (e *GameStateEntry) IsGameLost() bool {
	return e.is_dead
}

func (e *GameStateEntry) IsGameOver() bool {
	return e.is_end
}

func (e *GameStateEntry) Actions() []engine.GameActionInfo {
	if e.actions == nil {
		log.Fatalf("Missing action map")
	}
	result := make([]engine.GameActionInfo, len(e.actions))
	for i, a := range e.actions {
		result[i] = engine.GameActionInfo{ActionId: a.action_id, ActionName: a.action_name, Story: a.story}
		i += 1
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
