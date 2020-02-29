package yaml

import (
	"bytes"
	"fmt"
	"github.com/Wastack/adventure/engine"
	"gopkg.in/yaml.v2"
	"log"
)

func Parse_yaml(content []byte, is_verbose bool) (engine.GameDataI, error) {
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
		d.entries[i].actions = make(map[string]*GameActionInfo)
		if y.Entries[i].Name == y.Start_node {
			d.start_node = &d.entries[i]
		}
		for _, a := range y.Entries[i].Actions {
			d.entries[i].actions[a.Target] = &GameActionInfo{story: a.Text}
		}
	}

	// make pointers from actions by iterating over the structure again
	for i := range d.entries {
		for k := range d.entries[i].actions {
			found_action_to := false
			for j := range d.entries {
				if d.entries[j].name == k {
					found_action_to = true
					d.entries[i].actions[k].to = &d.entries[j]
				}
			}
			if !found_action_to {
				return nil, fmt.Errorf("Action with name '%s' does not correspond to any entries", k)
			}
		}
	}
	if is_verbose {
		log.Printf(log_data(&d))
	}
	return &d, nil
}

func log_data(game_data *GameData) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Debugging game data!\n\tStart node: %s\n", game_data.start_node.name))
	for i := range game_data.entries {
		buffer.WriteString(fmt.Sprintf("\tentry name: %s\n", game_data.entries[i].name))
		for k, v := range game_data.entries[i].actions {
			buffer.WriteString(fmt.Sprintf(
				"\t\taction name: %s, target node: %p, story: %s\n", k, v.to, v.story))
		}
	}
	return buffer.String()
}
