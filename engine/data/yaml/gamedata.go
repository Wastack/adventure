package yaml

import (
	"github.com/Wastack/adventure/engine"
	"log"
)

type GameActionInfo struct {
	to    *GameStateEntry
	story string
}

type GameStateEntry struct {
	actions map[string]*GameActionInfo
	name    string
	story   engine.StoryContent
	is_dead bool
	is_end  bool
}

func (e *GameStateEntry) Next(a string) engine.GameNodeI {
	if actionInfo, ok := e.actions[a]; ok {
		return actionInfo.to
	}
	return nil
}

func (e *GameStateEntry) IsGameLost() bool {
	return e.is_dead
}

func (e *GameStateEntry) IsGameOver() bool {
	return e.is_end
}

func (e *GameStateEntry) Actions() map[string]engine.GameActionInfo {
	if e.actions == nil {
		log.Fatalf("Missing action map")
	}
	result := make(map[string]engine.GameActionInfo, len(e.actions))
	i := 0
	for k, v := range e.actions {
		result[k] = engine.GameActionInfo{Story: v.story}
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
	entries    []GameStateEntry
	start_node *GameStateEntry
}

func (g *GameData) Start() engine.GameNodeI {
	return g.start_node
}

func (g *GameData) GetNodeByString(name string) engine.GameNodeI {
	for i := range g.entries {
		if g.entries[i].name == name {
			return &g.entries[i]
		}
	}
	log.Printf("Node not found: %s", name)
	return nil
}
