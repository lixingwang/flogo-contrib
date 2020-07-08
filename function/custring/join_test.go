package custring

import (
	"testing"

	"github.com/project-flogo/core/data/expression/function"

	"github.com/stretchr/testify/assert"
)

var in = &fnJoin{}

func init() {
	function.ResolveAliases()
}

func TestInt64Sample(t *testing.T) {
	var a = []string{"flogo", "test"}
	final, err := in.Eval(a, ";")
	assert.Nil(t, err)
	assert.Equal(t, "flogo;test", final)
}
