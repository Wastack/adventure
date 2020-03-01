package yaml

import (
	"bytes"
	"fmt"
	"github.com/Wastack/adventure/engine"
	"gopkg.in/yaml.v2"
	"log"
	"strconv"
)

func Parse_yaml(content []byte, is_verbose bool) (engine.GameDataI, error) {
	y := YamlDataSet{}

	err := yaml.Unmarshal(content, &y)
	if err != nil {
		log.Printf("error while parsing yaml: %v", err)
		return nil, err
	}

	// ensure entry names are unique
	if nue := check_non_unique(&y); len(nue) > 0 {
		return nil, fmt.Errorf("The following entry names are duplicated: %v", nue)
	}

	// yaml parsed, now it's time to turn to internal data by:
	// - copy node contents
	// - searching for start node in entries
	// - Connect entries through  the actions
	d := GameData{
		entries: make([]*GameStateEntry, len(y.Entries)),
	}

	var curr_action_id = 0

	for i, e := range y.Entries {
		d.entries[i] = &GameStateEntry{
			name:    e.Name,
			story:   engine.StoryContent(e.Story),
			is_end:  e.Is_end,
			is_dead: e.Is_dead,
			actions: make(map[engine.ActionId]*InnerActionInfo, len(e.Actions)),
		}
		if e.Name == y.Start_node {
			// assign start node
			d.start_node = d.entries[i]
		}
		for _, a := range e.Actions {
			d.entries[i].actions[engine.ActionId(strconv.Itoa(curr_action_id))] = &InnerActionInfo{
				GameActionInfo: engine.GameActionInfo{
					Target:     a.Target,
					ActionName: a.Name,
					Story:      a.Story,
					Secret:     a.Secret},
			}
			curr_action_id += 1
		}
	}

	// make pointers from actions by iterating over the structure again
	for i := range d.entries {
		for id, aPtr := range d.entries[i].actions {
			if aPtr.Target == "" {
				continue // the action has no target
			}
			found_action_to := false
			for _, e := range d.entries {
				if e.name == aPtr.Target {
					found_action_to = true
					aPtr.to = e
				}
			}
			if !found_action_to {
				return nil, fmt.Errorf("Action with id '%s', target: '%s' does not correspond to any entries", id, aPtr.Target)
			}
		}
	}
	if is_verbose {
		log.Printf(log_data(&d))
	}
	// not connected entries
	nce := check_story_connected(&d)
	if len(nce) > 0 {
		log.Printf("Warning: The following entries are not connected with start point: %v", nce)
	}

	return &d, nil
}

func check_non_unique(d *YamlDataSet) []string {
	occs := make(map[string]int, len(d.Entries))
	for i := range d.Entries {
		if v, has := occs[d.Entries[i].Name]; has {
			occs[d.Entries[i].Name] = v + 1
		} else {
			occs[d.Entries[i].Name] = 1
		}
	}

	result := make([]string, 0, 5)
	for k, v := range occs {
		if v > 1 {
			result = append(result, k)
		}
	}
	return result
}

func check_story_connected(game_data *GameData) []string {
	visited := make(map[*GameStateEntry]struct{}, len(game_data.entries))
	to_explore := make([]*GameStateEntry, 0, len(game_data.entries))
	to_explore = append(to_explore, game_data.start_node)
	for len(to_explore) > 0 {
		current := to_explore[0]
		//remove current
		to_explore = to_explore[1:]
		//check if current already visited
		if _, ok := visited[current]; ok {
			continue
		}
		// append new nodes
		for _, v := range current.actions {
			if v.to != nil {
				to_explore = append(to_explore, v.to)
			}
		}
		// save node as visited
		visited[current] = struct{}{}
	}

	// check equality of visited and original slices
	diff := difference(game_data.entries, visited)
	return diff
}

// difference returns the elements in `a` that aren't in `b`.
func difference(original []*GameStateEntry, b map[*GameStateEntry]struct{}) []string {
	var diff []string
	for _, x := range original {
		if _, found := b[x]; !found {
			diff = append(diff, x.name)
		}
	}
	return diff
}

func log_data(game_data *GameData) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Debugging game data!\n\tStart node: %s\n", game_data.start_node.name))
	for i := range game_data.entries {
		buffer.WriteString(fmt.Sprintf("\tentry name: %s\n", game_data.entries[i].name))
		for id, aPtr := range game_data.entries[i].actions {
			buffer.WriteString(fmt.Sprintf(
				"\t\taction id: %s, target: %s, target ptr: %p, name: %s, story: %s, Secret: %s\n",
				id, aPtr.Target, aPtr.to, aPtr.ActionName, aPtr.Story, aPtr.Secret))
		}
	}
	return buffer.String()
}
