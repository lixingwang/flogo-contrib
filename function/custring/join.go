package custring

import (
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

type fnJoin struct {
}

func init() {
	function.Register(&fnJoin{})
}

func (s *fnJoin) Name() string {
	return "join"
}

func (s *fnJoin) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeArray, data.TypeString}, false
}

func (s *fnJoin) Eval(in ...interface{}) (interface{}, error) {

	array, err := coerce.ToArray(in[0])
	if err != nil {
		return nil, err
	}
	split, err := coerce.ToString(in[1])
	if err != nil {
		return nil, err
	}
	return strings.Join(toStringArray(array), split), nil
}

func toStringArray(a []interface{}) []string {
	var strArray []string
	for _, element := range a {
		str, _ := coerce.ToString(element)
		strArray = append(strArray, str)
	}
	return strArray
}
