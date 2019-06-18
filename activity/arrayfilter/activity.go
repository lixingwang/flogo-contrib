package arrayfiler

import (
	"encoding/json"
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/resolve"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
)

func init() {
	_ = activity.Register(&Activity{})
}

type Input struct {
	Array []interface{} `md:"array"`      // The message to log
	Expr  string        `md:"expression"` // Append contextual execution information to the log message
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"array":      i.Array,
		"expression": i.Expr,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Array, err = coerce.ToArray(values["array"])
	if err != nil {
		return err
	}
	i.Expr, err = coerce.ToString(values["expression"])
	if err != nil {
		return err
	}

	return nil
}

var activityMd = activity.ToMetadata(&Input{})

// Activity is an Activity that is used to log a message to the console
// inputs : {message, flowInfo}
// outputs: none
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	ctx.GetInputObject(input)

	firstPart := `{
  "input": {
    "mapping": {
      "@foreach($.array, index, `

	objectMpper := firstPart + input.Expr + `)": {
        "=": "$loop"
      }
    }
  }
}`

	mappings := make(map[string]interface{})
	err = json.Unmarshal([]byte(objectMpper), &mappings)
	fmt.Println(err)

	attrs := map[string]interface{}{"array": input.Array}
	scope := data.NewSimpleScope(attrs, nil)

	factory := mapper.NewFactory(resolve.GetBasicResolver())
	mapper, err := factory.NewMapper(mappings)

	output, err := mapper.Apply(scope)

	if err != nil {
		return false, fmt.Errorf("Running array filter error: %s", err.Error())
	}
	out := output["input"]
	s, _ := json.Marshal(out)
	ctx.Logger().Debugf("Output: %s", string(s))
	ctx.SetOutput("outputarray", out)
	return true, nil
}
