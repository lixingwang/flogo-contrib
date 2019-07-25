package math

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

type fnConcat struct {
}

func init() {
	function.Register(&fnConcat{})
}

func (s *fnConcat) Name() string {
	return "concat"
}

func (s *fnConcat) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString, data.TypeString}, false
}

func (s *fnConcat) Eval(in ...interface{}) (interface{}, error) {

	str1, err := coerce.ToString(in[0])
	if err != nil {
		return nil, err
	}
	operator, err := coerce.ToString(in[1])
	if err != nil {
		return nil, err
	}

	str3, err := coerce.ToString(in[2])
	if err != nil {
		return nil, err
	}
	return str1 + operator + str3, nil

}
