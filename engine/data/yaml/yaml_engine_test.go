package yaml

import (
	"github.com/Wastack/adventure/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYamlEngine(t *testing.T) {
	assert := assert.New(t)
	data, err := Parse_yaml(utils.GetYamlFromFile("testdata/example.yml"), true)
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

	// assert that graph is connected
	game_data := data.(*GameData)
	assert.Empty(check_story_connected(game_data))
}
