package math

import (
	"fmt"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
	"testing"
)

var in = &fnSubSecond{}

func init() {
	function.ResolveAliases()
}

func TestInt64Sample(t *testing.T) {
	date1 := "2020-03-19T15:02:03"
	date2 := "2020-03-20T15:02:03"
	final, err := in.Eval(date1, date2)
	assert.Nil(t, err)
	fmt.Println("=========", final)
}
