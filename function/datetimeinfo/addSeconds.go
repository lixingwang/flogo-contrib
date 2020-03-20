package math

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"time"
)

type fnAddSeconds struct {
}

func init() {
	function.Register(&fnAddSeconds{})
}

func (s *fnAddSeconds) Name() string {
	return "addSeconds"
}

func (s *fnAddSeconds) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeInt}, false
}

func (s *fnAddSeconds) Eval(in ...interface{}) (interface{}, error) {

	startDate, err := coerce.ToString(in[0])
	if err != nil {
		return nil, err
	}
	seconds, err := coerce.ToInt64(in[1])
	if err != nil {
		return nil, err
	}
	var FORMAT = "2006-01-02T15:04:05"
	t, err := time.Parse(FORMAT, startDate)
	if err != nil {
		panic(err)
	}

	newT := t.Add(time.Duration(seconds) * time.Second)
	return newT.Format(FORMAT), nil

}
