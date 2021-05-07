package decoder

import (
	"bytes"
	"encoding/json"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

type fnDecoder struct {
}

func init() {
	function.Register(&fnDecoder{})
}

func (s *fnDecoder) Name() string {
	return "decodeNumberToString"
}

func (s *fnDecoder) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny}, false
}

func (s *fnDecoder) Eval(in ...interface{}) (interface{}, error) {

	data, err := coerce.ToBytes(in[0])
	if err != nil {
		return nil, err
	}

	var value interface{}

	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	d.Decode(&value)

	return value, nil

}
