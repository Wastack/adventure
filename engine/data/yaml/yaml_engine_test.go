package yaml

import (
	"github.com/Wastack/adventure/engine"
	"github.com/Wastack/adventure/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func findActionById(id string, node engine.GameNodeI) *engine.GameActionInfo {
	for _, v := range node.Actions() {
		if v.ActionId == id {
			return &v
		}
	}
	return nil
}

func TestYamlEngine(t *testing.T) {
	assert := assert.New(t)
	data, err := Parse_yaml(utils.GetYamlFromFile("testdata/example.yml"), true)
	assert.Nil(err)
	assert.NotNil(data.Start())

	// start node
	assert.Equal("kezdo_pont", data.Start().Name())
	assert.False(data.Start().IsGameLost())
	assert.False(data.Start().IsGameOver())

	mpont_action := findActionById("masodik_pont", data.Start())
	hpont_action := findActionById("harmadik_pont", data.Start())
	assert.NotNil(mpont_action)
	assert.NotNil(hpont_action)
	assert.Equal("masodik_pont", mpont_action.ActionId)
	assert.Equal("harmadik_pont", hpont_action.ActionId)
	assert.Equal("Lépés a második pontra", findActionById("masodik_pont", data.Start()).ActionName)

	// masodik_pont
	node := data.Start().Next("masodik_pont")
	assert.Equal("masodik_pont", node.Name())
	assert.False(data.Start().IsGameLost())
	assert.False(data.Start().IsGameOver())
	kpont_action := findActionById("kezdo_pont", node)
	assert.NotNil(kpont_action)
	assert.Equal("kezdo_pont", kpont_action.ActionId)
	assert.Equal("Visszalépés az első pontra", findActionById("kezdo_pont", node).ActionName)

	// go back to first node
	assert.Equal(node.Next("kezdo_pont").Name(), "kezdo_pont")

	// assert that graph is connected
	game_data := data.(*GameData)
	assert.Empty(check_story_connected(game_data))
}
