package math

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"time"
)

type fnSub struct {
}

func init() {
	function.Register(&fnSub{})
}

func (s *fnSub) Name() string {
	return "sub"
}

func (s *fnSub) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeInt, data.TypeInt, data.TypeInt}, false
}

func (s *fnSub) Eval(in ...interface{}) (interface{}, error) {

	startDate, err := coerce.ToString(in[0])
	if err != nil {
		return nil, err
	}
	var FORMAT = "2006-01-02T15:04:05"
	datetime, err := time.Parse(FORMAT, startDate)
	if err != nil {
		panic(err)
	}
	years, err := coerce.ToInt(in[1])
	if err != nil {
		return nil, err
	}
	months, err := coerce.ToInt(in[2])
	if err != nil {
		return nil, err
	}
	days, err := coerce.ToInt(in[3])
	if err != nil {
		return nil, err
	}

	year, month, day := datetime.Date()
	hour, min, sec := datetime.Clock()
	newDate := time.Date(year-years, month-time.Month(months), day-days, hour, min, sec, datetime.Nanosecond(), datetime.Location())
	return newDate.Format(FORMAT), nil

}
