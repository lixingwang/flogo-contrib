package decoder

import (
	"encoding/json"
	"github.com/project-flogo/core/data/expression/function"
	"testing"

	"github.com/stretchr/testify/assert"
)

var in = &fnDecoder{}

func init() {
	function.ResolveAliases()
}

func TestDecover(t *testing.T) {
	var data = []byte("{\n  \"zcrm_record_id\": 3570049000074512286883737,\n  \"zcrm_module\": \"Leads\",\n  \"zcrm_action\": \"Create\"\n}")

	value, err := in.Eval(data)

	assert.Nil(t, err)
	assert.Equal(t, "3570049000074512286883737", value.(map[string]interface{})["zcrm_record_id"].(json.Number).String())
}
