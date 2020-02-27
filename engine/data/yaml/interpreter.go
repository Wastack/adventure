package yaml

import (
	"fmt"
	"github.com/Wastack/adventure/engine"
	"gopkg.in/yaml.v2"
	"log"
)

func Parse_yaml(content []byte) (engine.GameDataI, error) {
	y := YamlDataSet{}

	err := yaml.Unmarshal(content, &y)
	if err != nil {
		log.Printf("error while parsing yaml: %v", err)
		return nil, err
	}

	// yaml parsed, now it's time to turn to internal data by:
	// - copy node contents
	// - searching for start node in entries
	// - Connect entries through  the actions
	d := GameData{
		entries: make([]GameStateEntry, len(y.Entries)),
	}

	for i := range y.Entries {
		d.entries[i].name = y.Entries[i].Name
		d.entries[i].story = engine.StoryContent(y.Entries[i].Story)
		d.entries[i].is_end = y.Entries[i].Is_end
		d.entries[i].is_dead = y.Entries[i].Is_dead
		d.entries[i].actions = make(map[engine.GameAction]GameActionInfo)
		if y.Entries[i].Name == y.Start_node {
			d.start_node = &d.entries[i]
		}
		for a_i, a := range y.Entries[i].Actions {
			d.entries[i].actions[engine.GameAction(a_i)] = GameActionInfo{name: a}
		}
	}

	// make pointers from actions by iterating over the structure again
	for i := range d.entries {
		for _, v := range d.entries[i].actions {
			found_action_to := false
			for j := range d.entries {
				if d.entries[j].name == v.name {
					found_action_to = true
					v.to = &d.entries[j]
				}
			}
			if !found_action_to {
				return nil, fmt.Errorf("Action with name '%s' does not correspond to any entries", v.name)
			}
		}
	}
	return &d, nil
}
