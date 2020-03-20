package math

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"time"
)

type fnSubSecond struct {
}

func init() {
	function.Register(&fnSubSecond{})
}

func (s *fnSubSecond) Name() string {
	return "subSeconds"
}

func (s *fnSubSecond) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (s *fnSubSecond) Eval(in ...interface{}) (interface{}, error) {

	startDate, err := coerce.ToString(in[0])
	if err != nil {
		return nil, err
	}
	endDate, err := coerce.ToString(in[1])
	if err != nil {
		return nil, err
	}

	var FORMAT = "2006-01-02T15:04:05"
	t, err := time.Parse(FORMAT, startDate)
	if err != nil {
		panic(err)
	}
	t2, err := time.Parse(FORMAT, endDate)
	if err != nil {
		panic(err)
	}

	sub := t2.Sub(t).Seconds()

	return sub, nil

}
