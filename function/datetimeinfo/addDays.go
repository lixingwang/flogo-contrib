package math

import (
	"encoding/json"
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"time"
)

type fnAddDays struct {
}

func init() {
	function.Register(&fnAddDays{})
}

func (s *fnAddDays) Name() string {
	return "addDays"
}

func (s *fnAddDays) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeInt}, false
}

func (s *fnAddDays) Eval(in ...interface{}) (interface{}, error) {

	startDate, err := coerce.ToString(in[0])
	fmt.Println(startDate)
	if err != nil {
		return nil, err
	}
	days, err := coerce.ToInt(in[1])
	if err != nil {
		return nil, err
	}

	var FORMAT = "2006-01-02T15:04:05"
	t, err := time.Parse(FORMAT, startDate)
	if err != nil {
		panic(err)
	}

	json.Unmarshal()
	newT := t.Add(time.Duration(days) * time.Hour * 24)
	return newT.Format(FORMAT), nil

}
