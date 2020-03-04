package graphgen

import (
	"github.com/Wastack/adventure/engine"
	"github.com/awalterschulze/gographviz"
	"os"
)

func GenerateDotGraph(data engine.GameDataI) (string, error) {
	graphAst, _ := gographviz.ParseString(`digraph G {}`)
	graph := gographviz.NewGraph()
	if err := gographviz.Analyse(graphAst, graph); err != nil {
		return "", err
	}
	visited := make(map[string]struct{})

	// initialize BFS with start node
	to_explore := make([]engine.GameNodeI, 1, 50)
	to_explore[0] = data.Start()
	graph.AddNode("G", data.Start().Name(), nil)

	for len(to_explore) > 0 {
		node := to_explore[0]
		to_explore = to_explore[1:]
		for k := range node.Actions() {
			new_node := node.Next(k)
			if new_node != nil {
				if _, ok := visited[new_node.Name()]; !ok {
					// Add new node
					graph.AddNode("G", new_node.Name(), nil)
					// Add to to_explore
					to_explore = append(to_explore, new_node)
				}
				graph.AddEdge(node.Name(), new_node.Name(), true, nil)
			}
		}
		visited[node.Name()] = struct{}{}
	}
	return graph.String(), nil
}

func WriteDotGraphTo(file_path string, data engine.GameDataI) error {
	f, err := os.Create(file_path)
	if err != nil {
		return err
	}
	defer f.Close()
	dot_string, err := GenerateDotGraph(data)
	if err != nil {
		return err
	}
	_, err = f.WriteString(dot_string)
	if err != nil {
		return err
	}

	f.Sync()
	return nil
}
