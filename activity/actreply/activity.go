package actreply

import (
	"fmt"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
)

func init() {
	_ = activity.Register(&Activity{}, New)
}

type Input struct {
	Input map[string]interface{} `md:"input"` // Set of mappings to execute when the activity runs
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"input": i.Input,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {
	var err error
	i.Input, err = coerce.ToObject(values["input"])
	if err != nil {
		return err
	}
	return nil
}

var activityMd = activity.ToMetadata(&Input{})

func New(ctx activity.InitContext) (activity.Activity, error) {
	act := &Activity{}
	return act, nil
}

// Activity is an Activity that is used to reply/return via the trigger
// inputs : {method,uri,params}
// outputs: {result}
type Activity struct {
}

func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Invokes a REST Operation
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {
	actionCtx := ctx.ActivityHost()
	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return false, fmt.Errorf("error getting input object: %s", err.Error())
	}
	actionCtx.Reply(input.Input, nil)
	return true, nil
}
