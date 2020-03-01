package yaml

import (
	"github.com/Wastack/adventure/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYamlNotConnected(t *testing.T) {
	assert := assert.New(t)
	data, err := Parse_yaml(utils.GetYamlFromFile("testdata/not_connected.yml"), true)
	game_data := data.(*GameData)
	assert.Nil(err)

	// check for not connected entries
	nce := check_story_connected(game_data)
	assert.Equal(1, len(nce))
	assert.Equal("harmadik_pont", nce[0])
}
