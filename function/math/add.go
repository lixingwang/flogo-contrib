package math

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

type fnAdd struct {
}

func init() {
	function.Register(&fnAdd{})
}

func (s *fnAdd) Name() string {
	return "add"
}

func (s *fnAdd) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeInt, data.TypeInt}, false
}

func (s *fnAdd) Eval(in ...interface{}) (interface{}, error) {

	number1, err := coerce.ToInt(in[0])
	if err != nil {
		return nil, err
	}
	number2, err := coerce.ToInt(in[1])
	if err != nil {
		return nil, err
	}
	return number1 + number2, nil

}
