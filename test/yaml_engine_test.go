package engine

import (
	"github.com/Wastack/adventure/engine/data/yaml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func getYamlFromFile(file_path string) []byte {
	file, err := os.Open(file_path)
	if err != nil {
		panic("missing test file")
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic("Error reading file")
	}
	return b
}

func TestYamlEngine(t *testing.T) {
	assert := assert.New(t)
	data, err := yaml.Parse_yaml(getYamlFromFile("data/example.yml"), true)
	assert.Nil(err)
	assert.NotNil(data.Start())

	// start node
	assert.Equal("kezdo_pont", data.Start().Name())
	assert.False(data.Start().IsGameLost())
	assert.False(data.Start().IsGameOver())
	actions := data.Start().Actions()
	assert.Contains(actions, "masodik_pont")
	assert.Contains(actions, "harmadik_pont")
	assert.Equal("Lépés a második pontra", actions["masodik_pont"].Story)

	// masodik_pont
	node := data.Start().Next("masodik_pont")
	assert.Equal("masodik_pont", node.Name())
	assert.False(data.Start().IsGameLost())
	assert.False(data.Start().IsGameOver())
	actions = node.Actions()
	assert.Contains(actions, "kezdo_pont")
	assert.Equal("Visszalépés az első pontra", actions["kezdo_pont"].Story)

	// go back to first node
	assert.Equal(node.Next("kezdo_pont").Name(), "kezdo_pont")
}
