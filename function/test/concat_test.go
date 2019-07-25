package math

import (
	"github.com/project-flogo/core/data/expression/function"
	"testing"

	"github.com/stretchr/testify/assert"
)

var in = &fnConcat{}

func init() {
	function.ResolveAliases()
}

func TestInt64Sample(t *testing.T) {
	final, err := in.Eval("123", "-", "456")
	assert.Nil(t, err)
	assert.Equal(t, "123-456", final)
}
