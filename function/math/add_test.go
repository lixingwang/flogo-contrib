package math

import (
	"fmt"
	"github.com/project-flogo/core/data/expression/function"
	"testing"

	"github.com/stretchr/testify/assert"
)

var in = &Int{}

func init() {
	function.ResolveAliases()
}

func TestInt64Sample(t *testing.T) {
	final, err := in.Eval("123")
	assert.Nil(t, err)
	assert.Equal(t, int(123), final)
}

func TestInt64Expression(t *testing.T) {
	fun, err := factory.NewExpr(`number.int64("123")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, int(123), v)
	fmt.Println(v)
}
